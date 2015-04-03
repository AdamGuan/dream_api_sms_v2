package helper

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"fmt"
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	//"net/url"
	"time"
	"regexp"
	"math/rand"
//	"github.com/astaxie/beego/config" 
)

func init() {
}

//类型转化 string  to int
func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

//类型转化 string  to float64
func StrToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

//类型转化 int to string
func IntToString(i int) string {
	return fmt.Sprintf("%d", i)
}

//类型转化 int64 to string
func Int64ToString(i int64) string {
	return fmt.Sprintf("%d", i)
}

//获取16位Guid
func GetGuid() string {
	f, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x", b)
	return uuid[0:16]
}

//创建密码
func CreatePwd(num int) string {
	return GetGuid()[0:num]
}

//leanCloud curl
func CurlLeanCloud(requestUri string, method string, requestData map[string]string, appId string, appKey string) (map[string]interface{},map[string][]string) {
	geturl := requestUri
	req, _ := http.NewRequest(method, geturl, nil)
	data, _ := json.Marshal(requestData)
	req.Body = ioutil.NopCloser(strings.NewReader(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("User-Agent", "SSTS Browser/1.0")
	//req.Header.Add("X-AVOSCloud-Application-Id", "l7dn2hfzry3yxdtfhnhciy0jt00dqd3ysbees8pjdhonwjb4")
	req.Header.Add("X-AVOSCloud-Application-Id", appId)
	//req.Header.Add("X-AVOSCloud-Application-Key", "76nib8wzexxcgaccrl9e4955ll8q11w1jwq36ofpx6u340q2")
	req.Header.Add("X-AVOSCloud-Application-Key", appKey)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	p := map[string]interface{}{}
	json.Unmarshal(bodyByte, &p)
	return p,resp.Header
}

//检查签名
func CheckSign(sign string, token string) bool {
	//sign = timestamp+md5(token+timestamp)
	if len(sign) == 46 && len(token) == 32 {
		timestamp := sign[0:14]
		//检测是否超时
//		appConf, _ := config.NewConfig("ini", "conf/app.conf")
//		debug,_ := appConf.Bool(beego.RunMode+"::debug")
//		if !debug{
		if 1 != 1{
			nowTime, _ := strconv.Atoi(time.Now().Format("20060102150405"))
			requestTime, _ := strconv.Atoi(timestamp)
			timedistince := nowTime - requestTime
			if timedistince > 60*5 {
				return false
			}
		}

		md5Str := sign[14:]
		str := token + timestamp
		md5Str2 := fmt.Sprintf("%x\r\n", md5.Sum([]byte(str)))
		md5Str2 = strings.TrimSpace(md5Str2)
		if string(md5Str) == string(md5Str2) {
			return true
		}
	}
	return false
}

//手机号码有效性验证
func CheckMPhoneValid(phone string)bool{
	matched, err := regexp.MatchString("^1[3|4|5|6|7|8][0-9]{9}$", phone)
	if err == nil && matched{
		return true
	}
	return false
}

//密码有效性验证
func CheckPwdValid(pwd string)bool{
	//pwd = strings.TrimSpace(pwd)
	matched, err := regexp.MatchString("^\\w{6,40}$", pwd)
	if err == nil && matched{
		return true
	}
	return false
}

//md5
func Md5(str string)string{
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5Str
}

//get now datatime
func GetNowDateTime()string{
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetDateTimeBeforeMinute(num int)string{
	return time.Now().Add(-time.Minute * time.Duration(num)).Format("2006-01-02 15:04:05")
}

func GetDateTimeAfterMinute(num int)string{
	return time.Now().Add(time.Minute * time.Duration(num)).Format("2006-01-02 15:04:05")
}

func Split(str string,flag string)[]string{
	return strings.Split(str, ",")
}

func JoinString(list []string,flag string)string{
	result := ""
	if len(list) > 0{
		for _,v := range list{
			result += v+","
		}
		result = strings.Trim(result,",")
	}
	return result
}

func StringInArray(value string,list []string)bool{
	result := false
	for _,item := range list{
		if value == item{
			result = true
			break
		}
	}
	return result
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//检查昵称有效性
func CheckNickNameValid(nickName string)bool{
	matched, err := regexp.MatchString("^[\u4e00-\u9fa5a-zA-Z0-9]{0,20}$", nickName)
	if err == nil && matched{
		return true
	}
	return false
}

//检查真实名有效性
func CheckRealNameValid(realName string)bool{
	matched, err := regexp.MatchString("^[\u4e00-\u9fa5a-zA-Z0-9]{0,20}$", realName)
	if err == nil && matched{
		return true
	}
	return false
}

//检查email有效性
func CheckEmailValid(email string)bool{
	matched, err := regexp.MatchString("^(\\w)+(\\.\\w+)*@(\\w)+((\\.\\w+)+)$", email)
	if err == nil && matched{
		return true
	}
	return false
}

//生成一个短信验证码
func GetSmsNum(bit int)string{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := ""
	for i:=0;i<bit;i++{
		str = str+IntToString(r.Intn(9))
	}
	return str
}