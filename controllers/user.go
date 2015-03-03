package controllers

import (
	"dream_api_sms_v2/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms_v2/helper"
	"github.com/astaxie/beego/config" 
//	"fmt"
	"strings"
	"os"
)

//用户
type UserController struct {
	beego.Controller
}

//上传文件用的size接口
type Sizer interface {
	Size() int64
}

//json echo
func (u0 *UserController) jsonEcho(datas map[string]interface{},u *UserController) {
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	
	datas["responseMsg"] = ""
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug,_ := appConf.Bool(beego.RunMode+"::debug")
	if debug{
		datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	}

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
// @Param	pwd					form	string	true	密码
// @Param	gender				form	string	false	性别(值: [男|女])
// @Param	grade				form	string	false	年级(小学一年级 -> 高中三年级)
// @Param	birthday			form	string	false	生日(格式:1999-09-10)
// @Param	school				form	int		false	学校ID
// @Param	province			form	int		false	省ID
// @Param	city				form	int		false	市ID
// @Param	county				form	int		false	县ID
// @Param	realname			form	string	false	真实姓名
// @Param	num					form	string	true	验证码(经过验证成功后的)
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
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
			parames := make(map[string]string)
			for k,v := range u.Ctx.Request.PostForm{
				parames[k] = v[0]
			}
			parames["mobilePhoneNumber"] = mobilePhoneNumber
			parames["pwd"] = pwd

			res2 := userObj.AddUser(parames)
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
					if len(info.F_phone_number) > 0{
						datas["F_phone_number"] = info.F_phone_number
						datas["F_gender"] = info.F_gender
						datas["F_grade"] = info.F_grade
						datas["F_grade_id"] = info.F_grade_id
						datas["F_birthday"] = info.F_birthday
						datas["F_school"] = info.F_school
						datas["F_school_id"] = info.F_school_id
						datas["F_province"] = info.F_province
						datas["F_province_id"] = info.F_province_id
						datas["F_city"] = info.F_city
						datas["F_city_id"] = info.F_city_id
						datas["F_county"] = info.F_county
						datas["F_county_id"] = info.F_county_id
						datas["F_user_realname"] = info.F_user_realname
						datas["F_user_nickname"] = info.F_user_nickname
						datas["F_crate_datetime"] = info.F_crate_datetime
						datas["F_modify_datetime"] = info.F_modify_datetime
						datas["F_class_id"] = info.F_class_id
						datas["F_class_name"] = info.F_class_name
						datas["F_avatar_url"] = info.F_avatar_url
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
		if len(info.F_phone_number) > 0{
			datas["responseNo"] = 0
			datas["F_phone_number"] = info.F_phone_number
			datas["F_gender"] = info.F_gender
			datas["F_grade"] = info.F_grade
			datas["F_grade_id"] = info.F_grade_id
			datas["F_birthday"] = info.F_birthday
			datas["F_school"] = info.F_school
			datas["F_school_id"] = info.F_school_id
			datas["F_province"] = info.F_province
			datas["F_province_id"] = info.F_province_id
			datas["F_city"] = info.F_city
			datas["F_city_id"] = info.F_city_id
			datas["F_county"] = info.F_county
			datas["F_county_id"] = info.F_county_id
			datas["F_user_realname"] = info.F_user_realname
			datas["F_user_nickname"] = info.F_user_nickname
			datas["F_crate_datetime"] = info.F_crate_datetime
			datas["F_modify_datetime"] = info.F_modify_datetime
			datas["F_class_id"] = info.F_class_id
			datas["F_class_name"] = info.F_class_name
			datas["F_avatar_url"] = info.F_avatar_url
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 修改用户信息
// @Description 修改用户信息(token: 登录时获取)(上传的头像name为"avatar",头像为用户上传的时候参数avatarType值为1)
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	gender				form	string	false	性别(值: [男|女])
// @Param	grade				form	string	false	年级(小学一年级 -> 高中三年级)
// @Param	birthday			form	string	false	生日(格式:1999-09-10)
// @Param	school				form	int		false	学校ID
// @Param	province			form	int		false	省ID
// @Param	city				form	int		false	市ID
// @Param	county				form	int		false	县ID
// @Param	realname			form	string	false	真实姓名
// @Param	nickname			form	string	false	昵称
// @Param	avatarType			form	int		false	头像类型(1:用户上传，2用户从系统头像选择)
// @Param	avatarId			form	int		false	系统头像ID(选择系统头像,参数avatarType为2)
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Param	pnum				header	string	true	手机号码
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /:mobilePhoneNumber [put]
func (u *UserController) ModifyUserInfo() {
	//uploadAvatar(u *UserController,mobilePhoneNumber string)
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	avatarType := helper.StrToInt(u.Ctx.Request.FormValue("avatarType"))
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		//头像修改
		avatarSysName := ""
		avatarExists := 0
		if avatarType == 1 || avatarType == 2{
			if avatarType == 1{	//上传
				datas["responseNo"] = u.uploadAvatar(u,mobilePhoneNumber)
			}else if avatarType == 2{	//系统选择
				avatarId := helper.StrToInt(u.Ctx.Request.FormValue("avatarId"))
				if avatarId <= 0{
					datas["responseNo"] = -10
				}else{
					//根据系统头像ID获取头像名称
					avatarSysName = userObj.GetAvatarNameFromId(avatarId)
					if len(avatarSysName) <= 0{
						datas["responseNo"] = -10
					}
				}
			}
			avatarExists = 1
		}
		//其它信息的修改
		if datas["responseNo"] == 0{
			datas["responseNo"] = -1
			parames := make(map[string]string)
			for k,v := range u.Ctx.Request.PostForm {
				if k != "avatarType" && k != "avatar" && k != "avatarId"{
					parames[k] = v[0]
				}
			}
			if len(avatarSysName) > 0{
				parames["avatarSysName"] = avatarSysName
			}
			if (avatarExists == 1 && len(parames) > 0) || (avatarExists != 1){
				parames["mobilePhoneNumber"] = mobilePhoneNumber
				datas["responseNo"] = userObj.ModifyUserInfo(parames)
			}else{
				datas["responseNo"] = 0
			}
		}
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

// @Title 修改用户的班级
// @Description 修改用户的班级(token: 登录时获取)
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	classId				query	int		true	班级ID
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Param	pnum				header	string	true	手机号码
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /class/:mobilePhoneNumber [put]
func (u *UserController) ModifyUserClass() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	classId := u.Ctx.Request.FormValue("classId")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 {
		datas["responseNo"] = userObj.UserChangeClass(mobilePhoneNumber,helper.StrToInt(classId))
	}
	//return
	u.jsonEcho(datas,u)
}


// @Title 上传用户头像
// @Description 上传用户头像(token: 登录时获取) (上传的头像name为"avatar")
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Param	pnum				header	string	true	手机号码
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /avatar/:mobilePhoneNumber [put]
func (u *UserController) UploadAvatar() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	if datas["responseNo"] == 0 {
		datas["responseNo"] = u.uploadAvatar(u,mobilePhoneNumber)
	}
	
	//return
	u.jsonEcho(datas,u)
}

// 上传用户头像
func (u0 *UserController) uploadAvatar(u *UserController,mobilePhoneNumber string) int{
	result := -1

	otherconf, _ := config.NewConfig("ini", "conf/other.conf")
	filename := otherconf.String("uploadAvatarFilename")
	allowType := otherconf.String("uploadAvatarType")
	savePath := otherconf.String("uploadAvatarSavePath")

	file,header,err := u.GetFile(filename)
	if err == nil {
		//文件类型
		contentType := header.Header.Get("Content-Type")
		typeList := strings.Split(allowType,",")
		valid := helper.StringInArray(contentType,typeList)
		if valid {
			contentType = strings.Replace(contentType,"image/","",-1)
			//文件大小
			if fileSizer, ok := file.(Sizer); ok {
				fileSize := fileSizer.Size()
				if fileSize <= 2*1024*1024{
					valid = true
				}else{
					valid = false
					result = -21
				}
			}
		}else{
			valid = false
			result = -22
		}

		//存储头像
		if valid{
			//文件存储
			saveFileName := "1_"+mobilePhoneNumber+"_"+helper.GetGuid()+"."+contentType
			saveFilePath := savePath+helper.Md5(saveFileName)[0:2]
			if !helper.Exist(saveFilePath){
				os.Mkdir(saveFilePath,0764)
				os.Create(saveFilePath+"/index.html")
			}
			err := u.SaveToFile(filename,saveFilePath+"/"+saveFileName)
			if err == nil{
				//数据库记录
				//model ini
				var userObj *models.MUser
				if userObj.UserAvatarNameModify(mobilePhoneNumber,saveFileName){
					result = 0
				}
			}
		}
	}
	
	return result
}

// @Title 获取服务端提供的头像
// @Description 获取服务端提供的头像
// @Success	200 {object} models.MAvatarlistResp
// @Failure 401 无权访问
// @router /avatarlist [get]
func (u *UserController) GetSystemAvatarList() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var userObj *models.MUser

	tmp := userObj.GetAvatarUrlList()
	if len(tmp) <= 0{
		datas["responseNo"] = -17
	}else{
		datas["avatarList"] = tmp
	}
	
	//return
	u.jsonEcho(datas,u)
}

// @Title 修改用户手机号码
// @Description 修改用户手机号码(token: 登录时获取)
// @Param	mobilePhoneNumber	path	string	true	手机号码(旧的手机号码)
// @Param	newPhone			form	string	true	手机号码(新的手机号码)
// @Param	num					form	string	true	验证码(经过验证成功后的)
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Param	pnum				header	string	true	手机号码(旧的手机号码)
// @Success	200 {object} models.MModifyPhoneResp
// @Failure 401 无权访问
// @router /phone/:mobilePhoneNumber [put]
func (u *UserController) ModifyPhone() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	var smsObj *models.MSms
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	newPhone := u.Ctx.Request.FormValue("newPhone")
	num := u.Ctx.Request.FormValue("num")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && helper.CheckMPhoneValid(newPhone) {
		datas["responseNo"] = -1
		if smsObj.CheckMsmActionvalid(newPhone,pkg,num) == true{
			datas["responseNo"] = userObj.ModifyUserPhone(mobilePhoneNumber,newPhone)
			if datas["responseNo"] == 0{
				//删除旧的手机号码的token
				var signObj *models.MSign
				signObj.DeleteAllPkgToken(mobilePhoneNumber)
				token,tokenExpireDatetime := userObj.GetToken(newPhone,pkg)
				datas["token"] = token
				datas["tokenExpireDatetime"] = tokenExpireDatetime
			}
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -10
	}
	//return
	u.jsonEcho(datas,u)
}