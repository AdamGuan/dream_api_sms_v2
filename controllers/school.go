package controllers

import (
	"dream_api_sms_v2/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms_v2/helper"
)

//学校
type SchoolController struct {
	beego.Controller
}

//json echo
func (u0 *SchoolController) jsonEcho(datas map[string]interface{},u *SchoolController) {
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}

// @Title 根据关键字查询学校
// @Description 根据关键字查询学校
// @Param	name	query	string	true	学校名
// @Success	200 {object} models.MSchoolResp
// @Failure 401 无权访问
// @router / [get]
func (u *SchoolController) QuerySchools() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//parse request parames
	u.Ctx.Request.ParseForm()
	name := u.Ctx.Request.FormValue("name")
	//model ini
	var schoolObj *models.MSchool
	schools := schoolObj.QuerySchools(name)
	datas["schoolList"] = schools
	//return
	u.jsonEcho(datas,u)
}

// @Title 获取所有年级
// @Description 获取所有年级
// @Success	200 {object} models.MGradeResp
// @Failure 401 无权访问
// @router /grades [get]
func (u *SchoolController) GetAllGrade() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var schoolObj *models.MSchool
	grades := schoolObj.GetAllGrade()
	datas["gradeList"] = grades
	//return
	u.jsonEcho(datas,u)
}