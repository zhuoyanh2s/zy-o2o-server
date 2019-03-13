package controllers

import "zy-o2o-server/utils"

type ObtainJWTController struct {
	baseController
}
type RefreshJWTController struct {
	baseController
}
type VerifyJWTController struct {
	baseController
}

//func (o *ObtainJWTController) ObtainJWT() {
//	username := o.GetString("username")
//
//	o.Data["token"] = utils.ObtainJWT()
//	o.ServeJSON()
//}
func (r *RefreshJWTController) RefreshJWT() {
	strToken := r.GetString("token")

	r.Data["token"] = utils.RefreshJWT(strToken)
	r.ServeJSON()
}
func (v *VerifyJWTController) VerifyJWT() {
	strToken := v.GetString("token")

	v.Data["token"] = utils.VerifyJWT(strToken)
	v.ServeJSON()
}
