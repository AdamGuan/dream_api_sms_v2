package models

import (
	"dream_api_sms_v2/helper"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/astaxie/beego/config" 
	"net/smtp"
	"strings"
	"github.com/astaxie/beego/utils"
)

var EmailMinute int
var EmailValidMinute int
var emailIdentify string
var emailUserName string
var emailUserPwd string
var emailHost string
var emailAddress string
var emailFrom string
var emailTitle string
var emailContent string
var emailHostSuffix string

func init() {
	otherconf, _ := config.NewConfig("ini", "conf/other.conf")
	EmailMinute,_ = otherconf.Int("emailMinute")
	EmailValidMinute,_ = otherconf.Int("emailValidMinute")

	emailIdentify = otherconf.String("emailIdentify")
	emailUserName = otherconf.String("emailUserName")
	emailUserPwd = otherconf.String("emailUserPwd")
	emailHost = otherconf.String("emailHost")
	emailAddress = otherconf.String("emailAddress")
	emailFrom = otherconf.String("emailFrom")
	emailTitle = otherconf.String("emailTitle")
	emailContent = otherconf.String("emailContent")
	emailHostSuffix = otherconf.String("emailHostSuffix")
}

type MEmail struct {
}

//获取一个email验证码并发送到邮箱
func (u *MEmail) GetEmailCode(email string) string {
	code := helper.CreatePwd(6)
	if u.sendEmail(email,code){
		return code
	}
	return ""
}

//发送email
func (u *MEmail) sendEmail(email string,code string) bool {
	content :=  "Message-ID: <"+helper.GetNowDateTime()+code+"@"+emailHostSuffix+">\r\nTo: "+email+"\r\nFrom: "+emailTitle+"<"+emailUserName+">\r\nSubject: "+emailTitle+"\r\n"+"Content-Type: text/plain; charset=UTF-8\r\n\r\n"+strings.Replace(emailContent, "$code", code,-1)
	mailList := []string{email}
	auth := smtp.PlainAuth(
		emailIdentify,
		emailUserName,
		emailUserPwd,
		emailHost,
	)
	//log
	logStr := "\nsend email\n"+utils.GetDisplayString("mailList", mailList, "content",content)
	EmailLog.Info(logStr)

	err := smtp.SendMail(
		emailAddress,
		auth,
		emailFrom,
		mailList,
		[]byte(content),
	)
	if err != nil {
		//log
		logStr := "\nsend email result\n"+utils.GetDisplayString("err", err)
		EmailLog.Info(logStr)
		return false
	}else{
		//log
		logStr := "\nsend email result true\n"
		EmailLog.Info(logStr)
	}
	return true
}

//valid a email code
func (u *MEmail) ValidEmail(pkgName string,code string,email string) bool {
	//检查本地数据库是否已保存对应的验证码
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_action FROM t_email_action_valid WHERE F_action=? AND F_last_timestamp > ? LIMIT 1", helper.Md5(email+pkgName+code),helper.GetDateTimeBeforeMinute(EmailValidMinute)).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}

//检查是否可以给用户发送email了
func (u *MEmail) CheckEmailRateValid(email string,pkgName string)bool{
	o := orm.NewOrm()
	var maps []orm.Params
	nowTime := time.Now().Add(-time.Minute * time.Duration(EmailMinute)).Format("2006-01-02 15:04:05")
	num, err := o.Raw("SELECT F_last_timestamp FROM t_email_rate WHERE F_action=? LIMIT 1", helper.Md5(email+pkgName)).Values(&maps)
	if err == nil && num > 0 {
		if maps[0]["F_last_timestamp"].(string) <= nowTime{
			return true
		}else{
			return false
		}
	}else{
		return true
	}
	return false
}

//写入email发送频率表
func (u *MEmail) AddEmailRate(email string,pkgName string){
	//写入数据库
	o := orm.NewOrm()
	o.Raw("replace into t_email_rate(F_action,F_last_timestamp) values('"+helper.Md5(email+pkgName)+"','"+time.Now().Format("2006-01-02- 15:04:05")+"')").Exec()
}

//删除email发送频率表
func (u *MEmail) DeleteEmailRate(email string,pkgName string){
	o := orm.NewOrm()
	o.Raw("UPDATE t_email_rate SET F_last_timestamp='1001-01-01 00:00:00' WHERE F_action=?",helper.Md5(email+pkgName)).Exec()
}

//写入t_email_action_valid
func (u *MEmail) AddEmailActionvalid(email string,pkgName string,code string)bool{
	//写入数据库
	o := orm.NewOrm()
	res, err := o.Raw("replace into t_email_action_valid(F_action,F_last_timestamp) values('"+helper.Md5(email+pkgName+code)+"','"+helper.GetNowDateTime()+"')").Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num >0{
			return true
		}
	}

	return false
}

//删除t_email_action_valid
func (u *MEmail) DeleteEmailActionvalid(email string,pkgName string,code string){
	o := orm.NewOrm()
	o.Raw("DELETE FROM t_email_action_valid WHERE F_action=?",helper.Md5(email+pkgName+code)).Exec()
}

//验证t_eamil_action_valid
func (u *MEmail) CheckEmailActionvalid(email string,pkgName string,code string)bool{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_action FROM t_email_action_valid WHERE F_action=? AND F_last_timestamp > ? LIMIT 1", helper.Md5(email+pkgName+code),helper.GetDateTimeBeforeMinute(EmailValidMinute)).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}