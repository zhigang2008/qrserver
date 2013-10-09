package controllers

import (
	"dqs/dao"
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
	this.TplNames = "index.html"
	this.Render()
}
