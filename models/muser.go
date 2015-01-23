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
func (u *MUser) GetUserInfo(userName string)map[string]string{
	info := make(map[string]string)
	if len(userName) > 0 {
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT * FROM t_user WHERE F_user_name=? LIMIT 1", userName).Values(&maps)
		if err == nil && num > 0 {
			info["F_phone_number"] = maps[0]["F_user_name"].(string)
			//性别
			gender := maps[0]["F_gender"].(string)
			genderint := helper.StrToInt(gender)
			info["F_gender"] = "男"
			if genderint != 1{
				info["F_gender"] = "女"
			}
			//年级
			info["F_grade"] = ""
			if maps[0]["F_grade_id"] != nil{
			tmp,ok := Grade[maps[0]["F_grade_id"].(string)]
			if ok{
				info["F_grade"] = tmp
			}
			}
			//生日
			info["F_birthday"] = ""
			if maps[0]["F_birthday"] != nil{
				info["F_birthday"] = maps[0]["F_birthday"].(string)
			}
			//学校
			info["F_school"] = ""
			if maps[0]["F_school_id"] != nil{
				tmp,ok := School[maps[0]["F_school_id"].(string)]
				if ok{
					info["F_school"] = tmp
				}
			}
			//省份
			info["F_province"] = ""
			if maps[0]["F_province_id"] != nil{
			tmp,ok := Province[maps[0]["F_province_id"].(string)]
			if ok{
				info["F_province"] = tmp
			}
			}
			//城市
			info["F_city"] = ""
			if maps[0]["F_city_id"] != nil{
			tmp,ok := City[maps[0]["F_city_id"].(string)]
			if ok{
				info["F_city"] = tmp
			}
			}
			//镇
			info["F_county"] = ""
			if maps[0]["F_county_id"] != nil{
			tmp,ok := County[maps[0]["F_county_id"].(string)]
			if ok{
				info["F_county"] = tmp
			}
			}
			//区
			info["F_town"] = ""
			if maps[0]["F_town_id"] != nil{
			tmp,ok := Town[maps[0]["F_town_id"].(string)]
			if ok{
				info["F_town"] = tmp
			}
			}
			//真实名
			info["F_user_realname"] = ""
			if maps[0]["F_user_realname"] != nil{
				info["F_user_realname"] = maps[0]["F_user_realname"].(string)
			}
			//创建时间
			info["F_crate_datetime"] = ""
			if maps[0]["F_crate_datetime"] != nil{
				info["F_crate_datetime"] = maps[0]["F_crate_datetime"].(string)
			}
			//修改时间
			info["F_modify_datetime"] = ""
			if maps[0]["F_modify_datetime"] != nil{
				info["F_modify_datetime"] = maps[0]["F_modify_datetime"].(string)
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
					tmp := 0
					for id,v := range School{
						if value == v{
							set += "F_school_id="+id+","
							tmp = 1
						}
					}
					if tmp == 0 {
						breakd = 1
						result = -10
					}
				case "province":
					tmp := 0
					for id,v := range Province{
						if value == v{
							set += "F_province_id="+id+","
							tmp = 1
						}
					}
					if tmp == 0 {
						breakd = 1
						result = -10
					}
				case "city":
					tmp := 0
					for id,v := range City{
						if value == v{
							set += "F_city_id="+id+","
							tmp = 1
						}
					}
					if tmp == 0 {
						breakd = 1
						result = -10
					}
				case "county":
					tmp := 0
					for id,v := range County{
						if value == v{
							set += "F_county_id="+id+","
							tmp = 1
						}
					}
					if tmp == 0 {
						breakd = 1
						result = -10
					}
				case "town":
					tmp := 0
					for id,v := range Town{
						if value == v{
							set += "F_town_id="+id+","
							tmp = 1
						}
					}
					if tmp == 0 {
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