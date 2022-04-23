package linktrace

import "context"

type LinkContext struct {
	context.Context
	kvMetric   *KVMetrics
	timeMetric *TimeMetrics
}

func NewLinkContext() *LinkContext {

	lc := &LinkContext{}

	return lc
}
