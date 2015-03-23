package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/config" 
	"encoding/json"
)

var ErrLog *logs.BeeLogger
var DebugLog *logs.BeeLogger
var LogLog *logs.BeeLogger
var SmsLog *logs.BeeLogger
var EmailLog *logs.BeeLogger

type MLog struct {
}

func init() {
	otherConf, _ := config.NewConfig("ini", "conf/other.conf")
	errLogFile := otherConf.String("errLogFile")
	debugLogFile := otherConf.String("debugLogFile")
	logLogFile := otherConf.String("logLogFile")
	smsLogFile := otherConf.String("smsLogFile")
	emailLogFile := otherConf.String("emailLogFile")

	ErrLog = logs.NewLogger(10000)
	str,_ := json.Marshal(map[string]string{"filename":errLogFile})
	ErrLog.SetLogger("file", string(str))

	DebugLog = logs.NewLogger(10000)
	str,_ = json.Marshal(map[string]string{"filename":debugLogFile})
	DebugLog.SetLogger("file", string(str))

	LogLog = logs.NewLogger(10000)
	str,_ = json.Marshal(map[string]string{"filename":logLogFile})
	LogLog.SetLogger("file", string(str))

	SmsLog = logs.NewLogger(10000)
	str,_ = json.Marshal(map[string]string{"filename":smsLogFile})
	SmsLog.SetLogger("file", string(str))

	EmailLog = logs.NewLogger(10000)
	str,_ = json.Marshal(map[string]string{"filename":emailLogFile})
	EmailLog.SetLogger("file", string(str))

}

//记录请求(log等级)
func (u0 *MLog) LogRequest(u *context.Context) {
	if Debug{
		//log
		str := "\nRquest\n"+utils.GetDisplayString("IP", u.Input.IP(), "Scheme",u.Input.Scheme(),"Uri", u.Input.Uri(),"Method", u.Input.Method(), "Params", u.Input.Params, "Post", u.Input.Request.Form,"Header", u.Input.Request.Header)

		LogLog.Info(str)
	}
}

//记录返回(log等级)
func (u0 *MLog) LogEcho(datas map[string]interface{}) {
	if Debug{
		//log
		str := "\nEcho\n"+utils.GetDisplayString("datas", datas)

		LogLog.Info(str)
	}
}

//记录请求(500 err等级)
func (u0 *MLog) LogRequestErr500(u *context.Context,code string) {
	if DebugErrlog{
		//log
		str := "\n"+code+"\n"+utils.GetDisplayString("IP", u.Input.IP(), "Scheme",u.Input.Scheme(),"Uri", u.Input.Uri(),"Method", u.Input.Method(), "Params", u.Input.Params, "Post", u.Input.Request.Form,"Header", u.Input.Request.Header)

		ErrLog.Info(str)
	}
}