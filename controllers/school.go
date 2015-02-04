package controllers

import (
	"dream_api_sms_v2/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms_v2/helper"
	"github.com/astaxie/beego/config" 
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
	
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug,_ := appConf.Bool(beego.RunMode+"::debug")
	if debug{
		datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	}

	u.Data["json"] = datas
	u.ServeJson()
}

// @Title 查询学校
// @Description 查询学校
// @Param	name	query	string	false	学校名(如果不传递，则查询结果中的全部学校)
// @Param	type	query	string	false	学校类型(值：[小学|初中|高中] ,如果不传递,则查询小初高全部)
// @Param	areaId	query	int		false	地域ID(只支持第三级的地域ID,如果不传递则查询全部地域的)
// @Param	areaName	query	string		false	地域名(只支持第三级的地域名,不支持模糊查找,如果不传递则查询全部地域的)
// @Success	200 {object} models.MSchoolResp
// @Failure 401 无权访问
// @router / [get]
func (u *SchoolController) QuerySchools() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//parse request parames
	u.Ctx.Request.ParseForm()
	name := u.Ctx.Request.FormValue("name")
	stype := u.Ctx.Request.FormValue("type")
	areaId := u.Ctx.Request.FormValue("areaId")
	areaName := u.Ctx.Request.FormValue("areaName")
	if len(name) > 0 || len(stype) > 0 || len(areaId) > 0 || len(areaName) > 0{
		//model ini
		stype2 := 0
		for k,v := range models.SchoolType{
			if stype == v{
				stype2 = k
			}
		}
		areaId2 := 0
		if len(areaId) > 0{
			areaId2 = helper.StrToInt(areaId)
		}
		var schoolObj *models.MSchool
		schools := schoolObj.QuerySchools(name,stype2,areaId2,areaName)
		datas["schoolList"] = schools
	}else{
		datas["responseNo"] = -10
	}
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

// @Title 根据学校查询所在地域
// @Description 根据学校查询所在地域
// @Param	schoolIds	path	string		true	学校ID(多个用","分隔,最多20个ID)
// @Success	200 {object} models.MSchoolAreaInfoResp
// @Failure 401 无权访问
// @router /area/:schoolIds [get]
func (u *SchoolController) GetSchoolArea() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var schoolObj *models.MSchool
	//parse request parames
	u.Ctx.Request.ParseForm()
	schoolIds := u.Ctx.Input.Param(":schoolIds")

	//check sign
	if datas["responseNo"] == 0 && len(schoolIds) > 0{
		schoolIdList := helper.Split(schoolIds,",")
		if len(schoolIdList) <= 20{
			tmp := make(models.MSchoolAreaInfoItemResp,len(schoolIdList))
			for _,schoolId := range schoolIdList{
				tmp2 := schoolObj.GetSchoolArea(helper.StrToInt(schoolId))
				if tmp2.F_school_id == helper.StrToInt(schoolId){
					tmp[schoolId] = schoolObj.GetSchoolArea(helper.StrToInt(schoolId))
				}
			}
			if len(tmp) > 0{
				datas["MSchoolAreaInfoResp"] = tmp
			}else{
				datas["MSchoolAreaInfoResp"] = make(models.MSchoolAreaInfoItemResp,1)
			}
		}else{
			datas["responseNo"] = -15
		}
	}
	//return
	u.jsonEcho(datas,u)
}