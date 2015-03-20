package controllers

import (
	"dream_api_sms_v2/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms_v2/helper"
	"github.com/astaxie/beego/config" 
)

//qq(第三方)
type QqController struct {
	beego.Controller
}

//json echo
func (u0 *QqController) jsonEcho(datas map[string]interface{},u *QqController) {
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
	//log
	u.logEcho(datas)

	u.ServeJson()
}

//sign check, , token为包名的md5值
func (u0 *QqController) checkSign(u *QqController)int {
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

// @Title 登录
// @Description 登录(token: md5(pkg))
// @Param	qq				path	string	true	qq号码
// @Param	access_token	query	string	true	access_token
// @Param	appid			query	string	true	appid(oauth_consumer_key)
// @Param	openid			query	string	true	openid
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MUserLoginResp
// @Failure 401 无权访问
// @router /login/:qq [get]
func (u *QqController) LoginQQ() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var userObj *models.MConsumer
	//parse request parames
	u.Ctx.Request.ParseForm()
	qq := u.Ctx.Input.Param(":qq")
	access_token := u.Ctx.Request.FormValue("access_token")
	appid := u.Ctx.Request.FormValue("appid")
	openid := u.Ctx.Request.FormValue("openid")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 {
		datas["responseNo"] = -1
		//检查qq信息的有效性
		if len(access_token) > 0 && len(appid) > 0 && len(openid) > 0 && len(qq) > 0{
			//检查qq号码是否已存在
			uid := userObj.GetUidByQQ(qq)
			if len(uid) <= 0{
				//写入一条qq数据
				uid = userObj.InsertQQ(qq)
			}
			if len(uid) > 0{
				//返回登录信息
				info := u.login(uid,pkg)
				if len(info) > 0{
					datas["responseNo"] = 0
					for key,value := range info{
						datas[key] = value
					}
				}
			}
		}else{
			datas["responseNo"] = -10
		}
	}
	//return
	u.jsonEcho(datas,u)
}

//登录
func (u *QqController) login(uid string,pkg string)map[string]interface{} {
	userInfo := map[string]interface{}{}
	//model ini
	var userObj *models.MConsumer
	//检查uid是否存在
	if userObj.CheckUserIdExists(uid){
		//获取token
		token,tokenExpireDatetime := userObj.GetTokenByUid(uid,pkg)
		//获取其它信息
		if len(token) > 0{
			userInfo["token"] = token
			userInfo["tokenExpireDatetime"] = tokenExpireDatetime
			info := userObj.GetUserInfoByUid(uid)
			if len(info.F_uid) > 0{
				userInfo["F_uid"] = info.F_uid
				userInfo["F_phone_number"] = info.F_phone_number
				userInfo["F_gender"] = info.F_gender
				userInfo["F_grade"] = info.F_grade
				userInfo["F_grade_id"] = info.F_grade_id
				userInfo["F_birthday"] = info.F_birthday
				userInfo["F_school"] = info.F_school
				userInfo["F_school_id"] = info.F_school_id
				userInfo["F_province"] = info.F_province
				userInfo["F_province_id"] = info.F_province_id
				userInfo["F_city"] = info.F_city
				userInfo["F_city_id"] = info.F_city_id
				userInfo["F_county"] = info.F_county
				userInfo["F_county_id"] = info.F_county_id
				userInfo["F_user_realname"] = info.F_user_realname
				userInfo["F_user_nickname"] = info.F_user_nickname
				userInfo["F_crate_datetime"] = info.F_crate_datetime
				userInfo["F_modify_datetime"] = info.F_modify_datetime
				userInfo["F_class_id"] = info.F_class_id
				userInfo["F_class_name"] = info.F_class_name
				userInfo["F_avatar_url"] = info.F_avatar_url
				userInfo["F_user_email"] = info.F_user_email
			}
		}
	}
	return userInfo
}

//记录请求
func (u *QqController) logRequest() {
	var logObj *models.MLog
	logObj.LogRequest(u.Ctx)
}

//记录返回
func (u *QqController) logEcho(datas map[string]interface{}) {
	var logObj *models.MLog
	logObj.LogEcho(datas)
}