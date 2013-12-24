package controllers

import (
//"dqs/dao"
//"github.com/astaxie/beego"
//log "github.com/cihub/seelog"
)

type HelpController struct {
	BaseController
}

func (this *HelpController) About() {
	this.Data["title"] = "关于本系统"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()

	this.TplNames = "about.html"
	this.Render()
}

func (this *HelpController) Help() {
	this.Data["title"] = "帮助"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()

	this.TplNames = "help.html"
	this.Render()
}
