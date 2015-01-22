package main

import (
	_ "dream_api_sms_v2/docs"
	_ "dream_api_sms_v2/routers"

	"github.com/astaxie/beego"
	"encoding/json"
	"io"
	"net/http"
	"runtime"
)

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	returndata := map[string]string{"responseCode": "404"}
	data, _ := json.Marshal(returndata)
	io.WriteString(rw, string(data))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if beego.RunMode == "dev"{
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Errorhandler("404", page_not_found)
	beego.Run()
}
