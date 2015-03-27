package controllers

import (
	"dream_api_sms_v2/models"
	"dream_api_sms_v2/helper"
)

//地域
type AreaController struct {
	BaseController
}

type areaInfoList map[string]struct{
		F_area_id int
		F_area_name string
}

type areaInfoItem struct{
		F_area_id int
		F_area_name string
}

// @Title 获取所有省份
// @Description 获取所有省份
// @Success	200 {object} models.MAreaResp
// @Failure 401 无权访问
// @router /provinces [get]
func (u *AreaController) GetAllProvinces() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var areaObj *models.MArea
	provinces := areaObj.GetAllProvinces()
	datas["areaList"] = provinces
	//return
	u.jsonEcho(datas)
}

// @Title 获取市
// @Description 获取市
// @Param	provinceId	query	int	true	省份ID
// @Success	200 {object} models.MAreaResp
// @Failure 401 无权访问
// @router /citys [get]
func (u *AreaController) GetCitys() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//parse request parames
	u.Ctx.Request.ParseForm()
	provinceId := u.Ctx.Request.FormValue("provinceId")
	//model ini
	var areaObj *models.MArea
	citys := areaObj.GetCitys(helper.StrToInt(provinceId))
	datas["areaList"] = citys
	//return
	u.jsonEcho(datas)
}

// @Title 获取县
// @Description 获取县
// @Param	cityId	query	int	true	市ID
// @Success	200 {object} models.MAreaResp
// @Failure 401 无权访问
// @router /countys [get]
func (u *AreaController) GetCountys() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//parse request parames
	u.Ctx.Request.ParseForm()
	cityId := u.Ctx.Request.FormValue("cityId")
	//model ini
	var areaObj *models.MArea
	countys := areaObj.GetCountys(helper.StrToInt(cityId))
	datas["areaList"] = countys
	//return
	u.jsonEcho(datas)
}

// @Title 根据省份ID获取名称
// @Description 根据省份ID获取名称
// @Param	ids	path	string		true	省份ID(多个ID用","分隔)
// @Success	200 {object} models.MAreaInfoResp
// @Failure 401 无权访问
// @router /province/:ids [get]
func (u *AreaController) GetProvinceName() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//parse request parames
	u.Ctx.Request.ParseForm()
	ids := u.Ctx.Input.Param(":ids")
	if len(ids) > 0{
		idList := helper.Split(ids,",")
		tmp := make(areaInfoList,len(idList))
		for _,id := range idList{
			name,ok := models.Province[id]
			if ok{
				tmp[id] = areaInfoItem{F_area_id:helper.StrToInt(id),F_area_name:name}
			}
		}
		if len(tmp) > 0{
			datas["areaInfoList"] = tmp
		}else{
			datas["responseNo"] = -17
		}
	}
	//return
	u.jsonEcho(datas)
}

// @Title 根据市ID获取名称
// @Description 根据市ID获取名称
// @Param	ids	path	string		true	市ID(多个ID用","分隔)
// @Success	200 {object} models.MAreaInfoResp
// @Failure 401 无权访问
// @router /city/:ids [get]
func (u *AreaController) GetCityName() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//parse request parames
	u.Ctx.Request.ParseForm()
	ids := u.Ctx.Input.Param(":ids")
	if len(ids) > 0{
		idList := helper.Split(ids,",")
		tmp := make(areaInfoList,len(idList))
		for _,id := range idList{
			name,ok := models.City[id]
			if ok{
				tmp[id] = areaInfoItem{F_area_id:helper.StrToInt(id),F_area_name:name}
			}
		}
		if len(tmp) > 0{
			datas["areaInfoList"] = tmp
		}else{
			datas["responseNo"] = -17
		}
	}
	//return
	u.jsonEcho(datas)
}

// @Title 根据县ID获取名称
// @Description 根据县ID获取名称
// @Param	ids	path	string		true	县ID(多个ID用","分隔)
// @Success	200 {object} models.MAreaInfoResp
// @Failure 401 无权访问
// @router /county/:ids [get]
func (u *AreaController) GetCountyName() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//parse request parames
	u.Ctx.Request.ParseForm()
	ids := u.Ctx.Input.Param(":ids")
	if len(ids) > 0{
		idList := helper.Split(ids,",")
		tmp := make(areaInfoList,len(idList))
		for _,id := range idList{
			name,ok := models.County[id]
			if ok{
				tmp[id] = areaInfoItem{F_area_id:helper.StrToInt(id),F_area_name:name}
			}
		}
		if len(tmp) > 0{
			datas["areaInfoList"] = tmp
		}else{
			datas["responseNo"] = -17
		}
	}
	//return
	u.jsonEcho(datas)
}