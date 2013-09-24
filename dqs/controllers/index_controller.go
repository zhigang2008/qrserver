package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["title"] = "首页"
	this.Data["author"] = "wangzhigang"
	this.TplNames = "index.html"
}
