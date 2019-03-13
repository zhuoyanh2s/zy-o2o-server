package models

import "time"

//店铺头像
type StoreHeadImage struct {
	Id      int
	Image   string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

//店铺Logo
type StoreLogoImage struct {
	Id      int
	Image   string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

//店铺二维码
type StoreQRImage struct {
	Id      int
	Image   string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

//店铺banner
type StoreBannerImage struct {
	Id      int
	Image   string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

//店铺banner
type EvaluationImage struct {
	Id      int
	Image   string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
