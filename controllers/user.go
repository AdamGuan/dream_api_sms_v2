package controllers

import (
	"dream_api_sms_v2/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms_v2/helper"
	//"fmt"
	//"strings"
)

//用户
type UserController struct {
	beego.Controller
}

//json echo
func (u0 *UserController) jsonEcho(datas map[string]interface{},u *UserController) {
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}

//sign check
func (u0 *UserController) checkSign(u *UserController)int {
	result := -6
	pkg := u.Ctx.Request.Header.Get("pkg")
	sign := u.Ctx.Request.Header.Get("sign")
	mobilePhoneNumber := u.Ctx.Request.Header.Get("pnum")
	var pkgObj *models.MPkg
	if !pkgObj.CheckPkgExists(pkg){
		result = -7
	}else{
		var signObj *models.MSign
		if re := signObj.CheckSign(sign, mobilePhoneNumber, pkg,""); re == true {
			result = 0
		}
	}
	return result
}

//sign check, , token为包名的md5值
func (u0 *UserController) checkSign2(u *UserController)int {
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

// @Title 注册
// @Description 注册(token: md5(pkg))
// @Param	mobilePhoneNumber	form	string	true	手机号码
// @Param	pwd			form	string	true	密码
// @Param	num			form	string	true	验证码(经过验证成功后的)
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /register [post]
func (u *UserController) Register() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	var smsObj *models.MSms
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Request.FormValue("mobilePhoneNumber")
	pwd := u.Ctx.Request.FormValue("pwd")
	num := u.Ctx.Request.FormValue("num")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//check sign
	datas["responseNo"] = u.checkSign2(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && helper.CheckPwdValid(pwd) {
		datas["responseNo"] = -1
		if smsObj.CheckMsmActionvalid(mobilePhoneNumber,pkg,num) == true{
			res2 := userObj.AddUser(mobilePhoneNumber,pwd)
			datas["responseNo"] = res2
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 重置密码
// @Description 重置密码(token: md5(pkg))
// @Param	mobilePhoneNumber	form	string	true	手机号码
// @Param	pwd			form	string	true	密码
// @Param	num			form	string	true	验证码(经过验证成功后的)
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /resetpwd [put]
func (u *UserController) ResetPwd() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	var smsObj *models.MSms
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Request.FormValue("mobilePhoneNumber")
	pwd := u.Ctx.Request.FormValue("pwd")
	num := u.Ctx.Request.FormValue("num")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//check sign
	datas["responseNo"] = u.checkSign2(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && helper.CheckPwdValid(pwd) {
		datas["responseNo"] = -1
		if smsObj.CheckMsmActionvalid(mobilePhoneNumber,pkg,num) == true{
			res2 := userObj.ModifyUserPwd(mobilePhoneNumber,pwd)
			datas["responseNo"] = res2
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 登录
// @Description 登录(token: md5(pkg))
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	pwd			query	string	true	密码
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MUserLoginResp
// @Failure 401 无权访问
// @router /login/:mobilePhoneNumber [get]
func (u *UserController) CheckUserAndPwd() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pwd := u.Ctx.Request.FormValue("pwd")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//check sign
	datas["responseNo"] = u.checkSign2(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && helper.CheckPwdValid(pwd) {
		datas["responseNo"] = -1
		if !userObj.CheckUserNameExists(mobilePhoneNumber){
			datas["responseNo"] = -4
		}else{
			res := userObj.CheckUserAndPwd(mobilePhoneNumber,pwd)
			if res{
				token,tokenExpireDatetime := userObj.GetToken(mobilePhoneNumber,pkg)
				if len(token) > 0{
					datas["responseNo"] = 0
					datas["token"] = token
					datas["tokenExpireDatetime"] = tokenExpireDatetime
					//获取用户信息
					info := userObj.GetUserInfo(mobilePhoneNumber)
					if len(info) > 0{
						for k,v := range info{
							datas[k] = v
						}
					}
				}
			}else{
				datas["responseNo"] = -9
			}
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -5
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 找回密码
// @Description 找回密码(token: md5(pkg))
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	num			query	string	true	验证码(经过验证成功后的)
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MFindPwdResp
// @Failure 401 无权访问
// @router /pwd/:mobilePhoneNumber [get]
func (u *UserController) FindPwd() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	var smsObj *models.MSms
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	num := u.Ctx.Request.FormValue("num")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//check sign
	datas["responseNo"] = u.checkSign2(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		datas["responseNo"] = -1
		if userObj.CheckUserNameExists(mobilePhoneNumber){
			if smsObj.CheckMsmActionvalid(mobilePhoneNumber,pkg,num) == true{
				res := userObj.GetUserPwd(mobilePhoneNumber)
				if len(res) > 0{
					datas["responseNo"] = 0
					datas["password"] = res
				}
			}
		}else{
			datas["responseNo"] = -4
		}
		
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 修改密码
// @Description 修改密码(token: 登录时获取)
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	oldPwd			form	string	true	旧密码
// @Param	newPwd			form	string	true	新密码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Param	pnum		header	string	true	手机号码
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /pwd/:mobilePhoneNumber [put]
func (u *UserController) ModifyPwd() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	oldPwd := u.Ctx.Request.FormValue("oldPwd")
	newPwd := u.Ctx.Request.FormValue("newPwd")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && helper.CheckPwdValid(oldPwd) && helper.CheckPwdValid(newPwd) {
		datas["responseNo"] = -1
		if userObj.CheckUserAndPwd(mobilePhoneNumber,oldPwd){
			res2 := userObj.ModifyUserPwd(mobilePhoneNumber,newPwd)
			datas["responseNo"] = res2
		}else{
			datas["responseNo"] = -8
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 验证手机号码是否已注册
// @Description 验证手机号码是否已注册(token: md5(pkg))
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /exists/:mobilePhoneNumber [get]
func (u *UserController) CheckUserExists() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	//check sign
	datas["responseNo"] = u.checkSign2(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		if userObj.CheckUserNameExists(mobilePhoneNumber){
			datas["responseNo"] = -2
		}else{
			datas["responseNo"] = -4
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 获取用户信息
// @Description 获取用户信息(token: 登录时获取)
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Param	pnum		header	string	true	手机号码
// @Success	200 {object} models.MUserInfoResp
// @Failure 401 无权访问
// @router /:mobilePhoneNumber [get]
func (u *UserController) GetUserInfo() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		datas["responseNo"] = -1
		info := userObj.GetUserInfo(mobilePhoneNumber)
		if len(info) > 0{
			datas["responseNo"] = 0
			for k,v := range info{
				datas[k] = v
			}
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 修改用户信息
// @Description 修改用户信息(token: 登录时获取)
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	gender				form	string	false	性别(男或女)
// @Param	grade				form	string	false	年级
// @Param	birthday			form	string	false	生日(格式:1999-09-10)
// @Param	school				form	string	false	学校
// @Param	province			form	string	false	省
// @Param	city				form	string	false	市
// @Param	county				form	string	false	县
// @Param	area				form	string	false	区
// @Param	realname			form	string	false	真实姓名
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Param	pnum				header	string	true	手机号码
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /:mobilePhoneNumber [put]
func (u *UserController) ModifyUserInfo() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		datas["responseNo"] = -1
		parames := make(map[string]string)
		for k,v := range u.Ctx.Request.PostForm{
			parames[k] = v[0]
		}
		parames["mobilePhoneNumber"] = mobilePhoneNumber
		datas["responseNo"] = userObj.ModifyUserInfo(parames)
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 用户登出
// @Description 用户登出(token: 登录时获取)
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Param	pnum				header	string	true	手机号码
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /logout/:mobilePhoneNumber [delete]
func (u *UserController) UserLogout() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		datas["responseNo"] = -1
		if userObj.UserLoginout(mobilePhoneNumber,pkg) == true{
			datas["responseNo"] = 0
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}