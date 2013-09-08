package quickserver

import (
	"dqs/controllers"
	"github.com/astaxie/beego"
)

func RouteConfig() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/alarm", &controllers.AlarmController{})

	beego.RESTRouter("/device", &controllers.DeviceController{})

	beego.SetStaticPath("/logs", "logs")
}
