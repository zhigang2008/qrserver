package quickserver

import (
	//"errors"
	//"fmt"
	log "github.com/cihub/seelog"
	//"net"
	"time"
)

func DelayGenerateReport(event *Event) {
	//延时处理.
	time.Sleep(3 * time.Minute)
}

func GenerateReport(event *Event) {
	summary := generateReportSummary(event)
	generateReportMap(event)
}

func generateReportSummary(event *Event) (sumary map[string]interface{}) {

	//基本信息
	sumary["时间"] = event.EventTimeStr
	sumary["报警数"] = event.AlarmCount

	alarms, err := server.dataManager.GetAlarmsByEvent(event)
	if err == nil {
		sumary[""]
		return
	}
	return
}

func generateReportMap(event *Event) {

}
