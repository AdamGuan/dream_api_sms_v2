package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"dream_api_sms_v2/helper"
)

func init() {
}

type MSign struct {
}

//检查sign是否正确
func (u *MSign) CheckSign(signOri string,userName string,pkg string,token string) bool{
	//get token
	o := orm.NewOrm()
	tokenNew := ""
	if len(token) > 0{
		tokenNew = token
		//计算sign是否正确
		return helper.CheckSign(signOri,tokenNew)
	}else{
		if len(userName) > 0 && len(pkg) > 0{
			var maps []orm.Params
			num, err := o.Raw("SELECT F_token FROM t_token WHERE F_user_name=? AND F_pkg=? AND F_expire_datetime > ? LIMIT 1", userName,pkg,helper.GetNowDateTime()).Values(&maps)
			if err == nil && num > 0 {
				tokenNew = maps[0]["F_token"].(string)
				//计算sign是否正确
				return helper.CheckSign(signOri,tokenNew)
			}
		}
	}
	return false
}