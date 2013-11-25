package controllers

import (
	"dqs/dao"
	"dqs/util"
	//"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
)

type AlarmController struct {
	BaseController
}

//报警信息列表
func (this *AlarmController) Get() {
	this.Data["title"] = "报警信息"
	this.Data["author"] = "wangzhigang"
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

	sid := this.GetString("sensorid")
	if sid != "" {
		pagination.AddParams("sensorid", sid)
	}

	//执行查询
	err = dao.AlarmList(&pagination)
	if err != nil {
		log.Warnf("查询报警信息失败:%s", err.Error())
	}
	pagination.Compute()

	this.Data["pagedata"] = pagination
	this.TplNames = "alarmlist.html"
	this.Render()

}

//实时报警信息列表
func (this *AlarmController) GetRealtimeAlarm() {

	timestep, err := this.GetInt("time")
	if err != nil {
		timestep = 5
	}
	//执行查询
	alarmlist, err2 := dao.GetRealtimeAlarm(timestep)
	if err2 != nil {
		log.Warnf("查询实时报警信息失败:%s", err2.Error())
	}

	this.Data["json"] = alarmlist
	this.ServeJson()

}

//获取波形图数据
func (this *AlarmController) ShowWaveInfoById() {

	oid := this.GetString(":objectid")

	//执行查询
	waveInfo, err2 := dao.GetWaveInfoById(oid)
	if err2 != nil {
		log.Warnf("查询波形数据信息失败:%s", err2.Error())
	}

	this.Data["waveData"] = waveInfo
	this.TplNames = "wave.html"
	this.Render()
}

//获取波形图数据
func (this *AlarmController) ShowWaveInfo() {

	sid := this.GetString(":sid")
	seqno := this.GetString(":seqno")

	//执行查询
	waveInfo, err2 := dao.GetWaveInfo(sid, seqno)
	if err2 != nil {
		log.Warnf("查询波形数据信息失败:%s", err2.Error())
	}

	this.Data["waveData"] = waveInfo
	this.TplNames = "wave.html"
	this.Render()
}
