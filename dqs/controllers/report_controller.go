package controllers

import (
	//"dqs/models"
	"dqs/dao"
	//"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
)

type ReportController struct {
	BaseController
}

//速报管理
func (this *ReportController) Get() {
	this.Data["title"] = "速报管理"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()

	reports, err := dao.ReportList(6)
	if err != nil {
		log.Infof("查询速报列表出错:%s", err.Error())
	}

	if len(*reports) >= 1 {
		this.Data["latestReport"] = (*reports)[0]
	}
	if len(*reports) > 1 {
		this.Data["reports"] = (*reports)[1:]
	}

	this.TplNames = "report.html"
	this.Render()
}
