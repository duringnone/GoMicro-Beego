package main

import (
	"gateway/controllers"
	_ "gateway/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	. "github.com/duringnone/microbase"
	_ "github.com/go-sql-driver/mysql"
	_ "net/http/pprof"
)

func main() {
	// pprof模式是否开启
	PprofOn, _ := ConfigInfo.Bool("PprofOn")
	//if beego.BConfig.RunMode == "dev" && true == PprofOn {
	if true == PprofOn && "DEV" == RunMode {
		beego.Router("/debug/pprof", &controllers.PprofCtrl{})
		beego.Router(`/debug/pprof/:app([\w]+)`, &controllers.PprofCtrl{})
	}
	//if beego.BConfig.RunMode == "dev" {
	if RunMode == "DEV" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	// 解决跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		//AllowOrigins:     []string{"*.test.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.Run()
}
