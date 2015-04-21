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
	//	mobilePhoneNumber := u.Ctx.Request.FormValue("mobilePhoneNumber")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//检查参数
	//	var userObj *models.MConsumer
	//	uid := userObj.GetUidByPhone(mobilePhoneNumber)
	result, uid := signObj.CheckToken(pkg, token)
	if !result {
		datas["responseNo"] = -18
	} else {
		if u.checkInWhiteListIp() {
			datas["F_uid"] = uid
		}
	}
	//return
	u.jsonEcho(datas)
}

// @Title 检查token是否正确(请用上面的)
// @Description 检查token是否正确(请用上面的)
// @Param	token				path	string	true	token
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
	//	uid := u.Ctx.Request.FormValue("uid")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//检查参数
	result, uid := signObj.CheckToken(pkg, token)
	if !result {
		datas["responseNo"] = -18
	} else {
		if u.checkInWhiteListIp() {
			datas["F_uid"] = uid
		}
	}
	//return
	u.jsonEcho(datas)
}

//检查ip是否存在于白名单中
func (u *TokenController) checkInWhiteListIp() bool {
	var userObj *models.MConsumer
	return userObj.CheckUpdateCoinWhiteIp(u.Ctx.Input.IP())
}
