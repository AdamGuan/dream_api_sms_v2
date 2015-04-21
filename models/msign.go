package models

import (
	"dream_api_sms_v2/helper"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
}

type MSign struct {
}

//检查sign是否正确
func (u *MSign) CheckSign(signOri string, userName string, pkg string, token string) bool {
	//get token
	o := orm.NewOrm()
	tokenNew := ""
	if len(token) > 0 {
		tokenNew = token
		//计算sign是否正确
		return helper.CheckSign(signOri, tokenNew)
	} else {
		if len(userName) > 0 && len(pkg) > 0 {
			var maps []orm.Params
			num, err := o.Raw("SELECT F_token FROM t_token WHERE F_user_name=? AND F_pkg=? AND F_expire_datetime > ? LIMIT 1", userName, pkg, helper.GetNowDateTime()).Values(&maps)
			if err == nil && num > 0 {
				tokenNew = maps[0]["F_token"].(string)
				//计算sign是否正确
				return helper.CheckSign(signOri, tokenNew)
			}
		}
	}
	return false
}

//检查token是否正确
func (u *MSign) CheckToken(pkg string, token string) (result bool, uid string) {
	//get token
	o := orm.NewOrm()
	if len(token) > 0 && len(pkg) > 0 {
		if len(pkg) > 0 {
			var maps []orm.Params
			num, err := o.Raw("SELECT F_user_name,F_token FROM t_token WHERE F_pkg=? AND F_expire_datetime > ? AND F_token = ? LIMIT 1", pkg, helper.GetNowDateTime(), token).Values(&maps)
			if err == nil && num > 0 {
				F_user_id := maps[0]["F_user_name"].(string)
				return true, F_user_id
			}
		}
	}
	return false, "0"
}

//删除token
func (u *MSign) DeleteAllPkgToken(userName string) bool {
	//get token
	o := orm.NewOrm()
	if len(userName) > 0 {
		o.Raw("DELETE FROM t_token WHERE F_user_name=?", userName).Exec()
		return true
	}
	return false
}
