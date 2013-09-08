package controllers

import (
	"dqs/models"
	"github.com/astaxie/beego"
)

type DeviceController struct {
	beego.Controller
}

func (this *DeviceController) Get() {
	this.Data["title"] = "设备信息"
	this.Data["author"] = "wangzhigang"
	this.Data["devices"] = models.DeviceList(20)
	this.TplNames = "device.html"
}
func (this *DeviceController) GetDeviceInfo() {
	sid := this.Ctx.Params[":id"]
	this.Data["title"] = "设备详细信息"
	this.Data["author"] = "wangzhigang"
	this.Data["device"] = models.GetDevice(sid)
	this.TplNames = "deviceinfo.html"
}
