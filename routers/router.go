// @APIVersion 2
// @Title 用户系统 API v2
package routers

import (
	"dream_api_sms_v2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v2",
		beego.NSNamespace("/sms",
			beego.NSInclude(
				&controllers.SmsController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/consumer",
			beego.NSInclude(
				&controllers.ConsumerController{},
			),
		),
		beego.NSNamespace("/area",
			beego.NSInclude(
				&controllers.AreaController{},
			),
		),
		beego.NSNamespace("/school",
			beego.NSInclude(
				&controllers.SchoolController{},
			),
		),
		beego.NSNamespace("/class",
			beego.NSInclude(
				&controllers.ClassController{},
			),
		),
		beego.NSNamespace("/token",
			beego.NSInclude(
				&controllers.TokenController{},
			),
		),
		beego.NSNamespace("/tmp",
			beego.NSInclude(
				&controllers.TmpController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
