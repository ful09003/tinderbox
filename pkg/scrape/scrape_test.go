// Package scrape encompasses one core concern, summarized as "scrape a Prometheus endpoint and return data from it"
package scrape

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/ful09003/tinderbox/pkg/types"
)

func TestScrape(t *testing.T) {
	s := newTestHTTPServer(t)

	type args struct {
		target string
		o      *types.TinderboxHTTPOptions
		c      http.Client
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "happy round trip",
			args: args{
				target: fmt.Sprintf("%s/happy_path", s.URL),
				o:      types.NewTinderboxHTTPOptions(),
				c:      *s.Client(),
			},
			want:    []byte(happyRes),
			wantErr: false,
		},
		{
			name: "5xx",
			args: args{
				target: fmt.Sprintf("%s/5xx", s.URL),
				o:      types.NewTinderboxHTTPOptions(),
				c:      *s.Client(),
			},
			want:    []byte(nil),
			wantErr: true,
		},
		{
			name: "retryable",
			args: args{
				target: fmt.Sprintf("%s/flappy", s.URL),
				o:      types.NewTinderboxHTTPOptions(),
				c:      *s.Client(),
			},
			want:    []byte(happyRes),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Scrape(tt.args.target, tt.args.o, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scrape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scrape() = %v, want %v", got, tt.want)
			}
		})
	}
}
