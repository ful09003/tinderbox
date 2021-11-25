package scrape

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	happyRes string = `# HELP joy_felt_total A counter of joy experienced.
	# TYPE joy_felt_total counter
	joy_felt_total{developer="me"} 9000
`
)
// newTestHTTPServer is a functional (if not small) HTTP Server used to unit test scraping needs.
func newTestHTTPServer(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var s string
		var rH int

		switch r.URL.String(){
		case "/happy_path":
			rH = http.StatusOK
			s = happyRes
		case "/5xx":
			rH = http.StatusInternalServerError
			s = ""
		case "/flappy":
			retryAttempt := r.Header.Get("X-Tinderbox-Retry")
			t.Logf("flappy attempt %s", retryAttempt)
			if retryAttempt == "0" {
				rH = http.StatusServiceUnavailable
				s = ""	
			} else {
				rH = http.StatusOK
				s = happyRes
			}
		default:
			rH = http.StatusForbidden
			s = "Forbidden"
		}

		rw.WriteHeader(rH)
		rw.Write([]byte(s))
	}))

	return server
}
