package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/ful09003/tinderbox/pkg/scrape"
	"github.com/ful09003/tinderbox/pkg/types"
	dto "github.com/prometheus/client_model/go"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	instance = "http://localhost:9100/metrics"
	jobName  = "demojob"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	job := scrape.NewScrapeJob(instance, types.NewTinderboxHTTPOptions())

	allFams := make([]*dto.MetricFamily, 0)

	for sRes := range scrape.OpenMetricScrape(job, *http.DefaultClient) {
		if sRes.Error() != nil {
			log.Fatalln(sRes.Error())
		}
		for _, v := range sRes.Families() {
			allFams = append(allFams, v)
		}
	}

	gen := types.NewGenerator(os.Stdout, types.Chaotic, instance, jobName, allFams).WithStepDuration(2 * time.Minute).WithGaugeVariance(1.0)

	gen.WriteOpenMetrics(done, time.Now().Add(-12*time.Hour), time.Now().Add(-2*time.Hour))

}
