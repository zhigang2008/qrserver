package quickserver

import (
	//"errors"
	//"fmt"
	"dqs/util"
	log "github.com/cihub/seelog"
	//"net"
	"time"
)

//震情分析器
type EarthquakeAnalyzer struct {
	name string
	dm   *DataManager
}

//构造震情分析器
func NewEarthquakeAnalyzer(name string, dm *DataManager) *EarthquakeAnalyzer {
	analyzer := new(EarthquakeAnalyzer)
	analyzer.name = name
	analyzer.dm = dm
	return analyzer
}

//震情分析过程
func (eq *EarthquakeAnalyzer) analyze(a *AlarmInfo) {
	log.Infof("进入震情分析过程[%s-%s]", a.SensorId, a.SeqNo)
	//判断是否已经有确认信号进入
	if es, ok := eq.hasEventSignal(); ok == true {
		log.Infof("已经有震情确认信号[%f,%f]-%d", es.Longitude, es.Latitude, es.Level)
		event, exist := eq.fetchEventBySignal(&es)
		if exist {
			a.EventId = event.EventId
			//更新报警信息
			eq.updateAlarmEvent(a)
			//更新事件信息
			event.AlarmCount++
			eq.updateEvent(&event)
		} else {
			//创建Event
			event = Event{}
			event.EventId = util.GUID()
			event.AlarmCount = 1
			event.EventTime = time.Now()
			event.IsConfirm = true
			event.SignalId = es.Id
			event.Signal = es
			//添加事件
			eq.saveEvent(&event)
			//更新报警事件
			a.EventId = event.EventId
			eq.updateAlarmEvent(a)
		}
	} else {
		//没有信号时的处理方法
		eq.eventJudge(a)
	}

}

//判断是否有震情确认信号
func (eq *EarthquakeAnalyzer) hasEventSignal() (event EventSignal, ok bool) {
	//通过数据库数据判定
	return EventSignal{}, false
}

//根据自主算法进行事件判断
func (eq *EarthquakeAnalyzer) eventJudge(a *AlarmInfo) {

}

//确定震情事件
func (eq *EarthquakeAnalyzer) eventRecord(a *AlarmInfo) {

}

//获取震情事件
func (eq *EarthquakeAnalyzer) fetchEventBySignal(es *EventSignal) (Event, bool) {

	return Event{}, false
}

//更新报警事件
func (eq *EarthquakeAnalyzer) updateAlarmEvent(a *AlarmInfo) {

}

//添加震情事件
func (eq *EarthquakeAnalyzer) saveEvent(event *Event) {

}

//更新震情事件
func (eq *EarthquakeAnalyzer) updateEvent(event *Event) {

}
