package models

import (
	"dream_api_sms_v2/helper"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/astaxie/beego/config" 
//	"strconv"
	"github.com/astaxie/beego/utils"
)

var SmsMinute int
var smsValidMinute int

func init() {
	otherconf, _ := config.NewConfig("ini", "conf/other.conf")
	SmsMinute,_ = otherconf.Int("smsMinute")
	smsValidMinute,_ = otherconf.Int("smsValidMinute")
}

type MSms struct {
}

//get a msm
func (u *MSms) GetMsm(mobilePhoneNumber string,appId string,appKey string,appName string,appTemplate string) map[string]interface{} {
	url := "https://leancloud.cn/1.1/requestSmsCode"
	method := "POST"
	data := map[string]string{"mobilePhoneNumber": mobilePhoneNumber,"template":appTemplate,"appname":appName}

	//log curl
	logStr := "\nCurl leancloud for request code\n"+utils.GetDisplayString("Url", url, "Method",method,"Data", data,"Appid", appId, "appKey", appKey)
	SmsLog.Info(logStr)
	
	//curl
	resBody,resHeader := helper.CurlLeanCloud(url,method,data,appId,appKey);

	//log curl return
	logStr2 := "\nCurl leancloud for request code return\n"+utils.GetDisplayString("resBody", resBody,"resHeader",resHeader)
	SmsLog.Info(logStr2)

	return resBody
}

//valid a msm
func (u *MSms) ValidMsm(pkgName string,sms string,mobilePhoneNumber string,appId string,appKey string) map[string]interface{} {
	//检查本地数据库是否已保存对应的验证码
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_action FROM t_sms_action_valid WHERE F_action=? AND F_last_timestamp > ? LIMIT 1", helper.Md5(mobilePhoneNumber+pkgName+sms),helper.GetDateTimeBeforeMinute(smsValidMinute)).Values(&maps)
	if err == nil && num > 0 {
		return make(map[string]interface{},0)
	}

	url := "https://leancloud.cn/1.1/verifySmsCode/"+sms+"?mobilePhoneNumber="+mobilePhoneNumber
	method := "POST"
	data := map[string]string{}

	//log curl
	logStr := "\nCurl leancloud for verify code\n"+utils.GetDisplayString("Url", url, "Method",method,"Data", data,"Appid", appId, "appKey", appKey)
	SmsLog.Info(logStr)

	//向leancloud发送验证码验证请求
	resBody,resHeader := helper.CurlLeanCloud(url,method,data,appId,appKey);

	//log curl return
	logStr2 := "\nCurl leancloud for verify code return\n"+utils.GetDisplayString("resBody", resBody,"resHeader",resHeader)
	SmsLog.Info(logStr2)

	return resBody
}

//检查是否可以给用户发送短信了
func (u *MSms) CheckMsmRateValid(phone string,pkgName string)bool{
	o := orm.NewOrm()
	var maps []orm.Params
	nowTime := time.Now().Add(-time.Minute * time.Duration(SmsMinute)).Format("2006-01-02 15:04:05")
	num, err := o.Raw("SELECT F_last_timestamp FROM t_sms_rate WHERE F_action=? LIMIT 1", helper.Md5(phone+pkgName)).Values(&maps)
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

//写入短信发送频率表
func (u *MSms) AddMsmRate(phone string,pkgName string){
	//写入数据库
	o := orm.NewOrm()
	o.Raw("replace into t_sms_rate(F_action,F_last_timestamp) values('"+helper.Md5(phone+pkgName)+"','"+time.Now().Format("2006-01-02- 15:04:05")+"')").Exec()
}

//删除短信发送频率表
func (u *MSms) DeleteMsmRate(phone string,pkgName string){
	o := orm.NewOrm()
	o.Raw("UPDATE t_sms_rate SET F_last_timestamp='1001-01-01 00:00:00' WHERE F_action=?",helper.Md5(phone+pkgName)).Exec()
}

//写入t_sms_action_valid
func (u *MSms) AddMsmActionvalid(phone string,pkgName string,sms string){
	//写入数据库
	o := orm.NewOrm()
	o.Raw("replace into t_sms_action_valid(F_action,F_last_timestamp) values('"+helper.Md5(phone+pkgName+sms)+"','"+helper.GetNowDateTime()+"')").Exec()
}

//删除t_sms_action_valid
func (u *MSms) DeleteMsmActionvalid(phone string,pkgName string,sms string){
	o := orm.NewOrm()
	o.Raw("DELETE FROM t_sms_action_valid WHERE F_action=?",helper.Md5(phone+pkgName+sms)).Exec()
}

//验证t_sms_action_valid
func (u *MSms) CheckMsmActionvalid(phone string,pkgName string,sms string)bool{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_action FROM t_sms_action_valid WHERE F_action=? AND F_last_timestamp > ? LIMIT 1", helper.Md5(phone+pkgName+sms),helper.GetDateTimeBeforeMinute(smsValidMinute)).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}