package controllers

import (
	"dream_api_sms_v2/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms_v2/helper"
	"github.com/astaxie/beego/config"
)

//公共controller
type BaseController struct {
	beego.Controller
}

//json echo
func (u *BaseController) jsonEcho(datas map[string]interface{}) {
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
//		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
//		u.Ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
		u.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	} 
	
	datas["responseMsg"] = ""
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug,_ := appConf.Bool(beego.RunMode+"::debug")
	if debug{
		datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	}

	u.Data["json"] = datas
	//log
	u.logEcho(datas)

	u.ServeJson()
}

//sign check
func (u *BaseController) checkSign()int {
	result := -6
	pkg := u.Ctx.Request.Header.Get("pkg")
	sign := u.Ctx.Request.Header.Get("sign")
	uid := u.Ctx.Request.Header.Get("huid")
	var pkgObj *models.MPkg
	if !pkgObj.CheckPkgExists(pkg){
		result = -7
	}else{
		var signObj *models.MSign
		if re := signObj.CheckSign(sign, uid, pkg,""); re == true {
			result = 0
		}
	}
	return result
}

//sign check, , token为包名的md5值
func (u *BaseController) checkSign2()int {
	result := -6
	pkg := u.Ctx.Request.Header.Get("pkg")
	sign := u.Ctx.Request.Header.Get("sign")
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

//sign check3 pnum可以是手机号码或uid
func (u *BaseController) checkSign3()int {
	result := -6
	pkg := u.Ctx.Request.Header.Get("pkg")
	sign := u.Ctx.Request.Header.Get("sign")
	uid := u.Ctx.Request.Header.Get("pnum")
	//判断是否为手机号码
	if helper.CheckMPhoneValid(uid){
		var userObj *models.MConsumer
		uid = userObj.GetUidByPhone(uid)
	}
	var pkgObj *models.MPkg
	if !pkgObj.CheckPkgExists(pkg){
		result = -7
	}else{
		var signObj *models.MSign
		if re := signObj.CheckSign(sign, uid, pkg,""); re == true {
			result = 0
		}
	}
	return result
}

//记录请求
func (u *BaseController) logRequest() {
	var logObj *models.MLog
	logObj.LogRequest(u.Ctx)
}

//记录返回
func (u *BaseController) logEcho(datas map[string]interface{}) {
	var logObj *models.MLog
	logObj.LogEcho(datas)
}