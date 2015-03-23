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

type MConsumer struct {
}

type userInfoa struct {
	F_uid string
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
	F_user_email string
}

type avatarSysInfoLista []struct{
	F_avatar_url string
	F_avatar_id int
}

//根据手机号码获取uid
func (u *MConsumer) GetUidByPhone(phone string)string{
	uid := ""
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_phone=? LIMIT 1", phone).Values(&maps)
	if err == nil && num > 0 {
		uid = maps[0]["F_user_name"].(string)
	}
	return uid
}

//根据email获取uid
func (u *MConsumer) GetUidByEmail(email string)string{
	uid := ""
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_email=? LIMIT 1", email).Values(&maps)
	if err == nil && num > 0 {
		uid = maps[0]["F_user_name"].(string)
	}
	return uid
}

//检查手机号码是否可用
func (u *MConsumer) CheckPhoneValid(phone string)int{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_phone=? LIMIT 1", phone).Values(&maps)
	if err == nil && num <= 0 {
		return 0
	}
	return -23

}

//检查email是否可做新用户使用
func (u *MConsumer) CheckEmailValid(email string)bool{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_email=? LIMIT 1", email).Values(&maps)
	if err == nil && num <= 0 {
		return true
	}
	return false
}

//检查用户ID是否可用
func (u *MConsumer) checkUserIdValid(uid string)bool{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_name=? LIMIT 1", uid).Values(&maps)
	if err == nil && num <= 0 {
		return true
	}
	return false
}

//创建一个可用的用户ID
func (u *MConsumer) CreateUid()string{
	uid := ""
	for{
		uid = helper.GetGuid()
		if u.checkUserIdValid(uid){
			break
		}
	}
	return uid
}

//检查手机号码是否存在
func (u *MConsumer) CheckPhoneExists(phone string)bool{
	if u.CheckPhoneValid(phone) != 0{
		return true
	}
	return false
}

//检查email是否存在
func (u *MConsumer) CheckEmailExists(email string)bool{
	if !u.CheckEmailValid(email) {
		return true
	}
	return false
}

//检查uid是否存在
func (u *MConsumer) CheckUserIdExists(uid string)bool{
	if !u.checkUserIdValid(uid){
		return true
	}
	return false
}

