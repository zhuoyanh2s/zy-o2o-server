package models

import (
	"time"
)

type PaymentAllinpaymentmessage struct {
	Id           int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn    time.Time `xorm:"not null DATETIME"`
	ModifiedOn   time.Time `xorm:"not null DATETIME"`
	CusId        string    `xorm:"not null VARCHAR(20)"`
	AllinAppid   string    `xorm:"not null VARCHAR(50)"`
	AllinKey     string    `xorm:"not null VARCHAR(64)"`
	Version      string    `xorm:"default 'NULL::character varying' VARCHAR(64)"`
	CustomsType  string    `xorm:"default 'NULL::character varying' VARCHAR(64)"`
	EshopEntCode string    `xorm:"default 'NULL::character varying' VARCHAR(64)"`
	EshopEntName string    `xorm:"default 'NULL::character varying' VARCHAR(64)"`
	IsDefault    bool      `xorm:"not null BOOL"`
	IsActive     bool      `xorm:"not null BOOL"`
}
