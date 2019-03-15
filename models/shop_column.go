package models

import (
	"time"
)

type ShopColumn struct {
	Id                   int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn            time.Time `xorm:"not null DATETIME"`
	ModifiedOn           time.Time `xorm:"not null DATETIME"`
	Name                 string    `xorm:"not null VARCHAR(8)"`
	NameIsDisplay        bool      `xorm:"not null BOOL"`
	Url                  string    `xorm:"default 'NULL::character varying' VARCHAR(250)"`
	HeadBackgroundPc     string    `xorm:"default 'NULL::character varying' VARCHAR(100)"`
	HeadBackgroundMobile string    `xorm:"default 'NULL::character varying' VARCHAR(100)"`
	Type                 string    `xorm:"not null VARCHAR(1)"`
	Sort                 int       `xorm:"not null INTEGER"`
	IsDisplay            bool      `xorm:"not null BOOL"`
}
