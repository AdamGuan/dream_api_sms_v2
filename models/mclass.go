package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"dream_api_sms_v2/helper"
)

func init() {
}

type MClass struct {
}

type allClassInfo []struct{
	F_class_id int
	F_class_name string
	F_class_person_total int
}

type classInfo struct{
	F_class_id int
	F_class_name string
	F_class_person_total int
}

//创建一个班级
func (u *MClass) CreateAClass(userName string,className string,schoolId int)int{
	o := orm.NewOrm()
	//查询学校是否存在
	result := -11
	var maps []orm.Params
	num, err := o.Raw("SELECT F_school_id FROM t_school WHERE F_school_id = ? LIMIT 1",schoolId).Values(&maps)
	if err == nil && num > 0 {
		result = 0
	}
	//判断班级名是否合法
	if result == 0{
		result = -12
		if len(className) > 0 && len(className) <= 40{
			result = -13
			var maps []orm.Params
			num, err := o.Raw("SELECT F_class_id FROM t_class WHERE F_school_id = ? AND F_class_name = ? LIMIT 1",schoolId,className).Values(&maps)
			if err == nil && num <= 0 {
				result = 0
			}
		}
	}
	//写入数据库
	if result == 0{
		result = -1
		res, err := o.Raw("INSERT INTO t_class SET F_class_name = ?,F_school_id=?", className,schoolId).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			classId, _ := res.LastInsertId()
			if num >0 && classId > 0{
				//给对应的用户绑定这个class
				res, err = o.Raw("UPDATE t_user SET F_class_id = ? WHERE F_user_name = ?", classId,userName).Exec()
				if err == nil {
					num, _ = res.RowsAffected()
					if num > 0{
						result = 0
					}
				}
			}
		}
	}
	return result
}

//获取某个学校下的所有班级信息
func (u *MClass) GetSchoolClassInfo(schoolId int)allClassInfo{
	schools := make(allClassInfo,0)

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT t_class.F_class_id,t_class.F_class_name,count(t_user.F_user_name) as total_person FROM t_class,t_user where t_class.F_school_id = ? AND t_class.F_class_id = t_user.F_class_id GROUP BY t_class.F_class_id",schoolId).Values(&maps)
	if err == nil && num > 0 {
		schools = make(allClassInfo,num)
		for key,item := range maps{
			schools[key] = classInfo{F_class_id:helper.StrToInt(item["F_class_id"].(string)),F_class_person_total:helper.StrToInt(item["total_person"].(string)),F_class_name:item["F_class_name"].(string)}
		}
	}
	return schools
}