package controllers

import (
	//"dqs/dao"
	//"dqs/util"
	"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
)

type AnalyzeController struct {
	BaseController
}

//报警信息列表
func (this *AnalyzeController) Get() {
	this.Data["title"] = "数据分析"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()

	usegis, err := beego.AppConfig.Bool("map.gis")
	if err != nil {
		usegis = false
		log.Warnf("无法从配置文件中获取gis启用信息.将使用地图模式.")
	}
	if usegis {
		this.Data["gisServiceUrl"] = beego.AppConfig.String("gis.service.url")
		this.Data["gisBasicLayer"] = beego.AppConfig.String("gis.layer.basic")
		this.TplNames = "analyze-gis.html"
	} else {

		this.TplNames = "analyze.html"
	}
	this.Render()
}
