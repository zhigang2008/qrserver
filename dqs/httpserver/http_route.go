package httpserver

import (
	"dqs/controllers"
	"github.com/astaxie/beego"
	//"net/http"
)

func routeConfig() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/loglist", &controllers.LogsController{})
	beego.Router("/analyze", &controllers.AnalyzeController{})
	beego.Router("/report", &controllers.ReportController{})
	beego.RESTRouter("/device", &controllers.DeviceController{})

	//报警
	beego.Router("/alarm", &controllers.AlarmController{}, "*:Get")
	beego.Router("/realtime_alarm", &controllers.AlarmController{}, "*:GetRealtimeAlarm")
	//波形图展示
	//beego.Router("/waveinfo/:objectid", &controllers.AlarmController{}, "*:ShowWaveInfoById")
	beego.Router("/waveinfo/:sid/:seqno", &controllers.AlarmController{}, "*:ShowWaveInfo")
	beego.Router("/getwaveinfo/:sid/:seqno", &controllers.AlarmController{}, "*:GetWaveInfo")
	//添加设备页面
	beego.Router("/adddevice", &controllers.DeviceController{}, "*:ToDeviceAddPage")
	//获取设备列表json
	beego.Router("/listdevice", &controllers.DeviceController{}, "*:DeviceList")
	//设备定位
	beego.Router("/location/:id", &controllers.DeviceController{}, "*:DeviceLocation")
	//设备参数
	beego.Router("/deviceparams/refresh", &controllers.DeviceController{}, "*:RefreshDeviceParams")
	beego.Router("/deviceparams/update", &controllers.DeviceController{}, "*:UpdateDeviceParams")
	//设备自定义参数
	beego.Router("/customparams/update/:id", &controllers.DeviceController{}, "*:UpdateCustomParams")

	//震情事件
	beego.Router("/eventlist", &controllers.EventController{}, "*:PageList")

	//用户管理
	beego.RESTRouter("/user", &controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{}, "*:ToUserAddPage")
	beego.Router("/resetpwd", &controllers.UserController{}, "*:ResetPassword")
	beego.Router("/toresetpwd/:uid", &controllers.UserController{}, "*:ToResetPassword")

	//用户自助
	beego.Router("/userinfo/view/:objectid", &controllers.UserSelfController{}, "*:View")
	beego.Router("/userinfo/update", &controllers.UserSelfController{}, "*:Update")
	beego.Router("/userinfo/resetpwd", &controllers.UserSelfController{}, "*:ResetPassword")

	//审计日志
	beego.Router("/audit", &controllers.AuditController{}, "*:List")
	//登录管理
	beego.Router("/sign", &controllers.CommonController{}, "*:Sign")
	beego.Router("/signout", &controllers.CommonController{}, "*:SignOut")
	beego.Router("/login", &controllers.CommonController{}, "*:Login")
	beego.Router("/register", &controllers.CommonController{}, "*:Register")
	beego.Router("/registersave", &controllers.CommonController{}, "*:RegisterSave")
	//数据交换
	beego.Router("/data/alarm", &controllers.ExchangeController{}, "*:ExportAlarms")
	//测试页面
	beego.Router("/testgis", &controllers.TestController{}, "*:TestGis")
	beego.Router("/test", &controllers.TestController{}, "*:Test")

	//静态文件路径
	beego.SetStaticPath("/logs", "logs")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/font", "static/font")

}
