package nginx_log_parse

import (
	"hamster-paas/pkg/utils/logger"
	"testing"
	"time"

	"github.com/meilisearch/meilisearch-go"
)

func TestNewLogParser(t *testing.T) {

	logger.InitLogger()
	p, err := NewLogParser()
	if err != nil {
		panic(err)
	}
	// spew.Dump(p)
	for {
		time.Sleep(time.Second)
		resp, err := p.client.Index("nginx").Search("jlvihv", &meilisearch.SearchRequest{
			Filter: "status = 404",
		})
		if err != nil {
			panic(err)
		}
		logger.Infof("resp: %v", resp.EstimatedTotalHits)
	}
}
