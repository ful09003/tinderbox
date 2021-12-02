package types

import (
	"io"
	"reflect"
	"testing"
	"time"

	dto "github.com/prometheus/client_model/go"
	"google.golang.org/protobuf/proto"
)

func TestMetricsGenerator_GetAlignment(t *testing.T) {
	type fields struct {
		alignment       GeneratorType
		stepDuration    time.Duration
		seedData        []*dto.MetricFamily
		stepMaxVariance float64
		alignmentFn     GeneratorAlignmentFn
		out             io.Writer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "happy path",
			fields: fields{
				alignment: Chaotic,
			},
			want: "chaotic",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MetricsGenerator{
				alignment:       tt.fields.alignment,
				stepDuration:    tt.fields.stepDuration,
				seedData:        tt.fields.seedData,
				stepMaxVariance: tt.fields.stepMaxVariance,
				alignmentFn:     tt.fields.alignmentFn,
				out:             tt.fields.out,
			}
			if got := m.GetAlignment(); got != tt.want {
				t.Errorf("MetricsGenerator.GetAlignment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetricsGenerator_WithStepDuration(t *testing.T) {
	type fields struct {
		alignment GeneratorType
		job, instance string
		seedData  []*dto.MetricFamily
		out       io.Writer
	}
	type args struct {
		t time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MetricsGenerator
	}{
		// TODO: Add test cases.
		{
			name: "happy path",
			fields: fields{
				out:       nil,
				alignment: Chaotic,
				seedData:  []*dto.MetricFamily{},
				job: "testj",
				instance: "testi",
			},
			args: args{
				t: 5 * time.Minute,
			},
			want: &MetricsGenerator{
				stepDuration:    5 * time.Minute,
				alignment:       Chaotic,
				alignmentFn:     chaoticMetricValues,
				seedData:        []*dto.MetricFamily{},
				stepMaxVariance: 0,
				out:             nil,
				instanceName: "testi",
				jobName: "testj",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGenerator(tt.fields.out, tt.fields.alignment, tt.fields.job, tt.fields.instance, tt.fields.seedData).WithStepDuration(tt.args.t); !reflect.DeepEqual(got.stepDuration, tt.want.stepDuration) {
				t.Errorf("MetricsGenerator.WithStepDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setReqdLabels(t *testing.T) {
	type args struct {
		onMetric *dto.Metric
		job      string
		instance string
	}
	tests := []struct {
		name string
		args args
		want dto.Metric
	}{
		// TODO: Add test cases.
		{
			name: "happy path",
			args: args{
				onMetric: &dto.Metric{
					Label: []*dto.LabelPair{
						{Name: proto.String("who"), Value: proto.String("you")},
					},
				},
				job:      "testjob",
				instance: "testinstance:7777",
			},
			want: dto.Metric{
				Label: []*dto.LabelPair{
					{Name: proto.String("who"), Value: proto.String("you")},
					{Name: proto.String("job"), Value: proto.String("testjob")},
					{Name: proto.String("instance"), Value: proto.String("testinstance:7777")},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setReqdLabels(tt.args.onMetric, tt.args.job, tt.args.instance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setReqdLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setTimestamp(t *testing.T) {
	testTime, err := time.Parse("2006-01-02", "2021-12-01")
	if err != nil {
		t.Error(err)
	}

	type args struct {
		onMetric *dto.Metric
		t        time.Time
	}
	tests := []struct {
		name string
		args args
		want dto.Metric
	}{
		// TODO: Add test cases.
		{
			name: "happy path",
			args: args{
				onMetric: &dto.Metric{
					TimestampMs: proto.Int64(0),
				},
				t: testTime,
			},
			want: dto.Metric{
				TimestampMs: proto.Int64(testTime.UnixNano() / int64(time.Millisecond)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setTimestamp(tt.args.onMetric, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
