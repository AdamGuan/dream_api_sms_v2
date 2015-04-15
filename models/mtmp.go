package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
}

type MTmp struct {
}

func (u *MTmp) DeleteAllUser(){
	o := orm.NewOrm()
	o.Raw("DELETE FROM t_user WHERE 1").Exec()
	o.Raw("DELETE FROM t_sms_rate WHERE 1").Exec()
	o.Raw("DELETE FROM t_token WHERE 1").Exec()
	o.Raw("DELETE FROM t_coin WHERE 1").Exec()
	o.Raw("DELETE FROM t_register_history WHERE 1").Exec()
	o.Raw("DELETE FROM t_auth_qq WHERE 1").Exec()
	o.Raw("DELETE FROM t_auth_weixin WHERE 1").Exec()
	o.Raw("DELETE FROM t_auth_xinlangweibo WHERE 1").Exec()
}

func (u *MTmp) DeleteUser(username string){
	if len(username) > 0 {
		o := orm.NewOrm()
		o.Raw("DELETE FROM t_user WHERE F_user_name=?",username).Exec()
		o.Raw("DELETE FROM t_token WHERE F_user_name=?",username).Exec()
		o.Raw("DELETE FROM t_coin WHERE F_user_name=?",username).Exec()
		o.Raw("DELETE FROM t_register_history WHERE F_user_name=?",username).Exec()
		o.Raw("DELETE FROM t_auth_qq WHERE F_user_name=?",username).Exec()
		o.Raw("DELETE FROM t_auth_weixin WHERE F_user_name=?",username).Exec()
		o.Raw("DELETE FROM t_auth_xinlangweibo WHERE F_user_name=?",username).Exec()
	}
}