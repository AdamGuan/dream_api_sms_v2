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

//获取省份
func (u *MArea) GetAllProvinces()[]string{
	provinces := make([]string,0)
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_area_name FROM t_area WHERE F_area_level=1").Values(&maps)
	if err == nil && num > 0 {
		provinces2 := make([]string,num)
		for key,item := range maps{
			provinces2[key] = item["F_area_name"].(string)
		}
		return provinces2
	}
	return provinces
}

//获取市
func (u *MArea) GetCitys(provinceValue string)[]string{
	citys := make([]string,0)

	//获取province id
	provinceId := ""
	for k,v := range Province{
		if v == provinceValue{
			provinceId = k
		}
	}
	if len(provinceId) > 0{
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT F_area_name FROM t_area WHERE F_area_level=2 AND F_area_parent=?",helper.StrToInt(provinceId)).Values(&maps)
		if err == nil && num > 0 {
			citys2 := make([]string,num)
			for key,item := range maps{
				citys[key] = item["F_area_name"].(string)
			}
			return citys2
		}
	}
	return citys
}

//获取县
func (u *MArea) GetCountys(cityValue string)[]string{
	countys := make([]string,0)

	//获取city id
	cityId := ""
	for k,v := range City{
		if v == cityValue{
			cityId = k
		}
	}
	if len(cityId) > 0{
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT F_area_name FROM t_area WHERE F_area_level=3 AND F_area_parent=?",helper.StrToInt(cityId)).Values(&maps)
		if err == nil && num > 0 {
			countys2 := make([]string,num)
			for key,item := range maps{
				countys[key] = item["F_area_name"].(string)
			}
			return countys2
		}
	}
	return countys
}

//获取镇
func (u *MArea) GetTowns(countyValue string)[]string{
	towns := make([]string,0)

	//获取county id
	countyId := ""
	for k,v := range County{
		if v == countyValue{
			countyId = k
		}
	}
	if len(countyId) > 0{
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT F_area_name FROM t_area WHERE F_area_level=3 AND F_area_parent=?",helper.StrToInt(countyId)).Values(&maps)
		if err == nil && num > 0 {
			towns2 := make([]string,num)
			for key,item := range maps{
				towns[key] = item["F_area_name"].(string)
			}
			return towns2
		}
	}
	return towns
}