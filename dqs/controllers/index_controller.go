package controllers

import (
	"dqs/dao"
	"dqs/util"
	//"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["title"] = "首页"
	this.Data["author"] = "wangzhigang"

	this.CheckUser()

	paginationDevices := util.Pagination{PageSize: 10, CurrentPage: 1}
	err := dao.DeviceList(&paginationDevices)
	if err != nil {
		log.Warnf("查询设备信息失败:%s", err.Error())
	}
	paginationDevices.Compute()
	this.Data["devices"] = paginationDevices.Data
	this.Data["devicePages"] = paginationDevices.PageCount
	/*devices, err := dao.GetAllDevices()
	if err != nil {
		log.Warnf("获取所有设备列表失败:%s", err.Error())
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
	this.Data["devices"] = devices
	*/
	paginationEvents := util.Pagination{PageSize: 10, CurrentPage: 1}
	err = dao.EventPageList(&paginationEvents)
	if err != nil {
		log.Warnf("查询震情事件列表失败:%s", err.Error())
	}
	paginationEvents.Compute()
	this.Data["events"] = paginationEvents.Data
	this.Data["eventPages"] = paginationEvents.PageCount

	allDevices, err := dao.GetAllDevices()
	if err != nil {
		log.Warnf("查询所有设备信息失败:%s", err.Error())
	}
	this.Data["allDevices"] = allDevices

	queryTimeSpan := RuntimeConfigs.IndexQueryTimeSpan
	if queryTimeSpan <= 0 {
		queryTimeSpan = 1
	}
	this.Data["queryTimeSpan"] = queryTimeSpan
	this.Data["GoogleMap"] = !SystemConfigs.DisableGoogleMap
	this.Data["mapExtend"] = SystemConfigs.GisImageCfg.BBOX

	usegis := SystemConfigs.UseGis
	if usegis {
		this.Data["gisServiceUrl"] = SystemConfigs.GisServiceUrl
		this.Data["gisServiceParams"] = SystemConfigs.GisServiceParams
		this.Data["gisBasicLayer"] = SystemConfigs.GisLayerBasic
		this.Data["gisChinaLayer"] = SystemConfigs.GisLayerChina
		this.TplNames = "index.html"
	} else {
		this.TplNames = "index-nogis.html"
	}

	hasPayment := false
	//续费通知
	if this.IsAuthRoles("role_admin", "role_device") {
		paymentCount := dao.GetPaymentCount()
		if paymentCount > 0 {
			this.Data["paymentCount"] = paymentCount
			hasPayment = true
		}
	}
	this.Data["hasPayment"] = hasPayment

	this.Render()
}
