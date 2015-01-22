package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SmsController"],
		beego.ControllerComments{
			"Smsvalid",
			`/smsvalid/:mobilePhoneNumber`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SmsController"],
		beego.ControllerComments{
			"RegisterGetSms",
			`/register/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SmsController"],
		beego.ControllerComments{
			"ResetPwdGetSms",
			`/resetpwd/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SmsController"],
		beego.ControllerComments{
			"FindPwdGetSms",
			`/pwd/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:TmpController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:TmpController"],
		beego.ControllerComments{
			"DeleteAllUser",
			`/alluser`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:TmpController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:TmpController"],
		beego.ControllerComments{
			"DeleteUser",
			`/user`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"],
		beego.ControllerComments{
			"Register",
			`/register`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"],
		beego.ControllerComments{
			"ResetPwd",
			`/resetpwd`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"],
		beego.ControllerComments{
			"CheckUserAndPwd",
			`/login/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"],
		beego.ControllerComments{
			"FindPwd",
			`/pwd/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"],
		beego.ControllerComments{
			"ModifyPwd",
			`/pwd/:mobilePhoneNumber`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"],
		beego.ControllerComments{
			"CheckUserExists",
			`/exists/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"],
		beego.ControllerComments{
			"GetUserInfo",
			`/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

}
