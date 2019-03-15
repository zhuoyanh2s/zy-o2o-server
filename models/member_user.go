package models

import (
	"time"
)

type MemberUser struct {
	Id            int       `xorm:"not null pk autoincr INTEGER"`
	Password      string    `xorm:"not null VARCHAR(128)"`
	LastLogin     time.Time `xorm:"DATETIME"`
	IsSuperuser   bool      `xorm:"not null BOOL"`
	Username      string    `xorm:"not null unique VARCHAR(150)"`
	FirstName     string    `xorm:"not null VARCHAR(30)"`
	LastName      string    `xorm:"not null VARCHAR(30)"`
	Email         string    `xorm:"not null VARCHAR(254)"`
	IsStaff       bool      `xorm:"not null BOOL"`
	IsActive      bool      `xorm:"not null BOOL"`
	DateJoined    time.Time `xorm:"not null DATETIME"`
	Mobile        string    `xorm:"default 'NULL::character varying' VARCHAR(20)"`
	HeadImg       string    `xorm:"not null BYTEA(255)"`
	RealName      string    `xorm:"default 'NULL::character varying' VARCHAR(20)"`
	Gender        bool      `xorm:"not null BOOL"`
	Birthday      time.Time `xorm:"not null DATE"`
	Address       string    `xorm:"default 'NULL::character varying' VARCHAR(50)"`
	Qq            string    `xorm:"default 'NULL::character varying' VARCHAR(50)"`
	Buyeridnumber string    `xorm:"default 'NULL::character varying' VARCHAR(60)"`
	IdCardImage   string    `xorm:"not null VARCHAR(100)"`
	IsVerify      string    `xorm:"not null VARCHAR(12)"`
}
