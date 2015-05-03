package controllers

import (
	"dqs/dao"
	"dqs/models"
	"dqs/util"
	"encoding/xml"
	//"github.com/astaxie/beego"
	"bytes"
	log "github.com/cihub/seelog"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	EarthQuakeTimeLayout = "20060102150405.000" //地震时间格式
)

type EventController struct {
	BaseController
}

//事件列表
func (this *EventController) EventPageList() {
	this.Data["title"] = "仪器观测事件"
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

//事件列表Json
func (this *EventController) EventJsonList() {

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
	/*
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
	*/
	//执行查询
	err = dao.EventPageList(&pagination)
	if err != nil {
		log.Warnf("查询震情事件列表失败:%s", err.Error())
	}
	pagination.Compute()

	this.Data["json"] = pagination.Data
	this.ServeJson()
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
	var eventSignal models.EventSignal = models.EventSignal{}
	//var dataArray []models.NetGrid
	//查找当前事件
	event, err0 := dao.GetEventById(eventid)
	if err0 != nil {
		log.Warnf("查找事件[%s]失败:%s", eventid, err0.Error())
	}
	if event.IsConfirm {
		eventSignal, err0 = dao.GetEventSignalById(event.SignalId)
	}

	//查找报警数据
	alarms, err := dao.GetAlarmsByEventId(eventid)
	if err != nil {
		log.Warnf("查找等值线的报警数据时出错:%s", err.Error())
	}

	//是否加入网格化虚拟站点
	dataArray := NetGridCompute(alarms, eventSignal)

	//传递的数据值
	DataArrayStr := ""
	DataArrayStrPGA := ""
	DataArrayStrSI := ""
	var lastlng, lastlat float32

	for k, v := range dataArray {
		if k < len(dataArray)-1 {
			DataArrayStr += v.String() + ","
			DataArrayStrPGA += v.StringPGA() + ","
			DataArrayStrSI += v.StringSI() + ","
		} else {
			DataArrayStr += v.String()
			DataArrayStrPGA += v.StringPGA()
			DataArrayStrSI += v.StringSI()
		}
		//设置地图中心点位置
		if k == 0 {
			lastlng = v.Longitude
			lastlat = v.Latitude
		}
	}
	//添加系统参数
	this.Data["dataArray"] = DataArrayStr
	this.Data["dataArrayPGA"] = DataArrayStrPGA
	this.Data["dataArraySI"] = DataArrayStrSI

	this.Data["dataSize"] = len(dataArray)

	this.Data["lastlng"] = lastlng
	this.Data["lastlat"] = lastlat

	usegis := SystemConfigs.UseGis
	if usegis {
		this.Data["gisServiceUrl"] = SystemConfigs.GisServiceUrl
		this.Data["gisServiceParams"] = SystemConfigs.GisServiceParams
		this.Data["gisBasicLayer"] = SystemConfigs.GisLayerBasic
		this.Data["gisChinaLayer"] = SystemConfigs.GisLayerChina
		this.TplNames = "eventline.html"
	} else {
		this.TplNames = "eventline-nogis.html"
	}
	this.Render()

}

//根据地震事件构造等值线
func (this *EventController) EventLineJson() {

	eventid := this.GetString(":id")
	var eventSignal models.EventSignal = models.EventSignal{}
	//查找当前事件
	event, err0 := dao.GetEventById(eventid)
	if err0 != nil {
		log.Warnf("查找事件[%s]失败:%s", eventid, err0.Error())
	}
	if event.IsConfirm {
		eventSignal, err0 = dao.GetEventSignalById(event.SignalId)
	}

	//查找报警数据
	alarms, err := dao.GetAlarmsByEventId(eventid)
	if err != nil {
		log.Warnf("查找等值线的报警数据时出错:%s", err.Error())
	}

	//是否加入网格化虚拟站点
	dataArray := NetGridCompute(alarms, eventSignal)

	//传递的数据值
	data := make(map[string]interface{})
	DataArrayStr := ""
	DataArrayStrPGA := ""
	DataArrayStrSI := ""
	var lastlng, lastlat float32

	for k, v := range dataArray {
		if k < len(dataArray)-1 {
			DataArrayStr += v.String() + ","
			DataArrayStrPGA += v.StringPGA() + ","
			DataArrayStrSI += v.StringSI() + ","
		} else {
			DataArrayStr += v.String()
			DataArrayStrPGA += v.StringPGA()
			DataArrayStrSI += v.StringSI()
		}
		//设置地图中心点位置
		if k == 0 {
			lastlng = v.Longitude
			lastlat = v.Latitude
		}
	}
	//添加系统参数
	data["dataArray"] = DataArrayStr
	data["dataArrayPGA"] = DataArrayStrPGA
	data["dataArraySI"] = DataArrayStrSI

	data["dataSize"] = len(dataArray)

	data["lastlng"] = lastlng
	data["lastlat"] = lastlat

	this.Data["json"] = data
	this.ServeJson()

}

//地震事件定位
func (this *EventController) QuakeLocation() {
	this.Data["title"] = "地震定位"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()
	sid := this.GetString(":id")
	signal, err := dao.GetEventSignalById(sid)
	if err != nil {
		log.Warnf("获取地震数据失败:%s", err.Error())
	}
	this.Data["signal"] = signal

	usegis := SystemConfigs.UseGis
	if usegis {
		this.Data["gisServiceUrl"] = SystemConfigs.GisServiceUrl
		this.Data["gisServiceParams"] = SystemConfigs.GisServiceParams
		this.Data["gisBasicLayer"] = SystemConfigs.GisLayerBasic
		this.Data["gisChinaLayer"] = SystemConfigs.GisLayerChina
		this.TplNames = "quakelocation.html"
	} else {
		this.TplNames = "quakelocation-nogis.html"
	}

	this.Render()
}

//网格化计算数据点
func NetGridCompute(alarms *[]models.AlarmInfo, eventSignal models.EventSignal) (ng []models.NetGrid) {

	dataArray := make([]models.NetGrid, len(*alarms), len(*alarms)*2)
	for k, v := range *alarms {
		dataArray[k] = models.NetGrid{Longitude: v.Longitude, Latitude: v.Latitude, Value: v.Intensity, PGAValue: v.PGA, SIValue: v.SI}
	}
	//计算虚拟的网格值
	if eventSignal.Id != "" {
		//TODO待添加算法,转移到GIS服务中进行
	}
	return dataArray
}

//添加地震事件
func (this *EventController) AddEventSignal() {
	body := this.Ctx.Request.Body
	defer body.Close()

	content, err := ioutil.ReadAll(body)
	if err != nil {
		log.Errorf("地震事件接口读入数据错误:%s", err.Error())
		this.writeResponse(false, err.Error())
		return
	}
	//xml解析
	earthQuake := new(models.EarthQuake)
	if err = xml.Unmarshal(content, earthQuake); err != nil {
		log.Errorf("地震事件接口xml解析错误:%s", err.Error())
		this.writeResponse(false, err.Error())
		return
	}
	eventSignal := new(models.EventSignal)

	eventSignal.Id = util.GUID()
	eventSignal.Longitude = earthQuake.Longitude
	eventSignal.Latitude = earthQuake.Latitude
	eventSignal.Level = earthQuake.Level
	eventSignal.EventId = earthQuake.EVENT_ID
	eventSignal.CODE = earthQuake.CODE
	eventSignal.CNAME = earthQuake.CNAME
	eventSignal.DEPTH = earthQuake.DEPTH
	eventSignal.LOCATION_CNAME = earthQuake.LOCATION_CNAME
	tm, errt := time.Parse(EarthQuakeTimeLayout, earthQuake.Time)
	if errt != nil {
		eventSignal.Time = time.Now()
	}
	eventSignal.Time = tm
	eventSignal.ReceiveTime = time.Now()

	err = dao.EventSignalSave(eventSignal)
	if err != nil {
		log.Errorf("保存接收的地震事件失败:%s", err.Error())
		this.writeResponse(false, err.Error())
		return
	}

	log.Infof("成功接收了地震事件%s [%f,%f] %d级", earthQuake.Time, earthQuake.Longitude, earthQuake.Latitude, earthQuake.Level)
	this.writeResponse(true, "success")

	//提供回送数据
	//go FeedbackData(eventSignal)
	return
}

//写入回复数据
func (this *EventController) writeResponse(ok bool, msg string) {
	reswriter := this.Ctx.ResponseWriter
	fb := models.Feedback{}
	fb.Ok = ok
	fb.Message = msg
	cont, err := xml.Marshal(fb)
	if err == nil {
		reswriter.Write(cont)
	}
}

//回送数据，调用川局服务器接口
func (this *EventController) FeedbackData(eventSignal *models.EventSignal) {

	alarmList := models.AlarmDataList{}
	body, err := xml.Marshal(alarmList)
	if err != nil {
		return
	}

	client := &http.Client{}
	reqest, _ := http.NewRequest("POST", "http://218.6.242.153/Service.asmx", bytes.NewBuffer(body))

	reqest.Header.Set("Accept", "application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")

	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		log.Info(bodystr)
	}
}
