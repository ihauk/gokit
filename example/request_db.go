package example

import (
	"context"

	"gorm.io/gorm"
)

func init() {
	dbcore.RegisterInjector(func(d *gorm.DB) {
		dbcore.SetupTableModel(d, &Request{})
	})
}

type IRequestDB interface {
	Get(id string) (*Request, error)
	List(query *Request, offset, limit int) ([]*Request, error)
	Create(query *Request) (*Request, error)
	Update(query *Request) (*Request, error)
	Delete(query *Request) error
}

type RequestDB struct {
	db *gorm.DB
}

func NewRequestDB(ctx context.Context) *RequestDB {
	return &RequestDB{dbcore.GetDB(ctx)}
}

func (x *RequestDB) Get(id string) (*Request, error) {
	panic("not implemented") // TODO: Implement
}

func (x *RequestDB) List(query *Request, offset int, limit int) ([]*Request, error) {
	panic("not implemented") // TODO: Implement
}

func (x *RequestDB) Create(query *Request) (*Request, error) {
	err := x.db.Create(query).Error
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (x *RequestDB) Update(query *Request) (*Request, error) {
	err := x.db.Save(query).Error
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (x *RequestDB) Delete(query *Request) error {
	panic("not implemented") // TODO: Implement
}
