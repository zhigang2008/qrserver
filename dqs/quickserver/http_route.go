package quickserver

import (
	"dqs/controllers"
	"github.com/astaxie/beego"
	//"net/http"
)

func RouteConfig() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/alarm", &controllers.AlarmController{})
	beego.Router("/loglist", &controllers.LogsController{})
	beego.RESTRouter("/device", &controllers.DeviceController{})

	beego.SetStaticPath("/logs", "logs")
}
