package models

import (
	"time"
)

type PromotionCouponlog struct {
	Id            int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn     time.Time `xorm:"not null DATETIME"`
	ModifiedOn    time.Time `xorm:"not null DATETIME"`
	CouponName    string    `xorm:"not null VARCHAR(50)"`
	OrderNo       string    `xorm:"not null index VARCHAR(60)"`
	Denominations string    `xorm:"not null NUMERIC"`
	OriginFee     string    `xorm:"not null NUMERIC"`
	DiscountFee   string    `xorm:"not null NUMERIC"`
	PayFee        string    `xorm:"not null NUMERIC"`
}
