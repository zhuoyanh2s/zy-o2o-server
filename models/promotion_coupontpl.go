package models

import (
	"time"
)

type PromotionCoupontpl struct {
	Id            int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn     time.Time `xorm:"not null DATETIME"`
	ModifiedOn    time.Time `xorm:"not null DATETIME"`
	CouponName    string    `xorm:"not null VARCHAR(50)"`
	IsActive      bool      `xorm:"not null BOOL"`
	MinOriginFee  string    `xorm:"not null NUMERIC"`
	Denominations string    `xorm:"not null NUMERIC"`
	Image         string    `xorm:"default 'NULL::character varying' VARCHAR(100)"`
	CouponCount   int64     `xorm:"not null BIGINT"`
	StartTime     time.Time `xorm:"not null DATETIME"`
	EndTime       time.Time `xorm:"not null DATETIME"`
}
