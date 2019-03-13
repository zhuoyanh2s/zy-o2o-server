package controllers

import (
	"zy-o2o-server/models"
	"zy-o2o-server/utils"
)

// Operations about Users
type UserController struct {
	baseController
}

//func (u *UserController) Post() {
//	var user models.User
//	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
//	username := models.AddUser(user)
//	u.Data["json"] = map[string]string{"username": username}
//	u.ServeJSON()
//}
//
//func (u *UserController) GetAll() {
//	users := models.GetAllUsers()
//	u.Data["json"] = users
//	u.ServeJSON()
//}
//
//func (u *UserController) Get() {
//	username := u.GetString("username")
//	if username != "" {
//		user, err := models.GetUser(username)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = user
//		}
//	}
//	u.ServeJSON()
//}
//
//func (u *UserController) Put() {
//	username := u.GetString("username")
//	if username != "" {
//		var user models.User
//		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
//		uu, err := models.UpdateUser(username, &user)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = uu
//		}
//	}
//	u.ServeJSON()
//}
//
//func (u *UserController) Delete() {
//	username := u.GetString("username")
//	models.DeleteUser(username)
//	u.Data["json"] = "delete success!"
//	u.ServeJSON()
//}

func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	//user, err := models.GetUser(username)
	//if err != nil {
	//	println(err)
	//	u.Data["error"] = "user not exist"
	//	return
	//}
	if models.Login(username, password) {
		u.Data["token"] = utils.ObtainJWT(username)
	} else {
		u.Data["error"] = "user not exist"
	}
	u.ServeJSON()
}

func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
