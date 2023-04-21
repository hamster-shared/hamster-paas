package nginx_log_parse

import (
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/utils/logger"
	"os"
	"time"

	"github.com/meilisearch/meilisearch-go"
)

type LogParser struct {
	client     *meilisearch.Client
	latestTime int64
}

func InitMeiliSearch() {
	for i := 0; i < 10; i++ {
		p, err := NewLogParser()
		if err != nil {
			logger.Errorf("init meili search error: %s, meili search host: %s, nginx log path: %s", err, os.Getenv("MEILI_SEARCH"), os.Getenv("NGINX_LOG_PATH"))
			time.Sleep(5 * time.Second)
			if i == 9 {
				panic(err)
			}
		} else {
			application.SetBean("meiliSearch", p.client)
			break
		}
	}
}

func NewLogParser() (*LogParser, error) {
	// 创建一个美丽搜索客户端
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: os.Getenv("MEILI_SEARCH"),
	})
	p := &LogParser{
		client: client,
	}
	// 读取 nginx 日志文件
	file, err := os.Open(os.Getenv("NGINX_LOG_PATH"))
	if err != nil {
		return nil, fmt.Errorf("open file: %s, error: %s", os.Getenv("NGINX_LOG_PATH"), err)
	}
	defer file.Close()

	// 创建索引
	index := p.client.Index("nginx")
	// 美丽搜索默认是容错的，我们不需要这个功能，禁用它
	_, err = index.UpdateTypoTolerance(&meilisearch.TypoTolerance{
		Enabled: false,
	})
	if err != nil {
		logger.Errorf("设置美丽搜索不容错失败: %s", err)
	}
	// 添加文档，主键为 request_id
	_, err = index.AddDocumentsNdjsonFromReader(file, "request_id")
	if err != nil {
		logger.Errorf("meili search index error: %s", err)
		return nil, err
	}
	taskInfo, err := index.UpdateSearchableAttributes(&[]string{"msec", "status", "request_uri"})
	if err != nil {
		logger.Errorf("meili search index error: %s", err)
		return nil, err
	}
	p.latestTime = time.Now().Unix()
	if taskInfo.Status == "failed" {
		logger.Errorf("meili search index failed: %v", taskInfo)
		return nil, fmt.Errorf("meili search index failed: %v", taskInfo)
	}

	// 设置过滤器
	filters := []string{
		"msec", "status", "request_uri",
	}
	settings := meilisearch.Settings{
		FilterableAttributes: filters,
		SortableAttributes:   filters,
	}
	_, err = client.Index("nginx").UpdateSettings(&settings)
	if err != nil {
		logger.Errorf("meili search update settings error: %s", err)
		return nil, err
	}

	// 定时更新文档，每 1 秒
	go p.TimedUpdate(1)

	return p, nil
}

func (p *LogParser) Update() error {
	file, err := os.Open(os.Getenv("NGINX_LOG_PATH"))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = p.client.Index("nginx").UpdateDocumentsNdjsonFromReader(file)
	return err
}

func (p *LogParser) TimedUpdate(s int64) {
	for {
		if time.Now().Unix()-p.latestTime > s {
			err := p.Update()
			if err != nil {
				logger.Errorf("meili search update error: %s", err)
			}
		}
		time.Sleep(time.Second)
	}
}