//检查手机号码与密码是否正确
func (u *MConsumer) CheckPhoneAndPwd(phone string,userPwd string)bool{
	if len(phone) <= 0 || len(userPwd) <= 0{
		return false
	}
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_phone=? AND F_user_password = ? LIMIT 1", phone,userPwd).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}

//检查email与密码是否正确
func (u *MConsumer) CheckEmailAndPwd(email string,userPwd string)bool{
	if len(email) <= 0 || len(userPwd) <= 0{
		return false
	}
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_email=? AND F_user_password = ? LIMIT 1", email,userPwd).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}

//检查用户ID与密码是否正确
func (u *MConsumer) CheckUserIdAndPwd(uid string,userPwd string)bool{
	if len(uid) <= 0 || len(userPwd) <= 0{
		return false
	}
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_name=? AND F_user_password = ? LIMIT 1", uid,userPwd).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}

//检查uid与email是否匹配
func (u *MConsumer) CheckUidAndEmail(uid string,email string)bool{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_name=? AND F_user_email = ? LIMIT 1", uid,email).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}

//检查用户密码
func (u *MConsumer) CheckUserPwdValid(userPwd string)int{
	if helper.CheckPwdValid(userPwd){
		return 0
	}
	return -3
}

//添加用户(根据手机号码)
func (u *MConsumer) AddUserByPhone(parames map[string]string)int{
	result := -16
	phone,ok := parames["mobilePhoneNumber"]
	if ok{
		result = 0
		if u.CheckPhoneValid(phone) != 0{
			result = -2
		}
	}
	if result == 0{
		return u.addUser(parames)
	}
	return result
}

//添加用户(根据email)
func (u *MConsumer) AddUserByEmail(parames map[string]string)int{
	result := -16
	email,ok := parames["email"]
	if ok{
		result = 0
		if !u.CheckEmailValid(email) {
			result = -2
		}
	}
	if result == 0{
		return u.addUser(parames)
	}
	return result
}

//添加用户
func (u *MConsumer) addUser(parames map[string]string)int{
	result := -10
	//检查pwd
	userPwd,ok := parames["pwd"]
	if ok{
		result = u.CheckUserPwdValid(userPwd)
	}else{
		result = -9
	}
	//foreach
	if result == 0{
		/**/
		breakd := 0
		now := helper.GetNowDateTime()
		set := "F_user_password = '"+userPwd+"',F_crate_datetime='"+now+"',F_modify_datetime='"+now+"',"
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
				case "mobilePhoneNumber":
					result = -10
					phone,ok := parames["mobilePhoneNumber"]
					if ok{
						if u.CheckPhoneValid(phone) != 0{
							result = -2
						}else{
							result = 0
						}
					}
					if result == 0{
						set += "F_user_phone='"+phone+"',"
					}else{
						breakd = 1
					}
				case "email":
					result = -10
					email,ok := parames["email"]
					if ok{
						if !u.CheckEmailValid(email) {
							result = -2
						}else{
							result = 0
						}
					}
					if result == 0{
						set += "F_user_email='"+email+"',"
					}else{
						breakd = 1
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
				set = "F_user_name='"+u.CreateUid()+"',"+set
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
func (u *MConsumer) ModifyUserPwdByUid(userId string,userPwd string)int{
	result := -1
	res := u.CheckUserIdExists(userId)
	if res {
		result = u.CheckUserPwdValid(userPwd)
	}
	if result == 0{
		result = -1
		//写入数据库
		o := orm.NewOrm()
		_, err := o.Raw("UPDATE t_user SET F_user_password=?,F_modify_datetime=? WHERE F_user_name=?",userPwd,helper.GetNowDateTime(),userId).Exec()
		if err == nil {
			result = 0
		}
	}
	return result
}

//修改用户密码(根据手机)
func (u *MConsumer) ModifyUserPwdByPhone(phone string,userPwd string)int{
	result := -1
	res := u.CheckPhoneExists(phone)
	if res {
		result = u.CheckUserPwdValid(userPwd)
	}
	if result == 0{
		result = -1
		//写入数据库
		o := orm.NewOrm()
		_, err := o.Raw("UPDATE t_user SET F_user_password=?,F_modify_datetime=? WHERE F_user_phone=?",userPwd,helper.GetNowDateTime(),phone).Exec()
		if err == nil {
			result = 0
		}
	}
	return result
}

//修改用户密码(根据email)
func (u *MConsumer) ModifyUserPwdByEmail(email string,userPwd string)int{
	result := -1
	res := u.CheckEmailExists(email)
	if res {
		result = u.CheckUserPwdValid(userPwd)
	}
	if result == 0{
		result = -1
		//写入数据库
		o := orm.NewOrm()
		_, err := o.Raw("UPDATE t_user SET F_user_password=?,F_modify_datetime=? WHERE F_user_email=?",userPwd,helper.GetNowDateTime(),email).Exec()
		if err == nil {
			result = 0
		}
	}
	return result
}

//获取用户的密码
func (u *MConsumer) GetUserPwdByUid(userId string)string{
	pwd := ""
	if len(userId) > 0{
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT F_user_password FROM t_user WHERE F_user_name=? LIMIT 1", userId).Values(&maps)
		if err == nil && num > 0 {
			pwd = maps[0]["F_user_password"].(string)
		}
	}
	return pwd

}

//获取用户的密码(根据手机号码)
func (u *MConsumer) GetUserPwdByPhone(phone string)string{
	pwd := ""
	if len(phone) > 0{
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT F_user_password FROM t_user WHERE F_user_phone=? LIMIT 1", phone).Values(&maps)
		if err == nil && num > 0 {
			pwd = maps[0]["F_user_password"].(string)
		}
	}
	return pwd

}

//获取用户token,并写入数据库
func (u *MConsumer) GetTokenByUid(userId string,pkg string)(token string,tokenExpireDatetime string){
	if len(userId) > 0 && len(pkg) > 0{
		//创建token
		token := fmt.Sprintf("%x", md5.Sum([]byte(helper.CreatePwd(4))))
		tokenExpireDatetime := helper.GetDateTimeAfterMinute(60*24*30)
		//写入数据库
		o := orm.NewOrm()
		res, err := o.Raw("REPLACE INTO t_token SET F_user_name = ?,F_pkg=?,F_token=?,F_expire_datetime=?", userId,pkg,token,tokenExpireDatetime).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			if num >0{
				return token,tokenExpireDatetime
			}
		}
	}
	return "",""
}

//获取用户token,并写入数据库(根据手机号码)
func (u *MConsumer) GetTokenByPhone(phone string,pkg string)(token string,tokenExpireDatetime string){
	uid := u.GetUidByPhone(phone)
	if len(uid) > 0{
		return u.GetTokenByUid(uid,pkg)
	}
	return "",""
}

//获取用户的信息
func (u *MConsumer) GetUserInfoByUid(userId string)userInfoa{
	info := userInfoa{}
	if len(userId) > 0 {
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT * FROM t_user WHERE F_user_name=? LIMIT 1", userId).Values(&maps)
		if err == nil && num > 0 {
			info.F_uid = maps[0]["F_user_name"].(string)
			info.F_phone_number = maps[0]["F_user_phone"].(string)
			info.F_user_email = maps[0]["F_user_email"].(string)
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

//获取用户的信息(根据手机号码)
func (u *MConsumer) GetUserInfoByPhone(phone string)userInfoa{
	uid := u.GetUidByPhone(phone)
	if len(uid) > 0{
		return u.GetUserInfoByUid(uid)
	}
	return userInfoa{}
}

//修改用户的信息
func (u *MConsumer) ModifyUserInfo(parames map[string]string)int{
	result := -1
	uid,ok := parames["uid"]
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
				res, err := o.Raw("UPDATE t_user SET "+set+" WHERE F_user_name=?",uid).Exec()
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
func (u *MConsumer) UserLoginout(userId string,pkg string)bool{
	result := false
	//写入数据库
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM t_token WHERE F_user_name=? AND F_pkg=?",userId,pkg).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num >0 {
			result = true
		}
	}
	return result
}

//用户修改班级
func (u *MConsumer) UserChangeClass(userId string,classId int)int{
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
		res, err := o.Raw("UPDATE t_user SET F_class_id = ?,F_school_id=?,F_grade_id=?,F_modify_datetime= ? WHERE F_user_name = ?",classId,F_school_id,F_grade_id,now,userId).Exec()
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
func (u *MConsumer) UserAvatarNameModify(userId string,avatarName string)bool{
	result := false
	o := orm.NewOrm()
	now := helper.GetNowDateTime()
	res, err := o.Raw("UPDATE t_user SET F_avatarname = ?,F_modify_datetime= ? WHERE F_user_name = ?",avatarName,now,userId).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num >0 {
			result = true
		}
	}
	return result
}

//获取用户头像url
func (u *MConsumer) getUserAvatarUrl(avatarName string,atype int)string{
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
func (u *MConsumer) GetAvatarUrlList()avatarSysInfoLista{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM t_sys_avatar where 1").Values(&maps)
	if err == nil && num > 0 {
		urlList := make(avatarSysInfoLista,num)
		for key,item := range maps{
			urlList[key].F_avatar_url = u.getUserAvatarUrl(item["F_avatar_name"].(string),2)
			urlList[key].F_avatar_id = helper.StrToInt(item["F_id"].(string))
		}
		return urlList
	}
	return make(avatarSysInfoLista,0)
}

//根据系统头像ID获取头像名称
func (u *MConsumer) GetAvatarNameFromId(id int)string{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_avatar_name FROM t_sys_avatar where F_id=?",id).Values(&maps)
	if err == nil && num > 0 {
		return maps[0]["F_avatar_name"].(string)
	}
	return ""
}

//修改用户手机号码
func (u *MConsumer) ModifyUserPhoneByPhone(phone string,newPhone string)int{
	result := -1
	res := u.CheckPhoneExists(phone)
	if res {
		result = u.CheckPhoneValid(newPhone)
		if result == 0 {
			//更新手机号码
			o := orm.NewOrm()
			_, err := o.Raw("UPDATE t_user SET F_user_phone=?,F_modify_datetime=? WHERE F_user_phone=?",newPhone,helper.GetNowDateTime(),phone).Exec()
			if err == nil {
				result = 0
			}
		}
	}
	return result
}

//修改用户手机号码
func (u *MConsumer) ModifyUserPhone(phone string,uid string)int{
	result := -1
	
	result = u.CheckPhoneValid(phone)
	if result == 0 {
		o := orm.NewOrm()
		res, err := o.Raw("UPDATE t_user SET F_user_phone=?,F_modify_datetime=? WHERE F_user_name=?",phone,helper.GetNowDateTime(),uid).Exec()
		if err == nil {
			num, err := res.RowsAffected()
			if err == nil && num > 0{
				result = 0
			}
		}
	}
	return result
}

//修改email
func (u *MConsumer) ModifyUserEmail(email string,uid string)int{
	result := -1
	
	res := u.CheckEmailValid(email)
	if res{
		o := orm.NewOrm()
		res, err := o.Raw("UPDATE t_user SET F_user_email=?,F_modify_datetime=? WHERE F_user_name=?",email,helper.GetNowDateTime(),uid).Exec()
		if err == nil {
			num, err := res.RowsAffected()
			if err == nil && num > 0{
				result = 0
			}
		}
	}else{
		result = -26
	}
	return result
}

//根据QQ号码返回UID
func (u *MConsumer) GetUidByQQ(qq string)string{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_auth_qq where F_qq_number=? LIMIT 1",qq).Values(&maps)
	if err == nil && num > 0 {
		return maps[0]["F_user_name"].(string)
	}
	return ""
}

//insert 一条qq认证信息
func (u *MConsumer) InsertQQ(qq string)string{
	uid := u.addUserQQ()
	if len(uid) > 0{
		o := orm.NewOrm()
		res, err := o.Raw("INSERT INTO t_auth_qq SET F_user_name=?,F_qq_number=?",uid,qq).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			if num >0{
				return uid
			}
		}else{
			o.Raw("DELETE FROM t_user WHERE F_user_name=? LIMIT 1",uid).Exec()
		}
	}
	return ""
}

//添加qq到t_user,并返回uid
func (u *MConsumer) addUserQQ()string{
	now := helper.GetNowDateTime()
	uid := u.CreateUid()
	set := "F_crate_datetime='"+now+"',F_modify_datetime='"+now+"'"
	set = "F_user_name='"+uid+"',"+set
	o := orm.NewOrm()
	res, err := o.Raw("INSERT INTO t_user SET "+set).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num >0{
			return uid
		}
	}

	return ""
}

//根据新浪微博用户名返回UID
func (u *MConsumer) GetUidByXinlangweibo(name string)string{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_auth_xinlangweibo where F_xinlangweibo_user=? LIMIT 1",name).Values(&maps)
	if err == nil && num > 0 {
		return maps[0]["F_user_name"].(string)
	}
	return ""
}

//insert 一条新浪微博认证信息
func (u *MConsumer) InsertXinlangweibo(name string)string{
	uid := u.addUserXinlangweibo()
	if len(uid) > 0{
		o := orm.NewOrm()
		res, err := o.Raw("INSERT INTO t_auth_xinlangweibo SET F_user_name=?,F_xinlangweibo_user=?",uid,name).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			if num >0{
				return uid
			}
		}else{
			o.Raw("DELETE FROM t_user WHERE F_user_name=? LIMIT 1",uid).Exec()
		}
	}
	return ""
}

//添加新浪微博到t_user,并返回uid
func (u *MConsumer) addUserXinlangweibo()string{
	now := helper.GetNowDateTime()
	uid := u.CreateUid()
	set := "F_crate_datetime='"+now+"',F_modify_datetime='"+now+"'"
	set = "F_user_name='"+uid+"',"+set
	o := orm.NewOrm()
	res, err := o.Raw("INSERT INTO t_user SET "+set).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num >0{
			return uid
		}
	}

	return ""
}