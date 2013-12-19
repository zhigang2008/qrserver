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

	reports, err := dao.GetValidReports(8)
	if err != nil {
		log.Infof("查询速报列表出错:%s", err.Error())
	}

	if len(*reports) > 3 {
		this.Data["newReports"] = (*reports)[0:3]
		this.Data["oldReports"] = (*reports)[3:]
	} else {
		this.Data["newReports"] = (*reports)
	}

	this.TplNames = "report.html"
	this.Render()
}

//速报无效
func (this *ReportController) SetInvalid() {
	sid := this.GetString(":id")

	err := dao.ReportInvalid(sid)
	answer := JsonAnswer{}
	if err != nil {
		answer.Ok = false
		answer.Msg = "置无效失败:" + err.Error()
		log.Warnf("设置速报无效失败[%s]:%s", sid, err.Error())

	} else {
		answer.Ok = true
		answer.Msg = "成功"
		log.Infof("设置速报无效[%s]", sid)
	}
	this.Data["json"] = &answer
	this.ServeJson()
}
