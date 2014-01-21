package httpserver

import (
	"dqs/controllers"
	"github.com/astaxie/beego"
	//"net/http"
)

func routeConfig() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/loglist", &controllers.LogsController{})
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

	//充值管理
	beego.Router("/payment", &controllers.FeeController{}, "*:GetPayments")
	beego.Router("/paymenthistory", &controllers.FeeController{}, "*:GetPaymentHistory")
	beego.Router("/payment/recharge", &controllers.FeeController{}, "*:Recharge")
	//beego.Router("/payment/refresh", &controllers.FeeController{}, "*:RefreshAfterRecharge")

	//震情事件
	beego.Router("/analyze", &controllers.AnalyzeController{})
	beego.Router("/eventlist", &controllers.EventController{}, "*:EventPageList")
	beego.Router("/eventjsonlist", &controllers.EventController{}, "*:EventJsonList")
	beego.Router("/eventsignallist", &controllers.EventController{}, "*:EventSignalPageList")
	beego.Router("/eventline/:id", &controllers.EventController{}, "*:EventLine")
	beego.Router("/eventlinejson/:id", &controllers.EventController{}, "*:EventLineJson")
	//传入地震事件
	beego.Router("/earthquake", &controllers.EventController{}, "*:AddEventSignal")
	beego.Router("/quakelocation/:id", &controllers.EventController{}, "*:QuakeLocation")
	//速报管理
	beego.Router("/reportlist", &controllers.ReportController{}, "*:Get")
	beego.Router("/report/:id", &controllers.ReportController{}, "*:GetReport")
	beego.Router("/reportinvalid/:id", &controllers.ReportController{}, "*:SetInvalid")
	beego.Router("/reportverify/:id", &controllers.ReportController{}, "*:SetVerify")
	beego.Router("/reportverifyandsend/:id", &controllers.ReportController{}, "*:SetVerifyAndSend")
	beego.Router("/reportdirectsend/:id", &controllers.ReportController{}, "*:DirectSend")
	beego.Router("/reportresend/:id", &controllers.ReportController{}, "*:ReSend")

	//用户管理
	beego.RESTRouter("/user", &controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{}, "*:ToUserAddPage")
	beego.Router("/resetpwd", &controllers.UserController{}, "*:ResetPassword")
	beego.Router("/toresetpwd/:uid", &controllers.UserController{}, "*:ToResetPassword")

	//用户自助
	beego.Router("/userinfo/view/:objectid", &controllers.UserSelfController{}, "*:View")
	beego.Router("/userinfo/update", &controllers.UserSelfController{}, "*:Update")
	beego.Router("/userinfo/resetpwd", &controllers.UserSelfController{}, "*:ResetPassword")

	//运行参数
	beego.RESTRouter("/runtimeconfig", &controllers.RuntimeConfigController{})
	//系统配置
	beego.RESTRouter("/systemconfig", &controllers.SysConfigController{})

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

	//关于与帮助页面
	beego.Router("/about", &controllers.HelpController{}, "*:About")
	beego.Router("/help", &controllers.HelpController{}, "*:Help")
	//测试页面
	beego.Router("/test", &controllers.TestController{}, "*:Test")

	//静态文件路径
	beego.SetStaticPath("/logs", "logs")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/font", "static/font")
	beego.SetStaticPath("/output", "output")

}
