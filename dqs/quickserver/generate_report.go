package quickserver

import (
	//"errors"
	"fmt"
	log "github.com/cihub/seelog"
	//"net"
	"dqs/util"
	"time"
)

//延时进行速报构建
func DelayGenerateReport(event *Event) {
	//延时处理.
	delaytime := GlobalConfig.ReportCfg.SleepTime
	time.Sleep(time.Duration(delaytime) * time.Minute)

	newEvent, err := server.dataManager.GetEventById(event.EventId)
	if err != nil {
		log.Warnf("速报处理-查询最新的时间信息失败:%s", err.Error())
		newEvent = *event
	}

	//log.Infof("级别判断%d-%d", event.MaxLevel, GlobalConfig.ReportCfg.ReportLevel)
	if event.MaxLevel >= GlobalConfig.ReportCfg.ReportLevel {
		//制作速报信息
		GenerateReport(&newEvent)

	}

}

//生成速报
func GenerateReport(event *Event) {
	log.Infof("准备生成速报...")

	summary := generateReportSummary(event)
	imgfile := generateReportMap(event)

	report := new(Report)
	report.ReportId = util.GUID()
	report.GenerateTime = time.Now()
	report.EventId = event.EventId
	report.Event = *event
	report.Summary = summary
	report.ImageFile = imgfile
	report.Verify = false
	report.Sended = false
	report.Valid = true

	//保存速报
	err := server.dataManager.ReportSave(report)
	if err != nil {
		log.Warnf("速报保存失败:%s", err.Error())
	} else {
		log.Info("速报已生成.")
	}
}

//生成概要信息
func generateReportSummary(event *Event) ReportSummary {

	sumary := ReportSummary{}
	//基本信息
	sumary.EventTime = event.EventTimeStr
	sumary.AlarmCount = event.AlarmCount

	cs := make(map[int]int)
	alarms, err := server.dataManager.GetAlarmsByEvent(event)
	if err == nil {
		for _, v := range *alarms {
			i := v.Intensity
			val, ok := cs[i]
			if ok {
				cs[i] = val + 1
			} else {
				cs[i] = 1
			}

		}
	} else {
		log.Infof("获取震情事件的所有报警信息失败:%s", err.Error())
	}

	alarmSumary := ""
	for k, v := range cs {
		alarmSumary += fmt.Sprintf("%d级-%d; ", k, v)

	}
	sumary.Brief = alarmSumary

	//实际地震信息
	if event.IsConfirm {
		signal := event.Signal
		stime := signal.Time.Format(CommonTimeLayout)
		sumary.QuakeInfo = fmt.Sprintf("%s 发生于[%f,%f] 震级%d级", stime, signal.Longitude, signal.Latitude, signal.Level)
	}
	return sumary
}

//生成速报的图片信息
func generateReportMap(event *Event) string {
	return ""
}
