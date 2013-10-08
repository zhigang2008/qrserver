package controllers

import (
	"dqs/dao"
	"dqs/util"
	//"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
)

type AlarmController struct {
	BaseController
}

//报警信息列表
func (this *AlarmController) Get() {
	this.Data["title"] = "报警信息"
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

	sid := this.GetString("sensorid")
	if sid != "" {
		pagination.AddParams("sensorid", sid)
	}

	//执行查询
	err = dao.AlarmList(&pagination)
	if err != nil {
		log.Warnf("查询报警信息失败:%s", err.Error())
	}
	pagination.Compute()

	this.Data["pagedata"] = pagination
	this.TplNames = "alarmlist.html"
	this.Render()

}
