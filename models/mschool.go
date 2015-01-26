package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
//	"dream_api_sms_v2/helper"
)

func init() {
}

type MSchool struct {
}

//查询学校
func (u *MSchool) QuerySchools(name string)[]string{
	schools := make([]string,0)
	if len(name) > 0 {
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT F_school FROM t_school WHERE F_school LIKE '%"+name+"%'").Values(&maps)
		if err == nil && num > 0 {
			schools := make([]string,num)
			for key,item := range maps{
				schools[key] = item["F_school"].(string)
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