package models

import (
	"time"
)

type TradeOrder struct {
	Id                  int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn           time.Time `xorm:"not null DATETIME"`
	ModifiedOn          time.Time `xorm:"not null DATETIME"`
	OrderNo             string    `xorm:"not null index VARCHAR(60)"`
	PayNo               string    `xorm:"not null VARCHAR(60)"`
	MailNo              string    `xorm:"default 'NULL::character varying' VARCHAR(60)"`
	OrderStatus         int       `xorm:"not null index SMALLINT"`
	PostFee             string    `xorm:"not null NUMERIC"`
	DiscountFee         string    `xorm:"not null NUMERIC"`
	OriginFee           string    `xorm:"not null NUMERIC"`
	PayFee              string    `xorm:"not null NUMERIC"`
	PayTime             time.Time `xorm:"DATETIME"`
	ReceivedPayment     string    `xorm:"not null NUMERIC"`
	ConsigneeName       string    `xorm:"not null VARCHAR(100)"`
	ConsigneeMobile     string    `xorm:"not null VARCHAR(20)"`
	ConsigneeTelephone  string    `xorm:"not null VARCHAR(50)"`
	ConsigneeAddress    string    `xorm:"not null VARCHAR(200)"`
	ConsigneePostalCode string    `xorm:"default 'NULL::character varying' VARCHAR(20)"`
	UserMem             string    `xorm:"TEXT"`
	UserId              int       `xorm:"not null index INTEGER"`
}
