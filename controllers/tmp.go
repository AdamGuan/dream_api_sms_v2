package controllers

import (
	"dream_api_sms_v2/models"
	"github.com/astaxie/beego"
	"dream_api_sms_v2/helper"
	"github.com/astaxie/beego/config" 
//	"fmt"
)

//临时工具
type TmpController struct {
	beego.Controller
}

//json echo
func (u0 *TmpController) jsonEcho(datas map[string]interface{},u *TmpController) {
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

/*
// @Title 清空全部用户数据(临时用)
// @Description 清空全部用户数据(临时用)
// @Param	wakaka			query	int	false	炸弹
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /alluser [delete]
func (u *TmpController) DeleteAllUser() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}

	//parse request parames
	u.Ctx.Request.ParseForm()
	wakaka := u.Ctx.Request.FormValue("wakaka")

	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err == nil{
		cwakaka,err := appConf.Int(beego.RunMode+"::wakaka")
		if err == nil && cwakaka == wakaka{
			//model ini
			var tmpObj *models.MTmp
			tmpObj.DeleteAllUser()
		}
	}

	//return
	u.jsonEcho(datas,u)
}
*/

// @Title 清空指定用户数据(临时用)
// @Description 清空指定用户数据(临时用)
// @Param	mobilePhoneNumber			query	string	true	手机号码
// @Param	wakaka			query	int	false	炸弹
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /user [delete]
func (u *TmpController) DeleteUser() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": 0}

	//parse request parames
	u.Ctx.Request.ParseForm()
	wakaka := helper.StrToInt(u.Ctx.Request.FormValue("wakaka"))

	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err == nil{
		cwakaka,err := appConf.Int(beego.RunMode+"::wakaka")
		if err == nil && cwakaka == wakaka{
			//
			mobilePhoneNumber := u.Ctx.Request.FormValue("mobilePhoneNumber")
			//model ini
			var userObj *models.MConsumer
			uid := userObj.GetUidByPhone(mobilePhoneNumber)

			var tmpObj *models.MTmp
			tmpObj.DeleteUser(uid)
		}
	}
	
	//return
	u.jsonEcho(datas,u)
}

/*
// @Title test
// @Description test
// @Param	phone			query	string	false	手机号码
// @Param	wakaka			query	int	false	炸弹
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /ttt [get]
	func (u *TmpController) Test() {
	
	//ini return
//	datas := map[string]interface{}{"responseNo": 0}

	u.Ctx.Request.ParseForm()
	phone := u.Ctx.Request.FormValue("phone")
	

	if phone == "abc"{
		c := 10
		b := 10
		d := c -b
		a := 10/d
		fmt.Println("a:",a)
	}
	
//	fmt.Println(u.Ctx.Output.IsServerError())
	//return
//	u.jsonEcho(datas,u)
	
}
*/

//记录请求
func (u *TmpController) logRequest() {
	var logObj *models.MLog
	logObj.LogRequest(u.Ctx)
}

//记录返回
func (u *TmpController) logEcho(datas map[string]interface{}) {
	var logObj *models.MLog
	logObj.LogEcho(datas)
}