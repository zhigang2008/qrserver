package controllers

import (
//"dqs/models"
//"dqs/util"
//"github.com/astaxie/beego"
//log "github.com/cihub/seelog"
)

type ReportController struct {
	BaseController
}

//速报管理
func (this *ReportController) Get() {
	this.Data["title"] = "速报管理"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()
	this.TplNames = "report.html"
	this.Render()
}
