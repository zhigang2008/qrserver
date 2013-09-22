package controllers

import (
	//"dqs/dao"
	//"dqs/util"
	"github.com/astaxie/beego"

//log "github.com/cihub/seelog"
)

type AnalyzeController struct {
	beego.Controller
}

//报警信息列表
func (this *AnalyzeController) Get() {
	this.Data["title"] = "数据分析"
	this.Data["author"] = "wangzhigang"
	this.TplNames = "analyze.html"
}
