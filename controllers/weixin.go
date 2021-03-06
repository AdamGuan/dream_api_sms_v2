package controllers

import (
	"dream_api_sms_v2/helper"
	"dream_api_sms_v2/models"
)

//微信(第三方)
type WeixinController struct {
	BaseController
}

// @Title 登录
// @Description 登录(token: md5(pkg))
// @Param	access_token	query	string	true	access_token
// @Param	openid			query	string	true	openid
// @Param	gender			query	string	false	性别(值: [男|女])
// @Param	nickname		query	string	false	昵称
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MUserLoginResp
// @Failure 401 无权访问
// @router /login [get]
func (u *WeixinController) LoginWeixin() {
	//log
	u.logRequest()
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var userObj *models.MConsumer
	//parse request parames
	u.Ctx.Request.ParseForm()
	access_token := u.Ctx.Request.FormValue("access_token")
	openid := u.Ctx.Request.FormValue("openid")
	gender := u.Ctx.Request.FormValue("gender")
	nickname := u.Ctx.Request.FormValue("nickname")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//check sign
	datas["responseNo"] = u.checkSign2()
	//检查参数
	if datas["responseNo"] == 0 {
		datas["responseNo"] = -1
		//检查微信信息的有效性
		if len(access_token) > 0 && len(openid) > 0 {
			res, _ := helper.CurlQq("https://api.weixin.qq.com/sns/auth?access_token="+access_token+"&openid="+openid, "GET")
			resErrcode, ok := res["errcode"]
			if ok {
				resErrcode2 := resErrcode.(float64)
				if int(resErrcode2) == 0 {
					//检查微信号码是否已存在
					uid := userObj.GetUidByWeixin(openid)
					if len(uid) <= 0 {
						datas["responseNo"] = 0
						//写入一条微信数据
						paramesData := make(map[string]string)
						paramesData["qq"] = openid
						if len(gender) > 0 {
							if gender != "男" && gender != "女" {
								datas["responseNo"] = -10
							} else {
								gender2 := "1"
								if gender == "男" {
									gender2 = "1"
								} else {
									gender2 = "2"
								}
								paramesData["gender"] = gender2
							}
						}
						if len(nickname) > 0 {
							paramesData["nickname"] = nickname
						}
						if datas["responseNo"] == 0 {
							uid = userObj.InsertWeixin(paramesData, pkg)
						}
					}
					if len(uid) > 0 {
						//返回登录信息
						info := u.login(uid, pkg)
						if len(info) > 0 {
							datas["responseNo"] = 0
							for key, value := range info {
								datas[key] = value
							}
						}
					}
				}
			}
		} else {
			datas["responseNo"] = -10
		}
	}
	//return
	u.jsonEcho(datas)
}

//登录
func (u *WeixinController) login(uid string, pkg string) map[string]interface{} {
	userInfo := map[string]interface{}{}
	//model ini
	var userObj *models.MConsumer
	//检查uid是否存在
	if userObj.CheckUserIdExists(uid) {
		//获取token
		token, tokenExpireDatetime := userObj.GetTokenByUid(uid, pkg)
		//获取其它信息
		if len(token) > 0 {
			userInfo["token"] = token
			userInfo["tokenExpireDatetime"] = tokenExpireDatetime
			info := userObj.GetUserInfoByUid(uid)
			if len(info.F_uid) > 0 {
				userInfo["F_uid"] = info.F_uid
				userInfo["F_phone_number"] = info.F_phone_number
				userInfo["F_gender"] = info.F_gender
				userInfo["F_grade"] = info.F_grade
				userInfo["F_grade_id"] = info.F_grade_id
				userInfo["F_birthday"] = info.F_birthday
				userInfo["F_school"] = info.F_school
				userInfo["F_school_id"] = info.F_school_id
				userInfo["F_province"] = info.F_province
				userInfo["F_province_id"] = info.F_province_id
				userInfo["F_city"] = info.F_city
				userInfo["F_city_id"] = info.F_city_id
				userInfo["F_county"] = info.F_county
				userInfo["F_county_id"] = info.F_county_id
				userInfo["F_user_realname"] = info.F_user_realname
				userInfo["F_user_nickname"] = info.F_user_nickname
				userInfo["F_crate_datetime"] = info.F_crate_datetime
				userInfo["F_modify_datetime"] = info.F_modify_datetime
				userInfo["F_class_id"] = info.F_class_id
				userInfo["F_class_name"] = info.F_class_name
				userInfo["F_avatar_url"] = info.F_avatar_url
				userInfo["F_user_email"] = info.F_user_email
			}
		}
	}
	return userInfo
}
