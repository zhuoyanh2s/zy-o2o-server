package models

import (
	"time"
)

type PaymentPayment struct {
	Id            int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn     time.Time `xorm:"not null DATETIME"`
	ModifiedOn    time.Time `xorm:"not null DATETIME"`
	OrderNo       string    `xorm:"not null VARCHAR(64)"`
	TradeNo       string    `xorm:"not null VARCHAR(64)"`
	PaymentId     string    `xorm:"default 'NULL::character varying' VARCHAR(64)"`
	Type          string    `xorm:"default 'NULL::character varying' VARCHAR(20)"`
	Message       string    `xorm:"not null TEXT"`
	Price         string    `xorm:"not null NUMERIC"`
	NotifyMessage string    `xorm:"default 'NULL::character varying' VARCHAR(1024)"`
}
