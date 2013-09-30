package controllers

import (
	"dqs/dao"
	"dqs/util"
	//"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
	"time"
)

type AuditController struct {
	BaseController
}

//审计日志信息列表
func (this *AuditController) List() {
	this.Data["title"] = "审计日志"
	this.Data["author"] = "wangzhigang"

	//权限检查
	this.AuthRoles("role_admin")

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
		pagination.PageSize = 20
	} else {
		pagination.PageSize = int(pagesize)
	}

	//查询参数

	acttype := this.GetString("acttype")
	if acttype != "" {
		pagination.AddParams("acttype", acttype)
	}
	userid := this.GetString("userid")
	if userid != "" {
		pagination.AddParams("userid", userid)
	}
	actcontent := this.GetString("actcontent")
	if actcontent != "" {
		pagination.AddParams("actcontent", actcontent)
	}
	begintime := this.GetString("begintime")
	if begintime != "" {
		pagination.AddParams("begintime", begintime)
	} else {
		now := time.Now()
		pagination.AddParams("begintime", now.Format(dao.AuditTimeLayout))
	}
	endtime := this.GetString("endtime")
	if begintime != "" {
		pagination.AddParams("endtime", endtime)
	}

	//执行查询
	err = dao.AuditList(&pagination)
	if err != nil {
		log.Warnf("查询审计日志失败:%s", err.Error())
	}
	pagination.Compute()

	this.Data["pagedata"] = pagination
	this.TplNames = "audit.html"
}
