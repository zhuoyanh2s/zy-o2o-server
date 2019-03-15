package models

import (
	"time"
)

type ShopIndexbanner struct {
	Id         int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn  time.Time `xorm:"not null DATETIME"`
	ModifiedOn time.Time `xorm:"not null DATETIME"`
	Name       string    `xorm:"not null VARCHAR(20)"`
	Image      string    `xorm:"not null VARCHAR(100)"`
	Url        string    `xorm:"not null VARCHAR(250)"`
	Sort       int       `xorm:"not null INTEGER"`
}
