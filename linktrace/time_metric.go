package linktrace

import (
	"sync"
	"time"

	"github.com/ihauk/gokit/logger"
	// "github.com/ihauk/gokit"
)

// usage :
// t :=NewTimeMetrics()
// fin := StartOneMetric("a"), fin()
// fin := StartOneMetric("b"), fin()
// ...
// fin := StartOneMetric("x"),fin()
// t.Finish(), then can call ts.String() for option
//
// best practice: one context(session) one instence for
type TimeMetrics struct {
	metricsCh chan interface{}
	once      sync.Once
	result    string
	finishCh  chan struct{}
	isFinish  bool
}

func NewTimeMetrics() *TimeMetrics {
	t := &TimeMetrics{}
	t.metricsCh = make(chan interface{}, 60)
	t.finishCh = make(chan struct{}, 1)

	go t.handleMetrics()
	return t
}

// 计时结束时调用返回值
func (ts *TimeMetrics) StartOneMetric(metric string) func() time.Duration {

	if ts.metricsCh == nil {
		panic("call NewTimeMetrics() first")
	}
	if ts.isFinish {
		panic("instence has finished, please renew an install. one context(session) one instence for best practice")
	}

	one := &timeMetric{}
	one.beginWithMetirc(metric)

	finish := func() time.Duration {
		// fmt.Println("finish ", one.metric)
		dur := one.endWithMetirc()
		if !ts.isFinish {
			ts.metricsCh <- one
		} else {
			logger.ErrorLog.Println("method use error: sending after finish")
		}

		return dur
	}

	return finish
}

func (ts *TimeMetrics) Finish() string {

	ts.once.Do(func() {
		// close(ts.metricsCh)
		// close(ts.isFinish)
		ts.metricsCh <- struct{}{}
	})

	// select {
	// case <-ts.finishCh:
	// }
	<-ts.finishCh

	return ts.result
}

func (ts *TimeMetrics) String() string {

	return ts.result
}

func (ts *TimeMetrics) handleMetrics() {
	ts.result = "{ "

METRICSLOOP:
	for m := range ts.metricsCh {

		switch val := m.(type) {
		case *timeMetric:
			ts.result += val.String() + " "
		case struct{}:
			break METRICSLOOP
		}

	}
	ts.result += "}"

	close(ts.finishCh)
	close(ts.metricsCh)
	ts.isFinish = true
}

type timeMetric struct {
	metric   string
	begin    time.Time
	end      time.Time
	duration time.Duration
}

func (t *timeMetric) beginWithMetirc(metric string) {
	t.metric = metric
	t.begin = time.Now()
}

func (t *timeMetric) endWithMetirc() time.Duration {
	t.end = time.Now()
	t.duration = t.end.Sub(t.begin)

	return t.duration
}

func (t *timeMetric) String() string {
	t.duration = t.end.Sub(t.begin)

	return "[" + t.metric + ":" + t.duration.String() + "]"
}
