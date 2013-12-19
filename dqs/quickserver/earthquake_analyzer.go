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
			//更新最高级别
			if a.Intensity > event.MaxLevel {
				event.MaxLevel = a.Intensity
			}
			eq.updateEvent(&event)
		} else {
			//创建Event
			event = Event{}
			event.EventId = util.GUID()
			event.AlarmCount = 1
			event.MaxLevel = a.Intensity
			event.MaxLevel = a.Intensity
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
	timespan := GlobalConfig.EventParams.SignalTimeSpan

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

	} else {
		//存在,则判断时间及其它因素.
		if eq.isNewEvent(*a, &lastEvent) {
			eq.eventRecord(a)
		} else {
			//更新当前报警事件
			a.EventId = lastEvent.EventId
			eq.updateAlarmEvent(a)
			//更新事件数量
			lastEvent.AlarmCount += 1
			eq.updateEvent(&lastEvent)
		}
	}

}

//确定震情事件
func (eq *EarthquakeAnalyzer) eventRecord(a *AlarmInfo) {

	//先判断该次报警的级别,达到规定级别才记录事件
	if a.Intensity >= GlobalConfig.EventParams.MinEventRecordLevel {

		newEvent := new(Event)
		newEvent.EventId = util.GUID()
		newEvent.AlarmCount = 1
		newEvent.IsConfirm = false
		newEvent.MaxLevel = a.Intensity

		if a.InitRealTime.IsZero() {
			newEvent.EventTime = time.Now()
		} else {
			newEvent.EventTime = a.InitRealTime
		}
		newEvent.EventTimeStr = newEvent.EventTime.Format(CommonTimeLayout)

		//创建观测事件
		eq.saveEvent(newEvent)

		a.EventId = newEvent.EventId

		err := eq.dm.updateAlarmEvent(a)
		if err != nil {
			log.Warnf("更新震情报警事件[%s-%s]失败:%s", a.SensorId, a.SeqNo, err.Error())
		}
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
	log.Infof("产生新的震情事件[%s]", event.EventTimeStr)
	//-------速报制作过程---------------
	go DelayGenerateReport(event)

}

//更新震情事件
func (eq *EarthquakeAnalyzer) updateEvent(event *Event) {
	err := eq.dm.EventUpdate(event)
	if err != nil {
		log.Warnf("更新地震事件[%s]出错:%s", event.EventId, err.Error())
	}

}

//是否是新事件
func (eq *EarthquakeAnalyzer) isNewEvent(a AlarmInfo, le *Event) bool {
	lastEventTime := le.EventTime
	alarmTime := a.InitRealTime
	var laterAlarmTime, earlyAlarmTime time.Time
	var eventGap, laterAlarmGap, earlyAlarmGap, averageGap time.Duration

	//如果时间距离够长,则判断为新事件
	eventGap = alarmTime.Sub(lastEventTime)
	if eventGap > time.Minute*time.Duration(GlobalConfig.EventParams.NewEventTimeGap) {
		return true
	}

	//根据事件报警信息进行判断
	lastAlarms, err0 := eq.dm.GetAlarmsByEvent(le)
	if err0 != nil {
		log.Warnf("查询最末事件[%s]的报警信息出错时出错,将该报警标记为新震情事件:%s", le.EventId, err0.Error())
		return true
	}
	//在此进行额外的数据校正
	le.AlarmCount = len(*lastAlarms)
	//eq.updateEvent(&le)

	//当报警数过少时,则无法进行群体分析,则断定为同一事件
	if len(*lastAlarms) < GlobalConfig.EventParams.ValidEventAlarmCount {
		log.Infof("最近的震情事件[%s]报警数量仅有[%d]个,该次报警判断属于同一事件.", le.EventId, len(*lastAlarms))
		return false
	}
	//较多数数量时,则进行数据距离差的分析
	laterAlarmTime = (*lastAlarms)[0].InitRealTime
	earlyAlarmTime = (*lastAlarms)[0].InitRealTime
	sumGap := time.Duration(0)
	for _, v := range *lastAlarms {
		if v.InitRealTime.After(laterAlarmTime) {
			laterAlarmTime = v.InitRealTime
		}
		if v.InitRealTime.Before(earlyAlarmTime) {
			earlyAlarmTime = v.InitRealTime
		}
		sumGap += v.InitRealTime.Sub(lastEventTime)
	}

	averageGap = sumGap / time.Duration(len(*lastAlarms))
	laterAlarmGap = laterAlarmTime.Sub(lastEventTime)
	earlyAlarmGap = earlyAlarmTime.Sub(lastEventTime)

	initMultiple := float64((laterAlarmGap + earlyAlarmGap)) / 2 / float64(averageGap)
	alarmMultiple := float64(eventGap) / float64(averageGap)

	//离散度过大,则判断为新事件
	if alarmMultiple > initMultiple*GlobalConfig.EventParams.NewEventGapMultiple {
		return true
	}

	return false
}
