package controllers

import (
	"dqs/models"
	"github.com/astaxie/beego"
)

type AlarmController struct {
	beego.Controller
}

func (this *AlarmController) Get() {
	this.Data["title"] = "报警信息"
	this.Data["author"] = "wangzhigang"
	this.Data["alarminfos"] = models.AlarmList(10)
	this.TplNames = "alarm.html"
}
