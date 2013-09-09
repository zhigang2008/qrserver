package quickserver

import (
	"dqs/controllers"
	"github.com/astaxie/beego"
	"net/http"
)

func RouteConfig() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/alarm", &controllers.AlarmController{})
	beego.RouterHandler("/logs/", http.StripPrefix("/logs/", http.FileServer(http.Dir("./logs"))))
	beego.RESTRouter("/device", &controllers.DeviceController{})

	beego.SetStaticPath("/logs", "logs")
}
