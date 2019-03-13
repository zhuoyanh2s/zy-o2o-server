package routers

import (
	"github.com/astaxie/beego"
	"zy-o2o-server/controllers"
)

func init() {
	beego.Router("login/", &controllers.UserController{}, "*:Login")
	beego.Router("api-token-auth/", &controllers.UserController{}, "*:Login")
	beego.Router("api-token-refresh/", &controllers.RefreshJWTController{}, "*:RefreshJWT")
	beego.Router("api-token-verify/", &controllers.VerifyJWTController{}, "*:VerifyJWT")

	//beego.AutoRouter(&controllers.AdminController{})
}
