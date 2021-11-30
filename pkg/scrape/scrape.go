// Package scrape encompasses one core concern, summarized as "scrape a Prometheus endpoint and return data from it"
package scrape

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ful09003/tinderbox/pkg/types"
)

// Scrape accepts a 'target' string, TinderboxHTTPOptions struct, and HTTP Client.
// Scrape returns the associated raw bytes discovered from scraping a (presumably) OpenMetrics-format endpoint
func Scrape(target string, o *types.TinderboxHTTPOptions, c http.Client) ([]byte, error) {
	req, err := o.ToRequest(target)
	if err != nil {
		return nil, err
	}

	return retryHttpOperation(req, c, o.Retries(), o.BackoffDelay())
}

// retryHttpOperation will retry an HTTP request using the provided client up to
// a configured number of times. Upon a non-success response or error,
// a delay of (current_retry_number * backoffDelay) seconds is introduced.
// A header is automatically added for each request to indicate to servers
// how many attempts this particular invocation has made
func retryHttpOperation(req *http.Request, c http.Client, r int, backoffDelay int) ([]byte, error) {

	var response *http.Response
	var err error

	for i := 0; i < r; i++ {
		req.Header.Set("X-Tinderbox-Retry", fmt.Sprintf("%d", i))
		response, err = c.Do(req)

		if err != nil || response.StatusCode != http.StatusOK {
			delay := i * backoffDelay
			time.Sleep(time.Duration(delay) * time.Second)
		} else {
			// Request was likely OK!
			return ioutil.ReadAll(response.Body)
		}
	}

	return nil, errors.New("exhausted retries for http operation")
}
