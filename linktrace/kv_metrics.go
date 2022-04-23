package linktrace

import (
	"fmt"
	"sync"
)

type KVMetrics struct {
	metricsCh chan interface{}
	once      sync.Once
	result    string
	finishCh  chan struct{}
	isFinish  bool
}

func NewKVMetrics() *KVMetrics {
	t := &KVMetrics{}
	t.metricsCh = make(chan interface{}, 60)
	t.finishCh = make(chan struct{}, 1)

	go t.handleMetrics()
	return t
}

func (kv *KVMetrics) SetMetricWithString(key string, value string) {
	kv.check()
	t := kvMetric{key: key, value: value}
	kv.metricsCh <- &t
	// select {
	// case <-kv.finishCh:
	// 	fmt.Println("finished when write")
	// case kv.metricsCh <- &t:
	// 	return
	// }

}

func (kv *KVMetrics) SetMetricWithStrings(key string, values []string) {
	kv.check()
	t := kvMetric{key: key, value: fmt.Sprint(values)}
	kv.metricsCh <- &t
	// select {
	// case <-kv.finishCh:
	// 	fmt.Println("finished when write")
	// case kv.metricsCh <- &t:

	// }

}

func (kv *KVMetrics) SetMetricWithObject(key string, object interface{}) {
	kv.check()
	t := kvMetric{key: key, value: fmt.Sprint(object)}
	kv.metricsCh <- &t
	// select {
	// case <-kv.finishCh:
	// 	fmt.Println("finished when write")
	// case kv.metricsCh <- &t:

	// }

}

func (ts *KVMetrics) Finish() string {

	ts.once.Do(func() {
		// close(ts.metricsCh)
		// close(ts.isFinish)
		ts.metricsCh <- struct{}{}
	})

	<-ts.finishCh

	return ts.result
}

func (kv *KVMetrics) String() string {
	return kv.result
}

func (kv *KVMetrics) check() {
	if kv.metricsCh == nil {
		panic("call NewTimeMetrics() first")
	}
	if kv.isFinish {
		panic("instence has finished, please renew an install. one context(session) one instence for best practice")
	}
}

func (kv *KVMetrics) handleMetrics() {
	kv.result = "{ "
METRICSLOOP:
	for one := range kv.metricsCh {

		switch val := one.(type) {
		case *kvMetric:
			kv.result += val.String() + " "
		case struct{}:
			fmt.Println("finish cmd")
			break METRICSLOOP
		}

	}
	kv.result += "}"

	close(kv.finishCh)
	close(kv.metricsCh)

	kv.isFinish = true
}

type kvMetric struct {
	key   string
	value string
}

func (t kvMetric) String() string {

	return "[" + t.key + ":" + t.value + "]"
}
