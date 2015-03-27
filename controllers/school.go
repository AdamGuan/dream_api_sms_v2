package controllers

import (
	"dream_api_sms_v2/models"
	"dream_api_sms_v2/helper"
)

//学校
type SchoolController struct {
	BaseController
}

// @Title 查询学校
// @Description 查询学校
// @Param	name	query	string	false	学校名(如果不传递，则查询结果中的全部学校)
// @Param	type	query	string	false	学校类型(值：[小学|初中|高中] ,如果不传递,则查询小初高全部)
// @Param	areaId	query	int		false	地域ID(只支持第三级的地域ID,如果不传递则查询全部地域的)
// @Param	areaName	query	string		false	地域名(只支持第三级的地域名,不支持模糊查找,如果不传递则查询全部地域的)
// @Param	needDefault	query	bool		false	缺省学校名(在没有查找到学校的情况下返回缺省值,为false 或是不传递则不返回缺省值)
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
	needDefault := u.Ctx.Request.FormValue("needDefault")
	needDefault2 := false
	if needDefault == "true"{
		needDefault2 = true
	}
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
		schools := schoolObj.QuerySchools(name,stype2,areaId2,areaName,needDefault2)
		if len(schools) > 0{
			datas["schoolList"] = schools
		}else{
			datas["responseNo"] = -17
		}
	}else{
		datas["responseNo"] = -10
	}
	//return
	u.jsonEcho(datas)
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
	u.jsonEcho(datas)
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
				datas["areaInfoList"] = tmp
			}else{
				datas["responseNo"] = -17
			}
		}else{
			datas["responseNo"] = -15
		}
	}
	//return
	u.jsonEcho(datas)
}

// @Title 根据学校ID查询学校名
// @Description 根据学校ID查询学校名
// @Param	schoolIds	path	string		true	学校ID(多个用","分隔,最多20个ID)
// @Success	200 {object} models.MSchoolResp
// @Failure 401 无权访问
// @router /name/:schoolIds [get]
func (u *SchoolController) GetSchoolName() {
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
			tmp := schoolObj.GetSchoolNameById(schoolIdList)
			if len(tmp) > 0{
				datas["schoolList"] = tmp
			}else{
				datas["responseNo"] = -17
			}
		}else{
			datas["responseNo"] = -15
		}
	}
	//return
	u.jsonEcho(datas)
}