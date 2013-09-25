package httpserver

import (
	"dqs/controllers"
	"github.com/astaxie/beego"
	//"net/http"
)

func routeConfig() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/alarm", &controllers.AlarmController{})
	beego.Router("/loglist", &controllers.LogsController{})
	beego.Router("/analyze", &controllers.AnalyzeController{})
	beego.Router("/report", &controllers.ReportController{})
	beego.RESTRouter("/device", &controllers.DeviceController{})

	//添加设备页面
	beego.Router("/device/add", &controllers.DeviceController{}, "*:ToDeviceAddPage")
	//设备参数
	beego.Router("/deviceparams/refresh/:id", &controllers.DeviceController{}, "*:RefreshDeviceParams")
	beego.Router("/deviceparams/update/:id", &controllers.DeviceController{}, "*:UpdateDeviceParams")
	//设备自定义参数
	beego.Router("/customparams/update/:id", &controllers.DeviceController{}, "*:UpdateCustomParams")

	//用户管理
	beego.RESTRouter("/user", &controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{}, "*:ToUserAddPage")

	//用户自助
	beego.Router("/userinfo/:objectid", &controllers.UserSelfController{}, "*:View")

	//登录管理
	beego.Router("/sign", &controllers.CommonController{}, "*:Sign")
	beego.Router("/signout", &controllers.CommonController{}, "*:SignOut")
	beego.Router("/login", &controllers.CommonController{}, "*:Login")

	//静态文件路径
	beego.SetStaticPath("/logs", "logs")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/font", "static/font")
}
