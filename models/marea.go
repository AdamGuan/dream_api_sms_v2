package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"dream_api_sms_v2/helper"
)

func init() {
}

type MArea struct {
}

type listItem struct{
	F_area_id	int
	F_area_name	string
}

type list []struct{
	F_area_id	int
	F_area_name	string
}


//获取省份
func (u *MArea) GetAllProvinces()list{
	provinces := make(list,0)
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_area_id,F_area_name FROM t_area WHERE F_area_level=1").Values(&maps)
	if err == nil && num > 0 {
		provinces2 := make(list,num)
		for key,item := range maps{
			provinces2[key] = listItem{F_area_id:helper.StrToInt(item["F_area_id"].(string)),F_area_name:item["F_area_name"].(string)}
		}
		return provinces2
	}
	return provinces
}

//获取市
func (u *MArea) GetCitys(provinceId int)list{
	citys := make(list,0)

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_area_id,F_area_name FROM t_area WHERE F_area_level=2 AND F_area_parent=?",provinceId).Values(&maps)
	if err == nil && num > 0 {
		citys2 := make(list,num)
		for key,item := range maps{
			citys2[key] = listItem{F_area_id:helper.StrToInt(item["F_area_id"].(string)),F_area_name:item["F_area_name"].(string)}
		}
		return citys2
	}

	return citys
}

//获取县
func (u *MArea) GetCountys(cityId int)list{
	countys := make(list,0)

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_area_id,F_area_name FROM t_area WHERE F_area_level=3 AND F_area_parent=?",cityId).Values(&maps)
	if err == nil && num > 0 {
		countys2 := make(list,num)
		for key,item := range maps{
			countys2[key] = listItem{F_area_id:helper.StrToInt(item["F_area_id"].(string)),F_area_name:item["F_area_name"].(string)}
		}
		return countys2
	}

	return countys
}

//获取镇
func (u *MArea) GetTowns(countyId int)list{
	towns := make(list,0)

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_area_id,F_area_name FROM t_area WHERE F_area_level=4 AND F_area_parent=?",countyId).Values(&maps)
	if err == nil && num > 0 {
		towns2 := make(list,num)
		for key,item := range maps{
			towns2[key] = listItem{F_area_id:helper.StrToInt(item["F_area_id"].(string)),F_area_name:item["F_area_name"].(string)}
		}
		return towns2
	}

	return towns
}