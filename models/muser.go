package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"dream_api_sms_v2/helper"
	"fmt"
	"crypto/md5"
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
		res, err := o.Raw("INSERT INTO t_user SET F_user_name = ?,F_user_password=?", userName,userPwd).Exec()
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
		_, err := o.Raw("UPDATE t_user SET F_user_password=? WHERE F_user_name=?",userPwd,userName).Exec()
		if err == nil {
			result = 0
			/*
			num, _ := res.RowsAffected()
			if num >0{
				result = 0
			}
			*/
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
func (u *MUser) GetToken(userName string,pkg string)string{
	if len(userName) > 0 && len(pkg) > 0{
		//创建token
		token := fmt.Sprintf("%x", md5.Sum([]byte(helper.CreatePwd(4))))
		//写入数据库
		o := orm.NewOrm()
		res, err := o.Raw("REPLACE INTO t_token SET F_user_name = ?,F_pkg=?,F_token=?", userName,pkg,token).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			if num >0{
				return token
			}
		}
	}
	return ""
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
			//生日
			info["F_birthday"] = ""
			//学校
			info["F_school"] = ""
			//省份
			info["F_province"] = ""
			//城市
			info["F_city"] = ""
			//镇
			info["F_county"] = ""
			//区
			info["F_area"] = ""
			//真实名
			info["F_user_realname"] = "管韩强"
		}
	}
	return info
}