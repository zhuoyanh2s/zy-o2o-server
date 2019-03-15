package models

import (
	"time"
)

type ShopHotsearch struct {
	Id         int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn  time.Time `xorm:"not null DATETIME"`
	ModifiedOn time.Time `xorm:"not null DATETIME"`
	Name       string    `xorm:"not null VARCHAR(20)"`
	Sort       int       `xorm:"not null INTEGER"`
}
