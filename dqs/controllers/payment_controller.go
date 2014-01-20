package controllers

import (
	"dqs/dao"
	//"dqs/models"
	"dqs/util"
	log "github.com/cihub/seelog"
	//"time"
)

type FeeController struct {
	BaseController
}

//用户登录
func (this *FeeController) GetPayments() {
	this.Data["title"] = "待充值设备信息"
	this.Data["author"] = "wangzhigang"
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

	SensorId := this.GetString("SensorId")
	if SensorId != "" {
		pagination.AddParams("SensorId", SensorId)
	}
	NetOperator := this.GetString("NetOperator")
	if NetOperator != "" {
		pagination.AddParams("NetOperator", NetOperator)
	}
	NetNo := this.GetString("NetNo")
	if NetNo != "" {
		pagination.AddParams("NetNo", NetNo)
	}
	begintime := this.GetString("begintime")
	if begintime != "" {
		pagination.AddParams("begintime", begintime)
	}
	endtime := this.GetString("endtime")
	if endtime != "" {
		pagination.AddParams("endtime", endtime)
	}

	//执行查询
	err = dao.GetPagePayments(&pagination)
	if err != nil {
		log.Warnf("查询待充值信息失败:%s", err.Error())
	}
	pagination.Compute()

	this.Data["pagedata"] = pagination
	this.TplNames = "payment.html"
	this.Render()
}

func (this *FeeController) GetPaymentHistory() {
}
