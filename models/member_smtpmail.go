package models

import (
	"time"
)

type MemberSmtpmail struct {
	Id                int       `xorm:"not null pk autoincr INTEGER"`
	CreatedOn         time.Time `xorm:"not null DATETIME"`
	ModifiedOn        time.Time `xorm:"not null DATETIME"`
	EmailHost         string    `xorm:"not null VARCHAR(50)"`
	EmailHostUser     string    `xorm:"not null VARCHAR(50)"`
	EmailHostPassword string    `xorm:"not null VARCHAR(50)"`
	DefaultFromEmail  string    `xorm:"default 'NULL::character varying' VARCHAR(250)"`
	IsDefault         bool      `xorm:"not null BOOL"`
	IsActive          bool      `xorm:"not null BOOL"`
}
