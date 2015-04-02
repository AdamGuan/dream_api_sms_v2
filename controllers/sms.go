package controllers

import (
	"dream_api_sms_v2/models"
	"dream_api_sms_v2/helper"
)

//短信(每个用户短信发送限制为1分钟的一次)
type SmsController struct {
	BaseController
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
	//log
	u.logRequest()
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
	datas["responseNo"] = u.checkSign2()
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && len(num) > 0 {
		datas["responseNo"] = -1
		pkgConfig := pkgObj.GetPkgConfig(pkg)
		if len(pkgConfig) > 0{
			res := smsObj.ValidMsm(pkg,num,mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"])
			if len(res) == 0{
				datas["responseNo"] = 0
			}
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas)
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
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var userObj *models.MConsumer
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pkg := u.Ctx.Request.Header.Get("Pkg")
	//check sign
	datas["responseNo"] = u.checkSign2()
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		datas["responseNo"] = -1
		res2 := userObj.CheckPhoneValid(mobilePhoneNumber)
		if res2 == 0{
			pkgConfig := pkgObj.GetPkgConfig(pkg)
			if len(pkgConfig) > 0 && smsObj.CheckMsmRateValid(mobilePhoneNumber,pkg){
				smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				res := smsObj.GetMsm(mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"],pkgConfig["F_app_name"],pkgConfig["F_app_msm_template"],pkg)
				if len(res) == 0{
					datas["responseNo"] = 0
					smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				}else{
					smsObj.DeleteMsmRate(mobilePhoneNumber,pkg)
				}
			}
		}else{
			if res2 == -23{
				res2 = -2
			}
			datas["responseNo"] = res2
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}

	//return
	u.jsonEcho(datas)
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
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var userObj *models.MConsumer
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pkg := u.Ctx.Request.Header.Get("Pkg")
	//check sign
	datas["responseNo"] = u.checkSign2()
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		datas["responseNo"] = -1
		res := userObj.CheckPhoneExists(mobilePhoneNumber)
		if res {
			pkgConfig := pkgObj.GetPkgConfig(pkg)
			if len(pkgConfig) > 0 && smsObj.CheckMsmRateValid(mobilePhoneNumber,pkg) {
				smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				res := smsObj.GetMsm(mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"],pkgConfig["F_app_name"],pkgConfig["F_app_msm_template"],pkg)
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
	u.jsonEcho(datas)
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
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var userObj *models.MConsumer
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pkg := u.Ctx.Request.Header.Get("Pkg")
	//check sign
	datas["responseNo"] = u.checkSign2()
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		datas["responseNo"] = -1
		res := userObj.CheckPhoneExists(mobilePhoneNumber)
		if res {
			pkgConfig := pkgObj.GetPkgConfig(pkg)
			if len(pkgConfig) > 0 && smsObj.CheckMsmRateValid(mobilePhoneNumber,pkg) {
				smsObj.AddMsmRate(mobilePhoneNumber,pkg)
				res := smsObj.GetMsm(mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"],pkgConfig["F_app_name"],pkgConfig["F_app_msm_template"],pkg)
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
	u.jsonEcho(datas)
}

// @Title 发送一条短信验证码(更换手机号码)(请使用下面的api代替,留这个api兼容之前的调用)
// @Description 发送一条短信验证码(更换手机号码)(token: 登录时获取)(请使用下面的api代替,留这个api兼容之前的调用)
// @Param	mobilePhoneNumber	path	string	true	手机号码(旧的号码)
// @Param	newPhone			query	string	true	手机号码(新的号码)
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Param	pnum				header	string	true	手机号码(旧的号码)
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /phone/:mobilePhoneNumber [get]
func (u *SmsController) ChangePhoneSms() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	pkg := u.Ctx.Request.Header.Get("Pkg")
	newPhone := u.Ctx.Request.FormValue("newPhone")
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pnum := u.Ctx.Request.Header.Get("pnum")
	//check sign
	datas["responseNo"] = u.checkSign3()
	//确定旧的手机号码是否是自己的
	if pnum != mobilePhoneNumber{
		datas["responseNo"] = -1
	}
	//检查新手机号码是否已被使用
	if datas["responseNo"] == 0{
		var userObj *models.MConsumer
		if userObj.CheckPhoneExists(newPhone){
			datas["responseNo"] = -23
		}
	}
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(newPhone) {
		datas["responseNo"] = -1
		pkgConfig := pkgObj.GetPkgConfig(pkg)
		if len(pkgConfig) > 0 && smsObj.CheckMsmRateValid(newPhone,pkg) {
			smsObj.AddMsmRate(newPhone,pkg)
			res := smsObj.GetMsm(newPhone,pkgConfig["F_app_id"],pkgConfig["F_app_key"],pkgConfig["F_app_name"],pkgConfig["F_app_msm_template"],pkg)
			if len(res) == 0{
				datas["responseNo"] = 0
				smsObj.AddMsmRate(newPhone,pkg)
			}else{
				smsObj.DeleteMsmRate(newPhone,pkg)
			}
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -10
	}

	//return
	u.jsonEcho(datas)
}

// @Title 发送一条短信验证码(更换手机号码)
// @Description 发送一条短信验证码(更换手机号码)(token: 登录时获取)
// @Param	mobilePhoneNumber	path	string	true	手机号码(新的号码)
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Param	huid				header	string	true	uid
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /mphone/:mobilePhoneNumber [get]
func (u *SmsController) ChangePhoneSms2() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	pkg := u.Ctx.Request.Header.Get("Pkg")
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	//check sign
	datas["responseNo"] = u.checkSign()
	//检查新手机号码是否已被使用
	if datas["responseNo"] == 0{
		var userObj *models.MConsumer
		if userObj.CheckPhoneExists(mobilePhoneNumber){
			datas["responseNo"] = -23
		}
	}
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		datas["responseNo"] = -1
		pkgConfig := pkgObj.GetPkgConfig(pkg)
		if len(pkgConfig) > 0 && smsObj.CheckMsmRateValid(mobilePhoneNumber,pkg) {
			smsObj.AddMsmRate(mobilePhoneNumber,pkg)
			res := smsObj.GetMsm(mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"],pkgConfig["F_app_name"],pkgConfig["F_app_msm_template"],pkg)
			if len(res) == 0{
				datas["responseNo"] = 0
				smsObj.AddMsmRate(mobilePhoneNumber,pkg)
			}else{
				smsObj.DeleteMsmRate(mobilePhoneNumber,pkg)
			}
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -10
	}

	//return
	u.jsonEcho(datas)
}