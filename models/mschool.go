package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"dream_api_sms_v2/helper"
)

func init() {
}

type MSchool struct {
}

type schoolList []struct {
	F_school_id int
	F_school string
	F_school_type string
}

type schoolItem struct {
	F_school_id int
	F_school string
	F_school_type string
}

type schoolArea struct{
	F_school_id int
	F_area_province_id int
	F_area_province_name string
	F_area_city_id int
	F_area_city_name string
	F_area_county_id int
	F_area_county_name string
}

//查询学校
func (u *MSchool) QuerySchools(name string,stype int,areaId int,areaName string,needDefault bool)schoolList{
	schools := make(schoolList,0)
	
	//查询地域ID
	o := orm.NewOrm()
	if len(areaName) > 0{
		var maps []orm.Params
		num, err := o.Raw("SELECT F_area_id FROM t_area WHERE F_area_level = 3 AND F_area_name = ? LIMIT 1",areaName).Values(&maps)
		if err == nil && num > 0 {
			areaId = helper.StrToInt(maps[0]["F_area_id"].(string))
		}
	}

	where := " WHERE 1 AND F_school_id > 0 "

	if stype != 0{
		where += " AND F_school_type = "+helper.IntToString(stype)
	}
	if areaId != 0{
		where += " AND F_belong_area_id = "+helper.IntToString(areaId)
	}
	if len(name) > 0{
		where += " AND F_school LIKE '%"+name+"%'"
	}

	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM t_school"+where+" LIMIT 100").Values(&maps)
	if err == nil && num > 0 {
		schools := make(schoolList,num)
		for key,item := range maps{
			schools[key] = schoolItem{F_school_id:helper.StrToInt(item["F_school_id"].(string)),F_school:item["F_school"].(string),F_school_type:SchoolType[helper.StrToInt(item["F_school_type"].(string))]}
		}
		return schools
	}
	if needDefault && num <= 0{
		schools := make(schoolList,0)
		for _,item := range DefaultSchoolList{
			if stype != 0{
				if item.F_school_type == stype{
					schools = append(schools, schoolItem{F_school_id:item.F_school_id,F_school:item.F_school,F_school_type:SchoolType[item.F_school_type]})
				}
			}else{
				schools = append(schools, schoolItem{F_school_id:item.F_school_id,F_school:item.F_school,F_school_type:SchoolType[item.F_school_type]})
			}
		}
		return schools
	}

	return schools
}

//获取所有年级
func (u *MSchool) GetAllGrade()[]string{
	grades := make([]string,0)
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_grade FROM t_grade WHERE 1").Values(&maps)
	if err == nil && num > 0 {
		grades := make([]string,num)
		for key,item := range maps{
			grades[key] = item["F_grade"].(string)
		}
		return grades
	}
	return grades
}

//根据某个学校查询所在地域
func (u *MSchool) GetSchoolArea(schoolId int)schoolArea{
	areaInfo := schoolArea{}
	areaInfoTmp := schoolArea{}
	//查学校所属地区ID
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_belong_area_id FROM t_school WHERE F_school_id = ? LIMIT 1",schoolId).Values(&maps)
	if err == nil && num > 0 {
		areaId := helper.StrToInt(maps[0]["F_belong_area_id"].(string))
		if schoolId < 0{
			//查学校所属省的ID,name
			var maps []orm.Params
			num, err := o.Raw("SELECT F_area_id,F_area_name FROM t_area WHERE F_area_level = 1 AND F_area_id = ? LIMIT 1",areaId).Values(&maps)
			if err == nil && num > 0 {
				areaInfoTmp.F_area_province_id = helper.StrToInt(maps[0]["F_area_id"].(string))
				areaInfoTmp.F_area_province_name = maps[0]["F_area_name"].(string)
				areaInfoTmp.F_school_id = schoolId

				areaInfo = areaInfoTmp
			}
		}else{
			if areaId > 0{
				//查学校所属县的ID,name
				var maps []orm.Params
				num, err := o.Raw("SELECT F_area_id,F_area_name,F_area_parent FROM t_area WHERE F_area_level = 3 AND F_area_id = ? LIMIT 1",areaId).Values(&maps)
				if err == nil && num > 0 {
					areaInfoTmp.F_area_county_id = helper.StrToInt(maps[0]["F_area_id"].(string))
					areaInfoTmp.F_area_county_name = maps[0]["F_area_name"].(string)
					parentId := helper.StrToInt(maps[0]["F_area_parent"].(string))
					if areaInfoTmp.F_area_county_id > 0{
						//查学校所属市的ID,name
						var maps []orm.Params
						num, err := o.Raw("SELECT F_area_id,F_area_name,F_area_parent FROM t_area WHERE F_area_level = 2 AND F_area_id = ? LIMIT 1",parentId).Values(&maps)
						if err == nil && num > 0 {
							areaInfoTmp.F_area_city_id = helper.StrToInt(maps[0]["F_area_id"].(string))
							areaInfoTmp.F_area_city_name = maps[0]["F_area_name"].(string)
							parentId := helper.StrToInt(maps[0]["F_area_parent"].(string))
							if areaInfoTmp.F_area_city_id > 0{
								//查学校所属省的ID,name
								var maps []orm.Params
								num, err := o.Raw("SELECT F_area_id,F_area_name FROM t_area WHERE F_area_level = 1 AND F_area_id = ? LIMIT 1",parentId).Values(&maps)
								if err == nil && num > 0 {
									areaInfoTmp.F_area_province_id = helper.StrToInt(maps[0]["F_area_id"].(string))
									areaInfoTmp.F_area_province_name = maps[0]["F_area_name"].(string)
									areaInfoTmp.F_school_id = schoolId
								}
							}
						}
					}
				}
			}
			if areaInfoTmp.F_area_province_id > 0 && len(areaInfoTmp.F_area_province_name) > 0 && areaInfoTmp.F_area_city_id > 0 &&		len(areaInfoTmp.F_area_city_name) > 0 && areaInfoTmp.F_area_county_id > 0 && len(areaInfoTmp.F_area_county_name) > 0{
				areaInfo = areaInfoTmp
			}
		}
	}
	
	return areaInfo
}

//根据学校ID查询学校名
func (u *MSchool) GetSchoolNameById(ids []string)schoolList{
	schools := make(schoolList,0)
	if len(ids) > 0{
		o := orm.NewOrm()
		var maps []orm.Params
		idStr := helper.JoinString(ids,",")
		num, err := o.Raw("SELECT F_school_id,F_school,F_school_type FROM t_school WHERE F_school_id IN("+idStr+")").Values(&maps)
		if err == nil && num > 0 {
			schools = make(schoolList,num)
			for key,item := range maps{
				schools[key] = schoolItem{F_school_id:helper.StrToInt(item["F_school_id"].(string)),F_school:item["F_school"].(string),F_school_type:SchoolType[helper.StrToInt(item["F_school_type"].(string))]}
			}
		}
	}
	return schools
}
