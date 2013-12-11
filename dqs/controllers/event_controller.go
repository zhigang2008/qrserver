package controllers

import (
	"dqs/dao"
	"dqs/models"
	"dqs/util"
	"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
	"strconv"
	"time"
)

type EventController struct {
	BaseController
}

//事件列表
func (this *EventController) EventPageList() {
	this.Data["title"] = "震情事件"
	this.Data["author"] = "wangzhigang"
	//权限检查
	//this.AuthRoles("role_admin")
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
		pagination.PageSize = 10
	} else {
		pagination.PageSize = int(pagesize)
	}

	//查询参数

	eventid := this.GetString("eventid")
	if eventid != "" {
		pagination.AddParams("eventid", eventid)
	}
	begintime := this.GetString("begintime")
	if begintime != "" {
		pagination.AddParams("begintime", begintime)
	} else {
		now := time.Now()
		pagination.AddParams("begintime", now.Format(dao.EventTimeLayout))
	}
	endtime := this.GetString("endtime")
	if endtime != "" {
		pagination.AddParams("endtime", endtime)
	}

	//执行查询
	err = dao.EventPageList(&pagination)
	if err != nil {
		log.Warnf("查询震情事件列表失败:%s", err.Error())
	}
	pagination.Compute()

	this.Data["pagedata"] = pagination
	this.TplNames = "eventlist.html"
	this.Render()
}

//确认事件列表
func (this *EventController) EventSignalPageList() {
	this.Data["title"] = "地震事件列表"
	this.Data["author"] = "wangzhigang"
	//权限检查
	//this.AuthRoles("role_admin")
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
		pagination.PageSize = 10
	} else {
		pagination.PageSize = int(pagesize)
	}

	//查询参数

	signalid := this.GetString("signalid")
	if signalid != "" {
		pagination.AddParams("signalid", signalid)
	}
	level, err3 := this.GetInt("level")
	if err3 == nil {
		pagination.AddParams("level", int(level))
	}

	begintime := this.GetString("begintime")
	if begintime != "" {
		pagination.AddParams("begintime", begintime)
	} else {
		now := time.Now()
		pagination.AddParams("begintime", now.Format(dao.EventTimeLayout))
	}
	endtime := this.GetString("endtime")
	if endtime != "" {
		pagination.AddParams("endtime", endtime)
	}

	//执行查询
	err = dao.EventSignalPageList(&pagination)
	if err != nil {
		log.Warnf("查询地震事件列表失败:%s", err.Error())
	}
	pagination.Compute()

	this.Data["pagedata"] = pagination
	this.TplNames = "eventsignallist.html"
	this.Render()
}

//根据地震事件构造等值线
func (this *EventController) EventLine() {
	this.Data["title"] = "等值线"
	this.Data["author"] = "wangzhigang"
	//权限检查
	//this.AuthRoles("role_admin")
	this.CheckUser()

	eventid := this.GetString(":id")
	var eventSignal models.EventSignal
	var dataArray *[]models.NetGrid
	//查找当前事件
	event, err0 := dao.GetEventById(eventid)
	if err0 != nil {
		log.Warnf("查找事件[%s]失败:%s", eventid, err0.Error())
	}
	if event.IsConfirm {
		eventSignal, err0 = dao.GetEventSignalById(event.SignalId)
	}
	alarms, err := dao.GetAlarmsByEventId(eventid)
	if err != nil {
		log.Warnf("查找等值线的报警数据时出错:%s", err.Error())
	}

	//是否加入网格化虚拟站点
	if eventSignal.Id != "" {
		dataArray = NetGridCompute(alarms, eventSignal)
	}

	DataArrayStr := ""
	for k, v := range *dataArray {
		if k == len(*dataArray) {
			DataArrayStr += strconv.FormatFloat(float64(v.Longitude), 'f', -1, 32) + "-" + strconv.FormatFloat(float64(v.Latitude), 'f', -1, 32) + "-" + strconv.FormatFloat(float64(v.Value), 'f', -1, 32)

		} else {
			DataArrayStr += strconv.FormatFloat(float64(v.Longitude), 'f', -1, 32) + "-" + strconv.FormatFloat(float64(v.Latitude), 'f', -1, 32) + "-" + strconv.FormatFloat(float64(v.Value), 'f', -1, 32) + ","
		}
	}
	//添加系统参数
	this.Data["dataArray"] = DataArrayStr
	this.Data["dataSize"] = len(*dataArray)

	usegis := false
	usegis, err = beego.AppConfig.Bool("map_gis")
	if err != nil {
		usegis = false
		log.Warnf("无法从配置文件中获取gis启用信息.将使用地图模式.")
	}
	if usegis {
		this.Data["gisServiceUrl"] = beego.AppConfig.String("gis_service_url")
		this.Data["gisServiceParams"] = beego.AppConfig.String("gis_service_params")
		this.Data["gisBasicLayer"] = beego.AppConfig.String("gis_layer_basic")
	}
	this.TplNames = "index-gis.html"
	this.Render()

}

func NetGridCompute(alarms *[]models.AlarmInfo, eventSignal models.EventSignal) (ng *[]models.NetGrid) {
	return nil
}
