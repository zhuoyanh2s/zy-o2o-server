package models

import "time"

type Store struct {
	Id          int
	Name        string
	StoreLogo   string    `orm:"null"`
	StoreHead   string    `orm:"null"`
	StoreQRCode string    `orm:"null"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
	User        *User     `orm:"rel(fk)"`
}
