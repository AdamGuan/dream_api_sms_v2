package controllers

import (
	"dream_api_sms_v2/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms_v2/helper"
	//"fmt"
	//"strings"
)

//短信(每个用户短信发送限制为1分钟的一次)
type SmsController struct {
	beego.Controller
}

//json echo
func (u0 *SmsController) jsonEcho(datas map[string]interface{},u *SmsController) {
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}

//sign check, token为包名md5
func (u0 *SmsController) checkSign(u *SmsController)int {
	result := -6
	pkg := u.Ctx.Request.Header.Get("pkg")
	sign := u.Ctx.Request.Header.Get("sign")
//	mobilePhoneNumber := u.Ctx.Request.Header.Get("pnum")
	var pkgObj *models.MPkg
	if !pkgObj.CheckPkgExists(pkg){
		result = -7
	}else{
		var signObj *models.MSign
		if re := signObj.CheckSign(sign, "", pkg,helper.Md5(pkg)); re == true {
			result = 0
		}
	}
	return result
}

// @Title 短信验证码验证
// @Description 短信验证码验证(token: md5(pkg))
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	num					form	string	true	验证码
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /smsvalid/:mobilePhoneNumber [post]
func (u *SmsController) Smsvalid() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	num := u.Ctx.Request.FormValue("num")
	pkg := u.Ctx.Request.Header.Get("Pkg")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && len(num) > 0 {
		datas["responseNo"] = -1
		pkgConfig := pkgObj.GetPkgConfig(pkg)
		if len(pkgConfig) > 0{
			res := smsObj.ValidMsm(num,mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"])
			if len(res) == 0{
				datas["responseNo"] = 0
				smsObj.AddMsmActionvalid(mobilePhoneNumber,pkg,num)
			}
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 发送一条短信验证码(注册时)
// @Description 发送一条短信验证码(注册时)(token: md5(pkg))
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /register/:mobilePhoneNumber [get]
func (u *SmsController) RegisterGetSms() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var userObj *models.MUser
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pkg := u.Ctx.Request.Header.Get("Pkg")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		datas["responseNo"] = -1
		res2 := userObj.CheckUserNameValid(mobilePhoneNumber)
		if res2 == 0{
			pkgConfig := pkgObj.GetPkgConfig(pkg)
			if len(pkgConfig) > 0 && smsObj.CheckMsmRateValid(mobilePhoneNumber,pkg){
				smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				res := smsObj.GetMsm(mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"],pkgConfig["F_app_name"],pkgConfig["F_app_msm_template"])
				if len(res) == 0{
					datas["responseNo"] = 0
					smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				}else{
					smsObj.DeleteMsmRate(mobilePhoneNumber,pkg)
				}
			}
		}else{
			datas["responseNo"] = res2
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}

	//return
	u.jsonEcho(datas,u)
}

// @Title 发送一条短信验证码(重置密码时)
// @Description 发送一条短信验证码(重置密码时)(token: md5(pkg))
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /resetpwd/:mobilePhoneNumber [get]
func (u *SmsController) ResetPwdGetSms() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var userObj *models.MUser
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pkg := u.Ctx.Request.Header.Get("Pkg")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		datas["responseNo"] = -1
		res := userObj.CheckUserNameExists(mobilePhoneNumber)
		if res {
			pkgConfig := pkgObj.GetPkgConfig(pkg)
			if len(pkgConfig) > 0 && smsObj.CheckMsmRateValid(mobilePhoneNumber,pkg) {
				smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				res := smsObj.GetMsm(mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"],pkgConfig["F_app_name"],pkgConfig["F_app_msm_template"])
				if len(res) == 0{
					datas["responseNo"] = 0
					smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				}else{
					smsObj.DeleteMsmRate(mobilePhoneNumber,pkg)
				}
			}
		}else{
			datas["responseNo"] = -4
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -4
	}

	//return
	u.jsonEcho(datas,u)
}

// @Title 发送一条短信验证码(找回密码时)
// @Description 发送一条短信验证码(找回密码时)(token: md5(pkg))
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /pwd/:mobilePhoneNumber [get]
func (u *SmsController) FindPwdGetSms() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var userObj *models.MUser
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pkg := u.Ctx.Request.Header.Get("Pkg")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		datas["responseNo"] = -1
		res := userObj.CheckUserNameExists(mobilePhoneNumber)
		if res {
			pkgConfig := pkgObj.GetPkgConfig(pkg)
			if len(pkgConfig) > 0 && smsObj.CheckMsmRateValid(mobilePhoneNumber,pkg) {
				smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				res := smsObj.GetMsm(mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"],pkgConfig["F_app_name"],pkgConfig["F_app_msm_template"])
				if len(res) == 0{
					datas["responseNo"] = 0
					smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				}else{
					smsObj.DeleteMsmRate(mobilePhoneNumber,pkg)
				}
			}
		}else{
			datas["responseNo"] = -4
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -4
	}

	//return
	u.jsonEcho(datas,u)
}