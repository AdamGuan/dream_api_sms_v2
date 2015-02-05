package controllers

import (
	"dream_api_sms_v2/models"
	"github.com/astaxie/beego"
	"dream_api_sms_v2/helper"
	"github.com/astaxie/beego/config" 
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
	u.ServeJson()
}

// @Title 清空全部用户数据(临时用)
// @Description 清空全部用户数据(临时用)
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /alluser [delete]
func (u *TmpController) DeleteAllUser() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}

	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug,_ := appConf.Bool(beego.RunMode+"::debug")
	if debug{
		//model ini
		var tmpObj *models.MTmp
		tmpObj.DeleteAllUser()
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 清空指定用户数据(临时用)
// @Description 清空指定用户数据(临时用)
// @Param	mobilePhoneNumber			query	string	true	手机号码
// @Param	wakaka			query	int	false	炸弹
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /user [delete]
func (u *TmpController) DeleteUser() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}

	//parse request parames
	u.Ctx.Request.ParseForm()
	wakaka := u.Ctx.Request.FormValue("wakaka")

	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug,_ := appConf.Bool(beego.RunMode+"::debug")
	if debug || wakaka == "-10"{
		mobilePhoneNumber := u.Ctx.Request.FormValue("mobilePhoneNumber")
		//model ini
		var tmpObj *models.MTmp
		tmpObj.DeleteUser(mobilePhoneNumber)
	}
	//return
	u.jsonEcho(datas,u)
}