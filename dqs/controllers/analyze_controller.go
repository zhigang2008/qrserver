package controllers

import (
	"dqs/dao"
	//"dqs/util"
	//"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
)

type AnalyzeController struct {
	BaseController
}

//报警信息列表
func (this *AnalyzeController) Get() {
	this.Data["title"] = "震情分析"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()

	/*
		usegis, err := beego.AppConfig.Bool("map_gis")
		if err != nil {
			usegis = false
			log.Warnf("无法从配置文件中获取gis启用信息.将使用地图模式.")
		}
		if usegis {
			this.Data["gisServiceUrl"] = beego.AppConfig.String("gis_service_url")
			this.Data["gisServiceParams"] = beego.AppConfig.String("gis_service_params")
			this.Data["gisBasicLayer"] = beego.AppConfig.String("gis_layer_basic")
			this.TplNames = "analyze-gis.html"
		} else {

			this.TplNames = "analyze.html"
		}
	*/

	eventList, err1 := dao.EventList(10)
	if err1 != nil {
		log.Warnf("查询震情事件列表失败:%s", err1.Error())
	}
	signalList, err2 := dao.EventSignalList(10)
	if err2 != nil {
		log.Warnf("查询震情确认信号列表失败:%s", err2.Error())
	}

	this.Data["eventList"] = eventList
	this.Data["eventSignalList"] = signalList

	//计算事件信息
	eventNum := len(*eventList)
	eventPages := 1
	if eventNum%10 == 0 {
		eventPages = eventNum / 10
	} else {
		eventPages = eventNum/10 + 1
	}
	this.Data["eventPages"] = eventPages

	//计算事件信号信息
	signalNum := len(*signalList)
	signalPages := 1
	if signalNum%10 == 0 {
		signalPages = signalNum / 10
	} else {
		signalPages = signalNum/10 + 1
	}

	this.Data["signalPages"] = signalPages

	this.TplNames = "analyze.html"
	this.Render()
}
