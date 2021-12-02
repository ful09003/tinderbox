package scrape

import (
	"bytes"
	"net/http"

	"github.com/ful09003/tinderbox/pkg/types"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

// ScrapeJob represents a desired target and options to scrape
type ScrapeJob struct {
	target string // Target to scrape, must represent a Prometheus exporter's full FQDN/path
	opts   *types.TinderboxHTTPOptions
}

func NewScrapeJob(t string, o *types.TinderboxHTTPOptions) *ScrapeJob {
	return &ScrapeJob{
		target: t,
		opts:   o,
	}
}

// ScrapeResults wraps the resulting Prometheus MetricFamily parse result, and any errors encountered during a scrape request.
type ScrapeResults struct {
	families map[string]*dto.MetricFamily
	err      error
}

func (r *ScrapeResults) Families() map[string]*dto.MetricFamily {
	return r.families
}
func (r *ScrapeResults) Error() error {
	return r.err
}

// OpenMetricScrape takes a scrape job, a client, and output channel and writes the resulting ScrapeResults to outCh.
// BUG(mfuller): This naming is incorrect, and should be updated to reflect that this is not truly OpenMetrics format, but rather Prometheus exposition format.
func OpenMetricScrape(in *ScrapeJob, c http.Client) <-chan ScrapeResults {

	out := make(chan ScrapeResults)

	go func() {
		defer close(out)
		scrapeBytes, err := Scrape(in.target, in.opts, c)
		if err != nil {
			out <- ScrapeResults{
				families: nil,
				err:      err,
			}
			return
		}
		f, e := Parse(scrapeBytes)
		out <- ScrapeResults{
			families: f,
			err:      e,
		}
	}()

	return out

	/*var res ScrapeResults

	scrapeBytes, err := Scrape(in.target, in.opts, c)
	if err != nil {
		res.err = err
		outCh <- res
		return
	}

	f, e := Parse(scrapeBytes)

	res.families = f
	res.err = e

	outCh <- res*/
}

func Parse(b []byte) (map[string]*dto.MetricFamily, error) {
	reader := bytes.NewReader(b)

	var parser expfmt.TextParser

	mFam, err := parser.TextToMetricFamilies(reader)

	return mFam, err
}
