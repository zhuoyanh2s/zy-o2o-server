package models

import (
	"time"
)

type PaymentPaymentlog struct {
	Id          int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn   time.Time `xorm:"not null DATETIME"`
	ModifiedOn  time.Time `xorm:"not null DATETIME"`
	OrderNo     string    `xorm:"not null VARCHAR(60)"`
	PaymentName string    `xorm:"not null VARCHAR(64)"`
	PaymentType string    `xorm:"not null VARCHAR(64)"`
	Req         string    `xorm:"TEXT"`
	Res         string    `xorm:"TEXT"`
}
