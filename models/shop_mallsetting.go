package models

import (
	"time"
)

type ShopMallsetting struct {
	Id                      int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn               time.Time `xorm:"not null DATETIME"`
	ModifiedOn              time.Time `xorm:"not null DATETIME"`
	MallLogo                string    `xorm:"not null VARCHAR(100)"`
	BrBackground            string    `xorm:"not null VARCHAR(100)"`
	StoreAdminUrl           string    `xorm:"default 'NULL::character varying' VARCHAR(50)"`
	InternetContentProvider string    `xorm:"default 'NULL::character varying' VARCHAR(50)"`
	MallColor               string    `xorm:"not null TEXT"`
	IsDefault               bool      `xorm:"not null BOOL"`
	SiteId                  int       `xorm:"not null unique INTEGER"`
}
