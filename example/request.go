package example

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type RequestStatus int64

const (
	InQueue RequestStatus = iota
	Fetched
	Succeed
	Failed
)

func (rs RequestStatus) String() string {
	return [...]string{"inqueue", "fetched", "succeed", "failed"}[rs]
}

func (u *RequestStatus) Scan(value interface{}) error {
	*u = RequestStatus(value.(int64))
	return nil
}
func (u RequestStatus) Value() (driver.Value, error) {
	return u.String(), nil
}

type Request struct {
	gorm.Model
	RequestEncoding  string
	ResponseEncoding string
	LanguageCode     string
	Volume           float32
	Pitch            float32
	Speaker          string
	RequestId        string
	RequestModel     string
	Uri              string
	JobId            string //For Redis Stream
	Name             string //For Operations
	Status           RequestStatus
	// dbCli mysql.MySQLClient `gorm:"-"`
}

func NewRequest() *Request {
	item := &Request{}
	//赋值

	return item
}
