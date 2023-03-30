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
	p, err := NewLogParser()
	if err != nil {
		panic(err)
	}
	application.SetBean("meiliSearch", p.client)
}

func NewLogParser() (*LogParser, error) {
	// 创建一个美丽搜索客户端
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: "http://localhost:7700",
	})
	p := &LogParser{
		client: client,
	}
	// 读取 nginx 日志文件
	file, err := os.Open(os.Getenv("NGINX_LOG_PATH"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建索引
	index := p.client.Index("nginx")
	// 添加文档，主键为 request_id
	taskInfo, err := index.AddDocumentsNdjsonFromReader(file, "request_id")
	if err != nil {
		logger.Errorf("meili search index error: %s", err)
		return nil, err
	}
	logger.Debugf("meili search index taskInfo: %v", taskInfo)
	p.latestTime = time.Now().Unix()
	if taskInfo.Status == "failed" {
		logger.Errorf("meili search index failed: %v", taskInfo)
		return nil, fmt.Errorf("meili search index failed: %v", taskInfo)
	}

	// 设置过滤器
	filters := []string{
		"msec", "connection", "connection_requests", "pid", "request_id", "request_length", "remote_addr", "remote_user", "remote_port", "time_local", "time_iso8601", "request", "request_uri", "args", "status", "body_bytes_sent", "bytes_sent", "http_referer", "http_user_agent", "http_x_forwarded_for", "http_host", "server_name", "request_time", "upstream", "upstream_connect_time", "upstream_header_time", "upstream_response_time", "upstream_response_length", "upstream_cache_status", "ssl_protocol", "ssl_cipher", "scheme", "request_method", "server_protocol", "pipe", "gzip_ratio", "http_cf_ray",
	}
	settings := meilisearch.Settings{
		FilterableAttributes: filters,
		SortableAttributes:   filters,
	}
	resp, err := client.Index("nginx").UpdateSettings(&settings)
	if err != nil {
		logger.Errorf("meili search update settings error: %s", err)
		return nil, err
	}
	logger.Debugf("meili search update settings: %v", resp)

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
