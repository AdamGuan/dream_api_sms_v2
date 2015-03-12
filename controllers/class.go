package controllers

import (
	"dream_api_sms_v2/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms_v2/helper"
	"github.com/astaxie/beego/config" 
)

//班级
type ClassController struct {
	beego.Controller
}

//json echo
func (u0 *ClassController) jsonEcho(datas map[string]interface{},u *ClassController) {
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
func (u0 *ClassController) checkSign(u *ClassController)int {
	result := -6
	pkg := u.Ctx.Request.Header.Get("pkg")
	sign := u.Ctx.Request.Header.Get("sign")
	mobilePhoneNumber := u.Ctx.Request.Header.Get("pnum")

	var userObj *models.MConsumer
	uid := userObj.GetUidByPhone(mobilePhoneNumber)

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

// @Title 添加一个班级
// @Description 添加一个班级(token: 登录时获取)
// @Param	className	form	string	true	班级名
// @Param	schoolId	form	int		true	学校ID
// @Param	gradeId		form	int		true	年级ID
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Param	pnum		header	string	true	手机号码
// @Success	200 {object} models.MClassAddResp
// @Failure 401 无权访问
// @router / [post]
func (u *ClassController) AddAClass() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var classObj *models.MClass
	//parse request parames
	u.Ctx.Request.ParseForm()
	className := u.Ctx.Request.FormValue("className")
	schoolId := u.Ctx.Request.FormValue("schoolId")
	gradeId := u.Ctx.Request.FormValue("gradeId")
	mobilePhoneNumber := u.Ctx.Request.Header.Get("pnum")

	//check sign
	datas["responseNo"] = u.checkSign(u)
	if datas["responseNo"] == 0 {
		var userObj *models.MConsumer
		uid := userObj.GetUidByPhone(mobilePhoneNumber)

		datas["responseNo"],datas["F_class_id"] = classObj.CreateAClass(uid,className,helper.StrToInt(schoolId),helper.StrToInt(gradeId))
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 获取某个学校下的所有班级信息
// @Description 获取某个学校下的所有班级信息(token: 登录时获取)
// @Param	schoolId	path	int		true	学校ID
// @Param	gradeId		query	int		true	年级ID
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Param	pnum		header	string	true	手机号码
// @Success	200 {object} models.MClassListInfoResp
// @Failure 401 无权访问
// @router /:schoolId [get]
func (u *ClassController) GetAllClasses() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var classObj *models.MClass
	//parse request parames
	u.Ctx.Request.ParseForm()
	schoolId := u.Ctx.Input.Param(":schoolId")
	gradeId := u.Ctx.Request.FormValue("gradeId")

	//check sign
	datas["responseNo"] = u.checkSign(u)
	if datas["responseNo"] == 0 {
		tmp := classObj.GetSchoolClassInfo(helper.StrToInt(schoolId),helper.StrToInt(gradeId))
		if len(tmp) > 0{
			datas["classList"] = tmp
		}else{
			datas["responseNo"] = -17
		}
	}
	//return
	u.jsonEcho(datas,u)
}