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

	this.TplNames = "reportlist.html"
	this.Render()
}

//查询速报
func (this *ReportController) GetReport() {
	this.Data["title"] = "速报内容"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()

	sid := this.GetString(":id")
	report, err := dao.GetReportById(sid)
	if err != nil {
		log.Infof("查询速报出错:%s", err.Error())
	}
	this.Data["report"] = report

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
		this.AuditLog("设置速报无效["+sid+"]", true)
	}
	this.Data["json"] = &answer
	this.ServeJson()
}

//速报审批通过
func (this *ReportController) SetVerify() {
	sid := this.GetString(":id")

	err := dao.ReportVerify(sid)
	answer := JsonAnswer{}
	if err != nil {
		answer.Ok = false
		answer.Msg = "置审核通过失败:" + err.Error()
		log.Warnf("设置速报审核通过失败[%s]:%s", sid, err.Error())

	} else {
		answer.Ok = true
		answer.Msg = "成功"
		log.Infof("设置速报审核通过[%s]", sid)
		this.AuditLog("审核速报["+sid+"]通过", true)
	}
	this.Data["json"] = &answer
	this.ServeJson()
}

//直接发送
func (this *ReportController) SetVerifyAndSend() {
	sid := this.GetString(":id")
	answer := JsonAnswer{}

	report, err := dao.GetReportById(sid)
	if err != nil {
		answer.Ok = false
		answer.Msg = "获取速报信息失败:" + err.Error()
		log.Warnf("获取速报信息失败[%s]:%s", sid, err.Error())

	} else {
		err1 := dao.ReportVerify(sid)
		if err1 != nil {
			answer.Ok = false
			answer.Msg = "置审核通过失败:" + err1.Error()
			log.Warnf("设置速报审核通过失败[%s]:%s", sid, err1.Error())
		} else {
			//发送速报
			prepareMms(report)
			//更新速报发送状态
			updateReportSendStatus(&report)
			answer.Ok = true
			answer.Msg = "成功"
			log.Infof("审核并发送速报成功[%s]", sid)
			this.AuditLog("审核并发送速报["+sid+"]", true)
		}
	}

	this.Data["json"] = &answer
	this.ServeJson()
}

//直接发送
func (this *ReportController) DirectSend() {
	sid := this.GetString(":id")
	answer := JsonAnswer{}

	report, err := dao.GetReportById(sid)
	if err != nil {
		answer.Ok = false
		answer.Msg = "获取速报信息失败:" + err.Error()
		log.Warnf("获取速报信息失败[%s]:%s", sid, err.Error())

	} else {

		//发送速报
		prepareMms(report)
		//更新速报发送状态
		updateReportSendStatus(&report)
		answer.Ok = true
		answer.Msg = "成功"
		log.Infof("发送速报成功[%s]", sid)
		this.AuditLog("直接发送速报"+sid+"]", true)

	}

	this.Data["json"] = &answer
	this.ServeJson()
}

//再次发送速报
func (this *ReportController) ReSend() {
	sid := this.GetString(":id")
	answer := JsonAnswer{}

	report, err := dao.GetReportById(sid)
	if err != nil {
		answer.Ok = false
		answer.Msg = "获取速报信息失败:" + err.Error()
		log.Warnf("获取速报信息失败[%s]:%s", sid, err.Error())

	} else {
		//发送速报
		prepareMms(report)
		answer.Ok = true
		answer.Msg = "成功"
		log.Infof("发送速报成功[%s]", sid)
		this.AuditLog("再次发送速报["+sid+"]", true)
	}

	this.Data["json"] = &answer
	this.ServeJson()
}

//速报列表
func (this *ReportController) ReportPageList() {
	this.Data["title"] = "速报列表"
	this.Data["author"] = "wangzhigang"
	//权限检查
	//this.AuthRoles("role_admin")
	this.CheckUser()

	pagination := util.Pagination{}
	page, err := this.GetInt("page")
	if err != nil {
		pagination.CurrentPage = 1
	} else {
		pagination.CurrentPage = int(page)
	}
	pagesize, err2 := this.GetInt("pagesize")
	if err2 != nil {
		pagination.PageSize = 10
	} else {
		pagination.PageSize = int(pagesize)
	}

	//查询参数

	eventid := this.GetString("eventid")
	if eventid != "" {
		pagination.AddParams("eventid", eventid)
	}
	begintime := this.GetString("begintime")
	if begintime != "" {
		pagination.AddParams("begintime", begintime)
	} else {
		now := time.Now()
		pagination.AddParams("begintime", now.Format(dao.EventTimeLayout))
	}
	endtime := this.GetString("endtime")
	if endtime != "" {
		pagination.AddParams("endtime", endtime)
	}

	//执行查询
	err = dao.ReportPageList(&pagination)
	if err != nil {
		log.Warnf("查询震情事件列表失败:%s", err.Error())
	}
	pagination.Compute()

	this.Data["pagedata"] = pagination
	this.TplNames = "reportmore.html"
	this.Render()
}
