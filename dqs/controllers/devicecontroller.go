package controllers

import (
	"dqs/models"
	"dqs/util"
	//"fmt"
	"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
)

type DeviceController struct {
	beego.Controller
}

func (this *DeviceController) Get() {
	sid := this.Ctx.Params[":objectId"]
	//单个设备查询
	if sid != "" {
		this.Data["title"] = "设备详细信息"
		this.Data["author"] = "wangzhigang"
		this.Data["device"] = models.GetDevice(sid)

		this.TplNames = "device.html"

	} else { //列表
		this.Data["title"] = "设备列表"
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
		err = models.DeviceList(&pagination)
		if err != nil {
			log.Warnf("查询所有设备信息失败:%s", err.Error())
		}
		pagination.Compute()

		this.Data["pagedata"] = pagination
		this.TplNames = "devicelist.html"
	}

}
