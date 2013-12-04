package controllers

import (
	"dqs/dao"
	"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["title"] = "首页"
	this.Data["author"] = "wangzhigang"

	this.CheckUser()

	devices, err := dao.GetAllDevices()
	if err != nil {
		log.Warnf("获取所有设备列表失败:%s", err.Error())
	}
	this.Data["devices"] = devices

	usegis := false
	usegis, err = beego.AppConfig.Bool("map_gis")
	if err != nil {
		usegis = false
		log.Warnf("无法从配置文件中获取gis启用信息.将使用地图模式.")
	}
	if usegis {
		this.Data["gisServiceUrl"] = beego.AppConfig.String("gis_service_url")
		this.Data["gisBasicLayer"] = beego.AppConfig.String("gis_layer_basic")
		this.TplNames = "index-gis.html"
	} else {
		this.TplNames = "index.html"
	}

	//计算设备信息
	deviceNum := len(devices)
	pages := 1
	if deviceNum%10 == 0 {
		pages = deviceNum / 10
	} else {
		pages = deviceNum/10 + 1
	}

	this.Data["devicePages"] = pages
	this.Render()
}
