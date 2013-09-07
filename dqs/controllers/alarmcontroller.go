package controllers

import (
	"github.com/astaxie/beego"
)

type AlarmController struct {
	beego.Controller
}

func (this *AlarmController) Get() {
	this.Data["title"] = "报警信息"
	this.Data["author"] = "wangzhigang"
	this.TplNames = "alarm.html"
}
