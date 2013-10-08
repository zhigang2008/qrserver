package controllers

import (
//"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["title"] = "首页"
	this.Data["author"] = "wangzhigang"

	this.CheckUser()
	this.TplNames = "index.html"
	this.Render()
}
