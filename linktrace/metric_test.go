package linktrace

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type Address struct {
	Country  string
	Shengfen string
	quhao    int
}

func TestKVMetric(t *testing.T) {

	kv := &kvMetric{}
	fmt.Println(kv)

	kt := NewKVMetrics()
	go func() {
		kt.SetMetricWithString("name", "haukzhu")

	}()
	go func() {
		kt.SetMetricWithStrings("phone", []string{"180", "134"})
	}()

	kt.SetMetricWithObject("address", Address{"china", "baoding", 1})

	time.Sleep(time.Second)
	fmt.Println(kt.Finish())

}

func TestTimeMetric(t *testing.T) {

	kt := NewTimeMetrics()
	go func() {
		aaFun := kt.StartOneMetric("aa")
		defer aaFun()

		time.Sleep(time.Second)

	}()
	go func() {
		aaFun := kt.StartOneMetric("bb")
		defer aaFun()

		time.Sleep(time.Millisecond * 1400)
	}()

	aaFun := kt.StartOneMetric("cc")
	aaFun()

	time.Sleep(time.Second * 2)

	fmt.Println(kt.Finish())

}

func TestLinkContext(t *testing.T) {
	ctx := LinkContext{}
	<-ctx.Done()

}

func CalleeA(ctx *LinkContext, val int) {
	_, c := context.WithCancel(ctx)
	defer c()

	select {
	case <-ctx.Done():
	}
}

func CalleeB(ctx *LinkContext, val string) {

}
