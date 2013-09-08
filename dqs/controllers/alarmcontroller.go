package controllers

import (
	"dqs/models"
	"dqs/util"
	"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
)

type AlarmController struct {
	beego.Controller
}

func (this *AlarmController) Get() {
	this.Data["title"] = "报警信息"
	this.Data["author"] = "wangzhigang"
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

	//执行查询
	err = models.AlarmList(&pagination)
	if err != nil {
		log.Warnf("查询报警信息失败:%s", err.Error())
	}
	pagination.Compute()

	this.Data["pagedata"] = pagination
	this.TplNames = "alarmlist.html"
}
