package controllers

import (
	"dqs/dao"
	"dqs/models"
	//"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
)

type ExchangeController struct {
	BaseController
}

//实时报警信息列表
func (this *ExchangeController) ExportAlarms() {

	begintime := this.GetString("btime")
	endtime := this.GetString("etime")
	sid := this.GetString("sid")

	returnList := models.AlarmList{}
	returnList.BeginTime = begintime
	returnList.EndTime = endtime
	returnList.SensorId = sid

	//执行查询
	alarmlist, err2 := dao.ExportAlarms(sid, begintime, endtime)
	if err2 != nil {
		log.Warnf("导出实时报警信息失败:%s", err2.Error())
	} else {
		returnList.Alarms = alarmlist
	}
	this.Data["xml"] = returnList
	this.ServeXml()

}
