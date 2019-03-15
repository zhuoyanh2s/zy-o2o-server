package models

import (
	"time"
)

type PaymentPaymentnotifylog struct {
	Id            int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn     time.Time `xorm:"not null DATETIME"`
	ModifiedOn    time.Time `xorm:"not null DATETIME"`
	PaymentName   string    `xorm:"not null VARCHAR(64)"`
	PaymentType   string    `xorm:"not null VARCHAR(64)"`
	NotifyMessage string    `xorm:"TEXT"`
}
