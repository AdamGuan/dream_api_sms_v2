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


//查询学校
func (u *MSchool) QuerySchools(name string,stype int,areaId int)schoolList{
	schools := make(schoolList,0)
	if len(name) > 0 {
		where := " WHERE 1"
		if stype != 0{
			where += " AND F_school_type = "+helper.IntToString(stype);
		}
		if areaId != 0{
			where += " AND F_belong_area_id = "+helper.IntToString(areaId);
		}
		where += " AND F_school LIKE '%"+name+"%'"

		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT * FROM t_school"+where).Values(&maps)
		if err == nil && num > 0 {
			schools := make(schoolList,num)
			for key,item := range maps{
				schools[key] = schoolItem{F_school_id:helper.StrToInt(item["F_school_id"].(string)),F_school:item["F_school"].(string),F_school_type:SchoolType[helper.StrToInt(item["F_school_type"].(string))]}
			}
			return schools
		}
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