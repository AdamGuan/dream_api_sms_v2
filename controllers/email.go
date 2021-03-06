package controllers

import (
	"dream_api_sms_v2/helper"
	"dream_api_sms_v2/models"
)

//Email(每个用户email发送限制为1分钟的一次)
type EmailController struct {
	BaseController
}

// @Title email验证码验证
// @Description email验证码验证(token: md5(pkg))
// @Param	email				path	string	true	email
// @Param	num					form	string	true	验证码
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /emailvalid/:email [post]
func (u *EmailController) Emailvalid() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var emailObj *models.MEmail
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	email := u.Ctx.Input.Param(":email")
	num := u.Ctx.Request.FormValue("num")
	pkg := u.Ctx.Request.Header.Get("Pkg")
	//check sign
	datas["responseNo"] = u.checkSign2()
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckEmailValid(email) && len(num) > 0 {
		datas["responseNo"] = -1
		pkgConfig := pkgObj.GetPkgConfig(pkg)
		if len(pkgConfig) > 0 {
			res := emailObj.ValidEmail(pkg, num, email)
			if res {
				datas["responseNo"] = 0
			}
		}
	} else if datas["responseNo"] == 0 {
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas)
}

// @Title 发送email验证码(注册时)
// @Description 发送email验证码(注册时)(token: md5(pkg))
// @Param	email		path	string	true	email
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /register/:email [get]
func (u *EmailController) RegisterGetEmail() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var emailObj *models.MEmail
	var userObj *models.MConsumer
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	email := u.Ctx.Input.Param(":email")
	pkg := u.Ctx.Request.Header.Get("Pkg")
	//check sign
	datas["responseNo"] = u.checkSign2()
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckEmailValid(email) {
		datas["responseNo"] = -1
		res2 := userObj.CheckEmailValid(email)
		if res2 {
			pkgConfig := pkgObj.GetPkgConfig(pkg)
			if len(pkgConfig) > 0 && emailObj.CheckEmailRateValid(email, pkg) {
				emailObj.AddEmailRate(email, pkg)
				res := emailObj.GetEmailCode(email)
				if len(res) == 6 {
					emailObj.AddEmailRate(email, pkg)
					if emailObj.AddEmailActionvalid(email, pkg, res) {
						datas["responseNo"] = 0
					}
				} else {
					emailObj.DeleteEmailRate(email, pkg)
				}
			}
		} else {
			datas["responseNo"] = -2
		}
	} else if datas["responseNo"] == 0 {
		datas["responseNo"] = -1
	}

	//return
	u.jsonEcho(datas)
}

// @Title 发送email验证码(重置密码时)
// @Description 发送email验证码(重置密码时)(token: md5(pkg))
// @Param	email		path	string	true	email
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /resetpwd/:email [get]
func (u *EmailController) ResetPwdGetEmail() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var emailObj *models.MEmail
	var userObj *models.MConsumer
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	pkg := u.Ctx.Request.Header.Get("Pkg")
	email := u.Ctx.Input.Param(":email")
	//check sign
	datas["responseNo"] = u.checkSign2()
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckEmailValid(email) {
		datas["responseNo"] = -1
		res := userObj.CheckEmailExists(email)
		if res {
			pkgConfig := pkgObj.GetPkgConfig(pkg)
			if len(pkgConfig) > 0 && emailObj.CheckEmailRateValid(email, pkg) {
				emailObj.AddEmailRate(email, pkg)
				res := emailObj.GetEmailCode(email)
				if len(res) == 6 {
					emailObj.AddEmailRate(email, pkg)
					if emailObj.AddEmailActionvalid(email, pkg, res) {
						datas["responseNo"] = 0
					}
				} else {
					emailObj.DeleteEmailRate(email, pkg)
				}
			}
		} else {
			datas["responseNo"] = -4
		}
	} else if datas["responseNo"] == 0 {
		datas["responseNo"] = -4
	}

	//return
	u.jsonEcho(datas)
}

// @Title 发送email验证码(更换email)
// @Description 发送email验证码(更换email)(token: 登录时获取)
// @Param	email		query	string	true	email(新的)
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Param	huid		header	string	true	uid
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /resetemail [get]
func (u *EmailController) ChangeEmailCode() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var emailObj *models.MEmail
	var userObj *models.MConsumer
	var pkgObj *models.MPkg
	//parse request parames
	u.Ctx.Request.ParseForm()
	pkg := u.Ctx.Request.Header.Get("Pkg")
	email := u.Ctx.Request.FormValue("email")
	//check sign
	datas["responseNo"] = u.checkSign()
	//检查新email是否已被使用
	if datas["responseNo"] == 0 {
		if userObj.CheckEmailExists(email) {
			datas["responseNo"] = -26
		}
	}
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckEmailValid(email) {
		datas["responseNo"] = -1
		pkgConfig := pkgObj.GetPkgConfig(pkg)
		if len(pkgConfig) > 0 && emailObj.CheckEmailRateValid(email, pkg) {
			emailObj.AddEmailRate(email, pkg)
			res := emailObj.GetEmailCode(email)
			if len(res) == 6 {
				emailObj.AddEmailRate(email, pkg)
				if emailObj.AddEmailActionvalid(email, pkg, res) {
					datas["responseNo"] = 0
				}
			} else {
				emailObj.DeleteEmailRate(email, pkg)
			}
		}
	} else if datas["responseNo"] == 0 {
		datas["responseNo"] = -10
	}

	//return
	u.jsonEcho(datas)
}
