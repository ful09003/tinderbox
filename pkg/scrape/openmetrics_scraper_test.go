package scrape

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/ful09003/tinderbox/pkg/types"
	dto "github.com/prometheus/client_model/go"
	"google.golang.org/protobuf/proto"
)

func TestOpenMetricScrape(t *testing.T) {
	s := newTestHTTPServer(t)

	type args struct {
		in    *ScrapeJob
		c     http.Client
		outCh chan ScrapeResults
	}
	tests := []struct {
		name    string
		args    args
		want    ScrapeResults
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "happy path",
			args: args{
				in: &ScrapeJob{
					target: fmt.Sprintf("%s/happy_path", s.URL),
					opts:   types.NewTinderboxHTTPOptions(),
				},
				c:     *s.Client(),
				outCh: make(chan ScrapeResults, 1),
			},
			want: ScrapeResults{
				err: nil,
				families: map[string]*dto.MetricFamily{
					"joy_felt_total": &dto.MetricFamily{
						Name: proto.String("joy_felt_total"),
						Help: proto.String("A counter of joy experienced."),
						Type: dto.MetricType_COUNTER.Enum(),
						Metric: []*dto.Metric{
							&dto.Metric{
								Label: []*dto.LabelPair{
									&dto.LabelPair{
										Name:  proto.String("developer"),
										Value: proto.String("me"),
									},
								},
								Counter: &dto.Counter{
									Value: proto.Float64(9000),
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for got := range OpenMetricScrape(tt.args.in, tt.args.c) {
				if (got.Error() != nil) != tt.wantErr {
					t.Errorf("OpenMetricScrape() error = %v, wantErr %v", got.err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got.Families(), tt.want.families) {
					t.Errorf("OpenMetricScrape() = %v, want %v", got, tt.want)
					return
				}	
			}
		})
	}
}
