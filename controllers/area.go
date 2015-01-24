package controllers

import (
	"dream_api_sms_v2/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms_v2/helper"
)

//地域
type AreaController struct {
	beego.Controller
}

//json echo
func (u0 *AreaController) jsonEcho(datas map[string]interface{},u *AreaController) {
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
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
	u.jsonEcho(datas,u)
}

// @Title 获取市
// @Description 获取市
// @Param	province	query	string	true	省份
// @Success	200 {object} models.MAreaResp
// @Failure 401 无权访问
// @router /citys [get]
func (u *AreaController) GetCitys() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//parse request parames
	u.Ctx.Request.ParseForm()
	province := u.Ctx.Request.FormValue("province")
	//model ini
	var areaObj *models.MArea
	citys := areaObj.GetCitys(province)
	datas["areaList"] = citys
	//return
	u.jsonEcho(datas,u)
}

// @Title 获取县
// @Description 获取县
// @Param	city	query	string	true	市
// @Success	200 {object} models.MAreaResp
// @Failure 401 无权访问
// @router /countys [get]
func (u *AreaController) GetCountys() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//parse request parames
	u.Ctx.Request.ParseForm()
	city := u.Ctx.Request.FormValue("city")
	//model ini
	var areaObj *models.MArea
	countys := areaObj.GetCountys(city)
	datas["areaList"] = countys
	//return
	u.jsonEcho(datas,u)
}

// @Title 获取镇
// @Description 获取镇
// @Param	county	query	string	true	县
// @Success	200 {object} models.MAreaResp
// @Failure 401 无权访问
// @router /towns [get]
func (u *AreaController) GetTowns() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//parse request parames
	u.Ctx.Request.ParseForm()
	county := u.Ctx.Request.FormValue("county")
	//model ini
	var areaObj *models.MArea
	towns := areaObj.GetTowns(county)
	datas["areaList"] = towns
	//return
	u.jsonEcho(datas,u)
}