package controllers

import (
	"dream_api_sms_v2/helper"
	"dream_api_sms_v2/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	//	"github.com/astaxie/beego/utils"
)

type ErrorController struct {
	beego.Controller
}

func init() {
}

//json echo
func (u0 *ErrorController) jsonEcho(datas map[string]interface{}, u *ErrorController) {
	//	u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	//u.Ctx.Output.SetStatus(401)
	u.Ctx.Output.ContentType("application/json; charset=utf-8")

	datas["responseMsg"] = ""
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug, _ := appConf.Bool(beego.RunMode + "::debug")
	if debug {
		datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	}
	u.Data["json"] = datas
	u.ServeJson()
}

func (u *ErrorController) Error401() {
	//ini return
	datas := map[string]interface{}{"responseNo": 401}
	//log
	u.errerLog("401")
	//return
	u.jsonEcho(datas, u)
}

func (u *ErrorController) Error403() {
	//ini return
	datas := map[string]interface{}{"responseNo": 403}
	//log
	u.errerLog("403")
	//return
	u.jsonEcho(datas, u)
}

func (u *ErrorController) Error404() {
	//ini return
	datas := map[string]interface{}{"responseNo": 404}
	//log
	u.errerLog("404")
	//return
	u.jsonEcho(datas, u)
}

func (u *ErrorController) Error500() {
	//ini return
	datas := map[string]interface{}{"responseNo": 500}
	//log
	u.errerLog("500")
	//return
	u.jsonEcho(datas, u)
}

func (u *ErrorController) Error503() {
	//ini return
	datas := map[string]interface{}{"responseNo": 503}
	//log
	u.errerLog("503")
	//return
	u.jsonEcho(datas, u)
}

func (u *ErrorController) errerLog(code string) {
	var logObj models.MLog
	logObj.LogRequestErr500(u.Ctx, code)
}
