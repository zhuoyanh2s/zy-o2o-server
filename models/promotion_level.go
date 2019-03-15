package models

import (
	"time"
)

type PromotionLevel struct {
	Id         int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn  time.Time `xorm:"not null DATETIME"`
	ModifiedOn time.Time `xorm:"not null DATETIME"`
	Points     int       `xorm:"not null INTEGER"`
	UserId     int       `xorm:"not null unique INTEGER"`
}
