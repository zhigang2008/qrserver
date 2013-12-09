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
func (eq *EarthquakeAnalyzer) hasEventSignal() (signal EventSignal, ok bool) {

	//通过数据库数据判定
	timespan := ServerConfigs.EventParams.SignalTimeSpan

	now := time.Now()
	btime := now.Add(time.Minute * -time.Duration(timespan))
	etime := now.Add(time.Minute * time.Duration(timespan))
	vsignal, err := eq.dm.GetValidEventSignal(btime, etime)
	if err != nil {
		return EventSignal{}, false
	}
	if vsignal.Id == "" {
		return EventSignal{}, false
	}
	return vsignal, true
}

//根据自主算法进行事件判断
func (eq *EarthquakeAnalyzer) eventJudge(a *AlarmInfo) {
	lastEvent, err := eq.dm.GetLastEvent()
	//如果不存在,则创建新的震情事件
	if err != nil || lastEvent.EventId == "" {
		eq.eventRecord(a)

	}
	//存在,则判断时间及其它因素.

}

//确定震情事件
func (eq *EarthquakeAnalyzer) eventRecord(a *AlarmInfo) {
	newEvent := new(Event)
	newEvent.EventId = util.GUID()
	newEvent.AlarmCount = 1
	newEvent.IsConfirm = false
	if a.InitRealTime.IsZero() {
		newEvent.EventTime = time.Now()
	} else {
		newEvent.EventTime = a.InitRealTime
	}

	err := eq.dm.EventAdd(newEvent)
	if err != nil {
		log.Warnf("创建新的震情事件记录失败:%s", err.Error())
	}
	a.EventId = newEvent.EventId

	err = eq.dm.updateAlarmEvent(a)
	if err != nil {
		log.Warnf("更新震情报警事件[%s-%s]失败:%s", a.SensorId, a.SeqNo, err.Error())
	}

}

//获取震情事件
func (eq *EarthquakeAnalyzer) fetchEventBySignal(es *EventSignal) (Event, bool) {

	e, err := eq.dm.GetEventBySingalId(es.Id)
	if err != nil {
		log.Warnf("根据信号[%s]获取震情事件出错:%s", es.Id, err.Error())
		return Event{}, false
	}
	if e.EventId == "" {
		return Event{}, false
	}
	return e, true
}

//更新报警事件
func (eq *EarthquakeAnalyzer) updateAlarmEvent(a *AlarmInfo) {
	err := eq.dm.updateAlarmEvent(a)
	if err != nil {
		log.Warnf("更新报警信息[%s-%s]事件号[%s]出错:%s", a.SensorId, a.SeqNo, a.EventId, err.Error())
	}
}

//添加震情事件
func (eq *EarthquakeAnalyzer) saveEvent(event *Event) {
	err := eq.dm.EventAdd(event)
	if err != nil {
		log.Warnf("添加地震事件[%s]出错:%s", event.EventId, err.Error())
	}
}

//更新震情事件
func (eq *EarthquakeAnalyzer) updateEvent(event *Event) {
	err := eq.dm.EventUpdate(event)
	if err != nil {
		log.Warnf("更新地震事件[%s]出错:%s", event.EventId, err.Error())
	}

}
