package quickserver

import (
	"dqs/controllers"
	"github.com/astaxie/beego"
	//"net/http"
)

func RouteConfig() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/device/params/:id", &controllers.DeviceController{}, "get:RefreshParams")
	beego.RESTRouter("/device", &controllers.DeviceController{})
	beego.Router("/alarm", &controllers.AlarmController{})
	beego.Router("/loglist", &controllers.LogsController{})
	beego.Router("/analyze", &controllers.AnalyzeController{})
	beego.Router("/report", &controllers.ReportController{})

	//静态文件路径
	beego.SetStaticPath("/logs", "logs")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/font", "static/font")
}
