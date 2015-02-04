package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"dream_api_sms_v2/helper"
	"fmt"
	"crypto/md5"
	"strings"
)

func init() {
}

type MUser struct {
}

type userInfo struct {
	F_phone_number string
	F_gender string
	F_grade string
	F_birthday string
	F_school string
	F_school_id int
	F_province string
	F_province_id int
	F_city string
	F_city_id int
	F_county string
	F_county_id int
	F_user_realname string
	F_crate_datetime string
	F_modify_datetime string
	F_class_id int
	F_class_name string
}

//检查用户名是否可用
func (u *MUser) CheckUserNameValid(userName string)int{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_name=?", userName).Values(&maps)
	if err == nil && num <= 0 {
		return 0
	}
	return -2

}

//检查用户名是否存在
func (u *MUser) CheckUserNameExists(userName string)bool{
	if len(userName) <= 0{
		return false
	}
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_name=?", userName).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}

//检查用户名密码是否正确
func (u *MUser) CheckUserAndPwd(userName string,userPwd string)bool{
	if len(userName) <= 0 || len(userPwd) <= 0{
		return false
	}
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_name=? AND F_user_password = ?", userName,userPwd).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}


//检查用户密码
func (u *MUser) CheckUserPwdValid(userPwd string)int{
	if len(userPwd) > 0{
		return 0
	}
	return -3
}

//添加用户
func (u *MUser) AddUser(userName string,userPwd string)int{
	result := u.CheckUserNameValid(userName)
	if result == 0{
		result = u.CheckUserPwdValid(userPwd)
	}
	if result == 0{
		result = -1
		//写入数据库
		o := orm.NewOrm()
		now := helper.GetNowDateTime()
		res, err := o.Raw("INSERT INTO t_user SET F_user_name = ?,F_user_password=?,F_crate_datetime=?,F_modify_datetime=?", userName,userPwd,now,now).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			if num >0{
				result = 0
			}
		}
	}
	return result
}

//修改用户密码
func (u *MUser) ModifyUserPwd(userName string,userPwd string)int{
	result := -1
	res := u.CheckUserNameExists(userName)
	if res {
		result = u.CheckUserPwdValid(userPwd)
	}
	if result == 0{
		result = -1
		//写入数据库
		o := orm.NewOrm()
		_, err := o.Raw("UPDATE t_user SET F_user_password=?,F_modify_datetime=? WHERE F_user_name=?",userPwd,helper.GetNowDateTime(),userName).Exec()
		if err == nil {
			result = 0
		}
	}
	return result
}

//获取用户的密码
func (u *MUser) GetUserPwd(userName string)string{
	pwd := ""
	if len(userName) > 0{
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT F_user_password FROM t_user WHERE F_user_name=? LIMIT 1", userName).Values(&maps)
		if err == nil && num > 0 {
			pwd = maps[0]["F_user_password"].(string)
		}
	}
	return pwd

}

//获取用户token,并写入数据库
func (u *MUser) GetToken(userName string,pkg string)(token string,tokenExpireDatetime string){
	if len(userName) > 0 && len(pkg) > 0{
		//创建token
		token := fmt.Sprintf("%x", md5.Sum([]byte(helper.CreatePwd(4))))
		tokenExpireDatetime := helper.GetDateTimeAfterMinute(60*24*30)
		//写入数据库
		o := orm.NewOrm()
		res, err := o.Raw("REPLACE INTO t_token SET F_user_name = ?,F_pkg=?,F_token=?,F_expire_datetime=?", userName,pkg,token,tokenExpireDatetime).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			if num >0{
				return token,tokenExpireDatetime
			}
		}
	}
	return "",""
}

