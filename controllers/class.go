package controllers

import (
	"dream_api_sms_v2/helper"
	"dream_api_sms_v2/models"
)

//班级
type ClassController struct {
	BaseController
}

// @Title 添加一个班级
// @Description 添加一个班级(token: 登录时获取)
// @Param	className	form	string	true	班级名
// @Param	schoolId	form	int		true	学校ID
// @Param	gradeId		form	int		true	年级ID
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Param	pnum		header	string	true	手机号码 或是 uid
// @Success	200 {object} models.MClassAddResp
// @Failure 401 无权访问
// @router / [post]
func (u *ClassController) AddAClass() {
	//log
	u.logRequest()
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
	datas["responseNo"] = u.checkSign3()
	if datas["responseNo"] == 0 {
		var userObj *models.MConsumer
		uid := userObj.GetUidByPhone(mobilePhoneNumber)

		datas["responseNo"], datas["F_class_id"] = classObj.CreateAClass(uid, className, helper.StrToInt(schoolId), helper.StrToInt(gradeId))
	}
	//return
	u.jsonEcho(datas)
}

// @Title 获取某个学校下的所有班级信息
// @Description 获取某个学校下的所有班级信息(token: 登录时获取)
// @Param	schoolId	path	int		true	学校ID
// @Param	gradeId		query	int		true	年级ID
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Param	pnum		header	string	true	手机号码 或是 uid
// @Success	200 {object} models.MClassListInfoResp
// @Failure 401 无权访问
// @router /:schoolId [get]
func (u *ClassController) GetAllClasses() {
	//log
	u.logRequest()

	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var classObj *models.MClass
	//parse request parames
	u.Ctx.Request.ParseForm()
	schoolId := u.Ctx.Input.Param(":schoolId")
	gradeId := u.Ctx.Request.FormValue("gradeId")

	//check sign
	datas["responseNo"] = u.checkSign3()
	if datas["responseNo"] == 0 {
		tmp := classObj.GetSchoolClassInfo(helper.StrToInt(schoolId), helper.StrToInt(gradeId))
		if len(tmp) > 0 {
			datas["classList"] = tmp
		} else {
			datas["responseNo"] = -17
		}
	}
	//return
	u.jsonEcho(datas)
}
