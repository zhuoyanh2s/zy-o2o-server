package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
)

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	dbssl := beego.AppConfig.String("dbssl")
	dbprefix := beego.AppConfig.String("dbprefix")
	if dbport == "" {
		dbport = "5432"
	}
	//ds := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	ds := "host=" + dbhost + " port=" + dbport + " user=" + dbuser + " password=" + dbpassword + " dbname=" + dbname + " sslmode=" + dbssl

	//orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", ds)
	//orm.RegisterModel(new(User), new(Profile), new(Object))
	orm.RegisterModelWithPrefix(dbprefix, new(User), new(Profile), new(Object), new(Store), new(StoreLogoImage))
	orm.RunSyncdb("default", false, true)
	orm.DefaultTimeLoc = time.UTC

}

//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