//获取用户的信息
func (u *MUser) GetUserInfo(userName string)userInfo{
	info := userInfo{}
	if len(userName) > 0 {
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT * FROM t_user WHERE F_user_name=? LIMIT 1", userName).Values(&maps)
		if err == nil && num > 0 {
			info.F_phone_number = maps[0]["F_user_name"].(string)
			//性别
			gender := maps[0]["F_gender"].(string)
			genderint := helper.StrToInt(gender)
			info.F_gender = "男"
			if genderint != 1{
				info.F_gender = "女"
			}
			//年级
			info.F_grade = ""
			if maps[0]["F_grade_id"] != nil{
			tmp,ok := Grade[maps[0]["F_grade_id"].(string)]
			if ok{
				info.F_grade = tmp
			}
			}
			//生日
			info.F_birthday = ""
			if maps[0]["F_birthday"] != nil{
				info.F_birthday = maps[0]["F_birthday"].(string)
			}
			//学校
			info.F_school,_ = School[maps[0]["F_school_id"].(string)]
			info.F_school_id = helper.StrToInt(maps[0]["F_school_id"].(string))
			//省份
			info.F_province,_ = Province[maps[0]["F_province_id"].(string)]
			info.F_province_id = helper.StrToInt(maps[0]["F_province_id"].(string))
			//城市
			info.F_city,_ = City[maps[0]["F_city_id"].(string)]
			info.F_city_id = helper.StrToInt(maps[0]["F_city_id"].(string))
			//县
			info.F_county,_ = County[maps[0]["F_county_id"].(string)]
			info.F_county_id = helper.StrToInt(maps[0]["F_county_id"].(string))
			//真实名
			info.F_user_realname = ""
			if maps[0]["F_user_realname"] != nil{
				info.F_user_realname = maps[0]["F_user_realname"].(string)
			}
			//创建时间
			info.F_crate_datetime = ""
			if maps[0]["F_crate_datetime"] != nil{
				info.F_crate_datetime = maps[0]["F_crate_datetime"].(string)
			}
			//修改时间
			info.F_modify_datetime = ""
			if maps[0]["F_modify_datetime"] != nil{
				info.F_modify_datetime = maps[0]["F_modify_datetime"].(string)
			}
			//年级
			info.F_class_id = helper.StrToInt(maps[0]["F_class_id"].(string))
			info.F_class_name = ""
			var maps2 []orm.Params
			num, err := o.Raw("SELECT F_class_name FROM t_class WHERE F_class_id = ? LIMIT 1",maps[0]["F_class_id"].(string)).Values(&maps2)
			if err == nil && num > 0 {
				info.F_class_name = maps2[0]["F_class_name"].(string)
			}
		}
	}
	return info
}

//修改用户的信息
func (u *MUser) ModifyUserInfo(parames map[string]string)int{
	result := -1
	userName,ok := parames["mobilePhoneNumber"]
	if ok{
		result = 0
		breakd := 0
		set := ""
		for filed,value := range parames{
			switch filed {
				case "gender":
					if value == "男" || value == "女"{
						g := "2"
						if value == "男" {
							g = "1"
						}
						set += "F_gender="+g+","
					}else{
						breakd = 1
						result = -10
					}
				case "grade":
					tmp := 0
					for id,v := range Grade{
						if value == v{
							set += "F_grade_id="+id+","
							tmp = 1
						}
					}
					if tmp == 0 {
						breakd = 1
						result = -10
					}
				case "birthday":
					if len(value) > 0{
						set += "F_birthday='"+value+"',"
					}else{
						breakd = 1
						result = -10
					}
				case "school":
					_,ok := School[value]
					if ok{
						set += "F_school_id="+value+","
					}else{
						breakd = 1
						result = -10
					}
				case "province":
					_,ok := Province[value]
					if ok{
						set += "F_province_id="+value+","
					}else{
						breakd = 1
						result = -10
					}
				case "city":
					_,ok := City[value]
					if ok{
						set += "F_city_id="+value+","
					}else{
						breakd = 1
						result = -10
					}
				case "county":
					_,ok := County[value]
					if ok{
						set += "F_county_id="+value+","
					}else{
						breakd = 1
						result = -10
					}
				case "realname":
					if len(value) > 0{
						set += "F_user_realname='"+value+"',"
					}else{
						breakd = 1
						result = -10
					}
				default:
			}
			if breakd == 1{
				break
			}
		}
		//写入数据库
		if result == 0{
			result = -1
			set = strings.Trim(set, ",")
			if len(set) > 0{
				set += ",F_modify_datetime='"+helper.GetNowDateTime()+"'"
				o := orm.NewOrm()
				res, err := o.Raw("UPDATE t_user SET "+set+" WHERE F_user_name=?",userName).Exec()
				if err == nil {
					num, _ := res.RowsAffected()
					if num >0 {
						result = 0
					}
				}
			}
		}
	}
	return result
}

//用户登出
func (u *MUser) UserLoginout(userName string,pkg string)bool{
	result := false
	//写入数据库
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM t_token WHERE F_user_name=? AND F_pkg=?",userName,pkg).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num >0 {
			result = true
		}
	}
	return result
}

//用户修改班级
func (u *MUser) UserChangeClass(userName string,classId int)int{
	result := -14
	//查询班级是否存在
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_class_id FROM t_class WHERE F_class_id = ? LIMIT 1",classId).Values(&maps)
	if err == nil && num > 0 {
		result = -1
		//修改用户的班级
		now := helper.GetNowDateTime()
		res, err := o.Raw("UPDATE t_user SET F_class_id = ?,F_modify_datetime= ? WHERE F_user_name = ?",classId,now,userName).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			if num >0 {
				result = 0
			}
		}
	}

	return result
}