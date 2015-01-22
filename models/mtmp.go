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
}

func (u *MTmp) DeleteUser(username string){
	if len(username) > 0{
		o := orm.NewOrm()
		o.Raw("DELETE FROM t_user WHERE F_user_name=?",username).Exec()
	}
}