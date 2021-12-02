package types

import (
	"io"
	"math/rand"
	"sync"
	"time"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"google.golang.org/protobuf/proto"
)

// GeneratorAlignmentFn is a function type which is capable of yielding a Prometheus MetricFamily for each step (time.Duration) between a start and end time.Time
type GeneratorAlignmentFn func(<-chan struct{}, *dto.MetricFamily, time.Time, time.Time, time.Duration, *MetricsGenerator) <-chan *dto.MetricFamily

type GeneratorJob struct {
	FamData *dto.MetricFamily
	GenType GeneratorType 
}

type GeneratorType int

func (g GeneratorType) String() string {
	switch g{
	case Repeatable:
		return "repeatable"
	case Chaotic:
		return "chaotic"
	case Lossy:
		return "lossy"
	default:
		return ""
	}
}

const (
	Repeatable GeneratorType = iota // Repeatable represents a generator that never changes values per interval
	Chaotic // Chaotic represents a generator that does change values per interval
	Lossy // Lossy represents a generator that will selectively drop a series
)

// MetricsGenerator is a coordinating structure for generating "fake" OpenMetrics datasets.
type MetricsGenerator struct {
	alignment GeneratorType // alignment represents this generator's treatment of values
	stepDuration time.Duration // stepDuration represents the interval between generated scrape times (e.g. 1 minute, 5 minutes, etc.)
	seedData []*dto.MetricFamily // seedData represents the base data (ideally, via a live scrape) this generator uses
	stepMaxVariance float64 // stepMaxVariance represents how the upper bound a given metric should change per interval
	alignmentFn GeneratorAlignmentFn
	
	out io.Writer // out is the destination io.Writer for this generator
}

// GetAlignment sets a generator's alignment. If already set, this is a noop
func (m *MetricsGenerator) GetAlignment() string {
	return m.alignment.String()
}

// WithStepDuration sets a generator's step.
func (m *MetricsGenerator) WithStepDuration(t time.Duration) *MetricsGenerator {
	m.stepDuration = t
	return m
}

// WithGaugeVariance sets a generator's gauge variance.
// Gauge variance is the measure of how much a gauge value changes, upon a change event.
func (m *MetricsGenerator) WithGaugeVariance(f float64) *MetricsGenerator {
	m.stepMaxVariance = f
	return m
}

// NewGenerator returns a bare-bones metrics generator.
func NewGenerator(o io.Writer, genType GeneratorType, seed []*dto.MetricFamily) *MetricsGenerator {
	retGen := &MetricsGenerator{
		out: o,
		alignment: genType,
		seedData: seed,
	}
	switch retGen.alignment {
	case Repeatable:
		retGen.alignmentFn = repeatMetricFam
	case Chaotic:
		retGen.alignmentFn = chaoticMetricValues
	}

	return retGen
}

// WriteOpenMetrics writes the OpenMetrics-compatible output of m.Generate to m's output buffer.
// It is preferred to call WriteOpenMetrics instead of Generate directly.
func (m *MetricsGenerator) WriteOpenMetrics(done <-chan struct{}, epoch, end time.Time) bool {
	defer expfmt.FinalizeOpenMetrics(m.out)

	for mFam := range m.Generate(done, epoch, end) {
		expfmt.MetricFamilyToOpenMetrics(m.out, mFam)
	}

	return true
}

// Generate provides a return channel representing generated MetricFamilies.
// It does this for every tick of epoch + m.stepDuration, up to forDuration.
// epoch is defined as the input variable of the same name.
// For each MetricFamily in the generator's seed data, a fanout pattern is employed.
func (m *MetricsGenerator) Generate(done <-chan struct{}, epoch, end time.Time) <-chan *dto.MetricFamily{
	returnCh := make(chan *dto.MetricFamily)

	var wg sync.WaitGroup

	merge := func(c <-chan *dto.MetricFamily) {
		defer wg.Done()
		for c1 := range c{
			select {
			case returnCh <- c1:
			case <- done:
				return
			}
		}
	}

	wg.Add(len(m.seedData))
	for f := range yieldFams(m.seedData) {
		go merge(m.alignmentFn(done, f, epoch, end, m.stepDuration, m))
	}
	
	go func(){
		wg.Wait()
		close(returnCh)
	}()

	return returnCh
}

func yieldFams(f []*dto.MetricFamily) <-chan *dto.MetricFamily {
	out := make(chan *dto.MetricFamily)

	go func() {
		defer close(out)
		for _, v := range f {
			out <- v
		}
	}()

	return out
}

