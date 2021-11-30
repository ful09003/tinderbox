package scrape

import (
	"bytes"
	"net/http"

	"github.com/ful09003/tinderbox/pkg/types"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

type ScrapeJob struct {
	target string
	opts   *types.TinderboxHTTPOptions
}

type ScrapeResults struct {
	families map[string]*dto.MetricFamily
	err      error
}

func OpenMetricScrape(in *ScrapeJob, c http.Client, outCh chan<- ScrapeResults) {
	var res ScrapeResults

	scrapeBytes, err := Scrape(in.target, in.opts, c)
	if err != nil {
		res.err = err
		outCh <- res
		return
	}

	reader := bytes.NewReader(scrapeBytes)

	var parser expfmt.TextParser

	mFam, err := parser.TextToMetricFamilies(reader)

	res.families = mFam
	res.err = err

	outCh <- res
}
