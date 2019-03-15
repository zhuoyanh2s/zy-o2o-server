package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

var engine *xorm.Engine

func Setup() {
	var err error
	engine, err = xorm.NewEngine("postgres", viper.GetString("database.conn"))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	engine.Sync2(new(User))
}