// chaoticMetricValues adjusts each gauge metric in a family to be +/- some input value per step, based on a 0.0 - 1.0 probability of change
func chaoticMetricValues(done <-chan struct{}, in *dto.MetricFamily, epoch, end time.Time, step time.Duration, m *MetricsGenerator) <-chan *dto.MetricFamily {
	out := make(chan *dto.MetricFamily)

	go func(){
		defer close(out)
		nTime := epoch

		for {
			if nTime.Add(step).After(end) {
				return
			}
			nTime = nTime.Add(step)

			if in.GetType() != dto.MetricType_GAUGE {
				out <- &dto.MetricFamily{
					Name: in.Name,
					Help: in.Help,
					Type: in.Type,
					Metric: modifyMetrics(in.Metric, nTime), 
				}
			} else{
				out <- &dto.MetricFamily{
					Name: in.Name,
					Help: in.Help,
					Type: in.Type,
					Metric: adjustVals(modifyMetrics(in.Metric, nTime), m.stepMaxVariance, 0.1),
				}	
			}
		}
	}()

	return out
}

// For each input gauge, evaluate if a psuedo-random float64 is >= prob.
// If so,
//	- Determine if a second psuedo-random float64 >= 0.50
//		- If so, increment the input metric's value by r
//		- If not, decrement the input metric's value by r
// If not, do nothing
func adjustVals(in []*dto.Metric, r, prob float64) []*dto.Metric {
	for _, m := range in {
		if rand.Float64() >= prob {
			if rand.Float64() >= 0.50 {
				m.Gauge.Value = proto.Float64(*m.Gauge.Value + r)
			} else {
				m.Gauge.Value = proto.Float64(*m.Gauge.Value - r)
			}
		} else {
			continue
		}
	}

	return in
}

// repeatMetricFam just returns the input family for each step
func repeatMetricFam(done <-chan struct{}, in *dto.MetricFamily, epoch, end time.Time, step time.Duration, g *MetricsGenerator) <-chan *dto.MetricFamily {
	out := make(chan *dto.MetricFamily)

	go func(){
		defer close(out)
		nTime := epoch

		for {
			if nTime.Add(step).After(end) {
				return
			}

			nTime = nTime.Add(step)
			out <- &dto.MetricFamily{
				Name: in.Name,
				Help: in.Help,
				Type: in.Type,
				Metric: modifyMetrics(in.Metric, nTime),
			}
		}
	}()

	return out
}

// modifyMetrics makes a set of established changes to an input slice of Metrics.
// TODO(mfuller): think about how to improve this. See the weird ducttape for chaotic metrics where we are chaning functions as a prime example.
func modifyMetrics(in []*dto.Metric, t time.Time) []*dto.Metric {
	r := make([]*dto.Metric, 0)

	for _, metric := range in {
		r1 := setTimestamp(metric, t)
		r1 = setReqdLabels(&r1, "testjob", "testinstance:7777")
		r = append(r, &r1)
	}

	return r
}

// setTimestamp sets a Metric's TimestampMs to the provided input Time.
// On Golang>=1.17, the division used may be done away with in lieu of 
// http://docs.studygolang.com/pkg/time/#UnixMilli
func setTimestamp(onMetric *dto.Metric, t time.Time) dto.Metric {
	return dto.Metric{
		Label: onMetric.Label,
		Gauge: onMetric.Gauge,
		Counter: onMetric.Counter,
		Summary: onMetric.Summary,
		Untyped: onMetric.Untyped,
		Histogram: onMetric.Histogram,
		TimestampMs: proto.Int64(t.UnixNano() / int64(time.Millisecond)),
	}
}

// setReqdLabels injects two labels which are required by Prometheus - job and instance
func setReqdLabels(onMetric *dto.Metric, job, instance string) dto.Metric {
	origLabels := onMetric.Label
	origLabels = append(origLabels,
		&dto.LabelPair{
			Name: proto.String("job"),
			Value: proto.String(job),
		},
		&dto.LabelPair{
			Name: proto.String("instance"),
			Value: proto.String(instance),
		},
	)

	return dto.Metric{
		Label: origLabels,
		Gauge: onMetric.Gauge,
		Counter: onMetric.Counter,
		Summary: onMetric.Summary,
		Untyped: onMetric.Untyped,
		Histogram: onMetric.Histogram,
		TimestampMs: onMetric.TimestampMs,
	}
}