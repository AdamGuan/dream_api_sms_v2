package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"dream_api_sms_v2/helper"
	"fmt"
	"crypto/md5"
	"strings"
	"github.com/astaxie/beego/config" 
	"github.com/astaxie/beego"
)

func init() {
}

type MUser struct {
}

type userInfo struct {
	F_phone_number string
	F_gender string
	F_grade string
	F_grade_id int
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
	F_user_nickname string
	F_crate_datetime string
	F_modify_datetime string
	F_class_id int
	F_class_name string
	F_avatar_url string
}

type avatarSysInfoList []struct{
	F_avatar_url string
	F_avatar_id int
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
	if helper.CheckPwdValid(userPwd){
		return 0
	}
	return -3
}

//添加用户
func (u *MUser) AddUser(parames map[string]string)int{
	result := -16
	userName,ok := parames["mobilePhoneNumber"]
	if ok{
		result = u.CheckUserNameValid(userName)
	}
	userPwd,ok := parames["pwd"]
	if result == 0 && ok{
		result = u.CheckUserPwdValid(userPwd)
	}
	if result == 0{
		/**/
		breakd := 0
		now := helper.GetNowDateTime()
		set := "F_user_name = '"+userName+"',F_user_password = '"+userPwd+"',F_crate_datetime='"+now+"',F_modify_datetime='"+now+"',"
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
					if helper.CheckRealNameValid(value){
						set += "F_user_realname='"+value+"',"
					}else{
						breakd = 1
						result = -24
					}
				default:
			}
			if breakd == 1{
				break
			}
		}
		/**/
		//写入数据库
		if result == 0 {
			result = -1
			set = strings.Trim(set, ",")
			if len(set) > 0{
				o := orm.NewOrm()
				res, err := o.Raw("INSERT INTO t_user SET "+set).Exec()
				if err == nil {
					num, _ := res.RowsAffected()
					if num >0{
						result = 0
					}
				}
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
					info.F_grade_id = helper.StrToInt(maps[0]["F_grade_id"].(string))
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
			//昵称
			info.F_user_nickname = ""
			if maps[0]["F_user_nickname"] != nil{
				info.F_user_nickname = maps[0]["F_user_nickname"].(string)
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
			//班级
			info.F_class_id = helper.StrToInt(maps[0]["F_class_id"].(string))
			info.F_class_name = ""
			var maps2 []orm.Params
			num, err := o.Raw("SELECT F_class_name FROM t_class WHERE F_class_id = ? LIMIT 1",maps[0]["F_class_id"].(string)).Values(&maps2)
			if err == nil && num > 0 {
				info.F_class_name = maps2[0]["F_class_name"].(string)
			}
			//头像
			avatartmp := maps[0]["F_avatarname"].(string)
			if len(avatartmp) > 0{
				info.F_avatar_url = u.getUserAvatarUrl(avatartmp,helper.StrToInt(avatartmp[0:1]))
			}else{
				info.F_avatar_url = ""
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
					if helper.CheckRealNameValid(value){
						set += "F_user_realname='"+value+"',"
					}else{
						breakd = 1
						result = -24
					}
				case "nickname":
					if helper.CheckNickNameValid(value){
						set += "F_user_nickname='"+value+"',"
					}else{
						breakd = 1
						result = -25
					}
				case "avatarSysName":
					if len(value) > 0{
						set += "F_avatarname='"+value+"',"
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
	num, err := o.Raw("SELECT t_class.* FROM t_class,t_user WHERE t_class.F_class_id = ? AND t_class.F_class_id = t_user.F_class_id LIMIT 1",classId).Values(&maps)
	if err == nil && num > 0 {
		result = -1
		//修改用户的班级
		F_school_id := maps[0]["t_class.F_school_id"]
		F_grade_id := maps[0]["t_class.F_grade_id"]
		now := helper.GetNowDateTime()
		res, err := o.Raw("UPDATE t_user SET F_class_id = ?,F_school_id=?,F_grade_id=?,F_modify_datetime= ? WHERE F_user_name = ?",classId,F_school_id,F_grade_id,now,userName).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			if num >0 {
				result = 0
			}
		}
	}

	return result
}

//用户头像修改
func (u *MUser) UserAvatarNameModify(userName string,avatarName string)bool{
	result := false
	o := orm.NewOrm()
	now := helper.GetNowDateTime()
	res, err := o.Raw("UPDATE t_user SET F_avatarname = ?,F_modify_datetime= ? WHERE F_user_name = ?",avatarName,now,userName).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num >0 {
			result = true
		}
	}
	return result
}

//获取用户头像url
func (u *MUser) getUserAvatarUrl(avatarName string,atype int)string{
	url := ""
	//domain
	appconf, _ := config.NewConfig("ini", "conf/app.conf")
	doamin := appconf.String(beego.RunMode+"::domain")
	//path
	pre := "avatar/"
	if atype == 2{
		pre = "avatar2/"
		//build url
		url = doamin+"/"+pre+avatarName
	}else{
		//build url
		url = doamin+"/"+pre+helper.Md5(avatarName)[0:2]+"/"+avatarName
	}
	return url
}

//获取系统内置用户头像url列表
func (u *MUser) GetAvatarUrlList()avatarSysInfoList{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM t_sys_avatar where 1").Values(&maps)
	if err == nil && num > 0 {
		urlList := make(avatarSysInfoList,num)
		for key,item := range maps{
			urlList[key].F_avatar_url = u.getUserAvatarUrl(item["F_avatar_name"].(string),2)
			urlList[key].F_avatar_id = helper.StrToInt(item["F_id"].(string))
		}
		return urlList
	}
	return make(avatarSysInfoList,0)
}

//根据系统头像ID获取头像名称
func (u *MUser) GetAvatarNameFromId(id int)string{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_avatar_name FROM t_sys_avatar where F_id=?",id).Values(&maps)
	if err == nil && num > 0 {
		return maps[0]["F_avatar_name"].(string)
	}
	return ""
}

//修改用户手机号码
func (u *MUser) ModifyUserPhone(userName string,newUserName string)int{
	result := -1
	res := u.CheckUserNameExists(userName)
	if res {
		//检查是否新的手机号码已注册
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT F_user_name FROM t_user where F_user_name=?",newUserName).Values(&maps)
		if err == nil && num <= 0 {
			//更新手机号码
			o := orm.NewOrm()
			_, err := o.Raw("UPDATE t_user SET F_user_name=?,F_modify_datetime=? WHERE F_user_name=?",newUserName,helper.GetNowDateTime(),userName).Exec()
			if err == nil {
				result = 0
			}
		}else{
			result = -23
		}
	}
	return result
}