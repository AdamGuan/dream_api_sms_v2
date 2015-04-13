package controllers

import (
	"dream_api_sms_v2/models"
)

//token
type TokenController struct {
	BaseController
}

// @Title 检查token是否正确
// @Description 检查token是否正确
// @Param	token				path	string	true	token
// @Param	mobilePhoneNumber	query	string	true	手机号码
// @Param	pkg					header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /verify/:token [get]
func (u *TokenController) CheckToken() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var signObj *models.MSign
	//parse request parames
	u.Ctx.Request.ParseForm()
	token := u.Ctx.Input.Param(":token")
	mobilePhoneNumber := u.Ctx.Request.FormValue("mobilePhoneNumber")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//检查参数
	var userObj *models.MConsumer
	uid := userObj.GetUidByPhone(mobilePhoneNumber)
	result := signObj.CheckToken(uid,pkg,token)
	if !result{
		datas["responseNo"] = -18
	}
	//return
	u.jsonEcho(datas)
}

// @Title 检查token是否正确(根据UID)
// @Description 检查token是否正确(根据UID)
// @Param	token				path	string	true	token
// @Param	uid					query	string	true	uid
// @Param	pkg					header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /verify2/:token [get]
func (u *TokenController) CheckTokenByUid() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var signObj *models.MSign
	//parse request parames
	u.Ctx.Request.ParseForm()
	token := u.Ctx.Input.Param(":token")
	uid := u.Ctx.Request.FormValue("uid")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//检查参数
	result := signObj.CheckToken(uid,pkg,token)
	if !result{
		datas["responseNo"] = -18
	}
	//return
	u.jsonEcho(datas)
}