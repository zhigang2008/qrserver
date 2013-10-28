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
	usegis, err = beego.AppConfig.Bool("map.gis")
	if err != nil {
		usegis = false
		log.Warnf("无法从配置文件中获取gis启用信息.将使用地图模式.")
	}
	if usegis {
		this.TplNames = "index-gis.html"
	} else {
		this.TplNames = "index.html"
	}

	this.Render()
}
