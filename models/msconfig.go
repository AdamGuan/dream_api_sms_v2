package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"os"
	//"dream_api_sms/helper"
	"time"
	"github.com/astaxie/beego/config" 
	"dream_api_sms_v2/helper"
)

//key:响应代码，value:响应信息
var ConfigMyResponse map[string]string

//key:id，value:名称
var City map[string]string
var County map[string]string
var Grade map[string]string
var Province map[string]string
var School map[string]string

//学校类型(小学,初中,高中)
var SchoolType map[int]string

//缺省学校列表
type DefaultSchoolItemType struct{
	F_school_id			int
	F_school			string
	F_school_type		int
	F_belong_area_id	int
}

type DefaultSchoolListType map[string]struct{
	F_school_id			int
	F_school			string
	F_school_type		int
	F_belong_area_id	int
}

var DefaultSchoolList DefaultSchoolListType

func init() {
	
	dbconf, _ := config.NewConfig("ini", "conf/db.conf")
	maxIdle,_ := dbconf.Int("maxIdle")
	maxConn,_ := dbconf.Int("maxConn")
	userName := dbconf.String(beego.RunMode+"::userName")
	password := dbconf.String(beego.RunMode+"::password")
	dbName := dbconf.String(beego.RunMode+"::dbName")
	orm.RegisterDataBase("default", "mysql", userName+":"+password+"@/"+dbName+"?charset=utf8&loc=Asia%2FShanghai",maxIdle, maxConn)
	orm.DefaultTimeLoc = time.UTC
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug,_ := appConf.Bool(beego.RunMode+"::debug")
	if debug{
		orm.Debug = true
	}
	logFile, _ := os.OpenFile("./db.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	orm.DebugLog = orm.NewLog(logFile)

	getResponseConfig()
	getCity()
	getCounty()
	getGrade()
	getProvince()
	getSchool()
	getDefaultSchool()
	
	SchoolType = make(map[int]string)
	SchoolType[1] = "小学"
	SchoolType[2] = "初中"
	SchoolType[3] = "高中"
}

//获取config  im
func getResponseConfig() {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM t_config_response").Values(&maps)
	if err == nil && num > 0 {
		ConfigMyResponse = make(map[string]string)
		for _, item := range maps {
			ConfigMyResponse[item["F_response_no"].(string)] = item["F_response_msg"].(string)
		}
	}
}

//获取city
func getCity() {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM t_area WHERE F_area_level = 2").Values(&maps)
	if err == nil && num > 0 {
		City = make(map[string]string)
		for _, item := range maps {
			City[item["F_area_id"].(string)] = item["F_area_name"].(string)
		}
	}
}

//获取county
func getCounty() {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM t_area WHERE F_area_level = 3").Values(&maps)
	if err == nil && num > 0 {
		County = make(map[string]string)
		for _, item := range maps {
			County[item["F_area_id"].(string)] = item["F_area_name"].(string)
		}
	}
}

//获取grade
func getGrade() {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM t_grade").Values(&maps)
	if err == nil && num > 0 {
		Grade = make(map[string]string)
		for _, item := range maps {
			Grade[item["F_grade_id"].(string)] = item["F_grade"].(string)
		}
	}
}

//获取province
func getProvince() {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM t_area WHERE F_area_level = 1").Values(&maps)
	if err == nil && num > 0 {
		Province = make(map[string]string)
		for _, item := range maps {
			Province[item["F_area_id"].(string)] = item["F_area_name"].(string)
		}
	}
}

//获取school
func getSchool() {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM t_school").Values(&maps)
	if err == nil && num > 0 {
		School = make(map[string]string)
		for _, item := range maps {
			School[item["F_school_id"].(string)] = item["F_school"].(string)
		}
	}
}

//获取缺省的school列表
func getDefaultSchool() {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM t_school WHERE F_school_id < 0").Values(&maps)
	if err == nil && num > 0 {
		DefaultSchoolList = make(DefaultSchoolListType,num)
		for _, item := range maps {
			DefaultSchoolList[item["F_school_id"].(string)] = DefaultSchoolItemType{F_school_id:helper.StrToInt(item["F_school_id"].(string)),F_school:item["F_school"].(string),F_school_type:helper.StrToInt(item["F_school_type"].(string)),F_belong_area_id:helper.StrToInt(item["F_belong_area_id"].(string))}
		}
	}
}