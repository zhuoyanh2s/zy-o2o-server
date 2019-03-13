package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type baseController struct {
	beego.Controller
	o              orm.Ormer
	controllerName string
	actionName     string
}
