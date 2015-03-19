package controllers

import (
	"dream_api_sms_v2/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms_v2/helper"
	"github.com/astaxie/beego/config" 
	//"fmt"
	//"strings"
)

//token
type TokenController struct {
	beego.Controller
}

//json echo
func (u0 *TokenController) jsonEcho(datas map[string]interface{},u *TokenController) {
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

// @Title 检查token是否正确
// @Description 检查token是否正确
// @Param	mobilePhoneNumber	query	string	true	手机号码
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
	u.jsonEcho(datas,u)
}

//记录请求
func (u *TokenController) logRequest() {
	var logObj *models.MLog
	logObj.LogRequest(u.Ctx)
}

//记录返回
func (u *TokenController) logEcho(datas map[string]interface{}) {
	var logObj *models.MLog
	logObj.LogEcho(datas)
}