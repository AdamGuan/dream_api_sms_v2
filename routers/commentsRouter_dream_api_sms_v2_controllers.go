package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:AreaController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:AreaController"],
		beego.ControllerComments{
			"GetAllProvinces",
			`/provinces`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:AreaController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:AreaController"],
		beego.ControllerComments{
			"GetCitys",
			`/citys`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:AreaController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:AreaController"],
		beego.ControllerComments{
			"GetCountys",
			`/countys`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:AreaController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:AreaController"],
		beego.ControllerComments{
			"GetProvinceName",
			`/province/:ids`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:AreaController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:AreaController"],
		beego.ControllerComments{
			"GetCityName",
			`/city/:ids`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:AreaController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:AreaController"],
		beego.ControllerComments{
			"GetCountyName",
			`/county/:ids`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:ClassController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:ClassController"],
		beego.ControllerComments{
			"AddAClass",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:ClassController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:ClassController"],
		beego.ControllerComments{
			"GetAllClasses",
			`/:schoolId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SchoolController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SchoolController"],
		beego.ControllerComments{
			"QuerySchools",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SchoolController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SchoolController"],
		beego.ControllerComments{
			"GetAllGrade",
			`/grades`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SchoolController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SchoolController"],
		beego.ControllerComments{
			"GetSchoolArea",
			`/area/:schoolIds`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SchoolController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:SchoolController"],
		beego.ControllerComments{
			"GetSchoolName",
			`/name/:schoolIds`,
			[]string{"get"},
			nil})

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

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:TokenController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:TokenController"],
		beego.ControllerComments{
			"CheckToken",
			`/verify/:token`,
			[]string{"get"},
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

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"],
		beego.ControllerComments{
			"ModifyUserInfo",
			`/:mobilePhoneNumber`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"],
		beego.ControllerComments{
			"UserLogout",
			`/logout/:mobilePhoneNumber`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"] = append(beego.GlobalControllerRouter["dream_api_sms_v2/controllers:UserController"],
		beego.ControllerComments{
			"ModifyUserClass",
			`/class/:mobilePhoneNumber`,
			[]string{"put"},
			nil})

}
