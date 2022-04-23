package example

import "context"

type IRequestDomain interface {
	RequestDB(ctx context.Context) IRequestDB
}

type RequestDomain struct {
}

func NewRequestDomain() IRequestDomain {
	return &RequestDomain{}
}

func (x *RequestDomain) RequestDB(ctx context.Context) IRequestDB {
	return NewRequestDB(ctx)
}
