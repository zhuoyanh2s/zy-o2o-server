package models

import (
	"time"
)

type ShopFriendshiplink struct {
	Id         int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn  time.Time `xorm:"not null DATETIME"`
	ModifiedOn time.Time `xorm:"not null DATETIME"`
	Name       string    `xorm:"not null VARCHAR(20)"`
	Url        string    `xorm:"not null VARCHAR(250)"`
	Sort       int       `xorm:"not null INTEGER"`
}
