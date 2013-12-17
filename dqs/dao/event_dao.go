package dao

import (
	"dqs/models"
	"dqs/util"
	"labix.org/v2/mgo/bson"
	"time"
)

const (
	EventTimeLayout = "2006-01-02"
)

//保存确认信号
func EventSignalAdd(signal *models.EventSignal) (err error) {
	c := GetSession().DB(DatabaseName).C(EventSignalCollection)

	err = c.Insert(signal)
	if err != nil {
		return err
	}
	return nil
}

//事件列表
func EventSignalList(n int) (*[]models.EventSignal, error) {
	c := GetSession().DB(DatabaseName).C(EventSignalCollection)

	eventSignals := []models.EventSignal{}
	//先查找
	err := c.Find(bson.M{}).Sort("-time").Limit(n).All(&eventSignals)
	if err != nil {
		return nil, err
	}
	return &eventSignals, nil
}

//获取事件确认信号
func GetEventSignalById(sid string) (signal models.EventSignal, err error) {
	c := GetSession().DB(DatabaseName).C(EventSignalCollection)

	err0 := c.Find(bson.M{"id": sid}).One(&signal)
	if err0 != nil {
		return models.EventSignal{}, err0
	}
	return signal, nil
}

//获取当前时间段内有效的信号
func GetValidEventSignal(begintime, endtime time.Time) (signal models.EventSignal, err error) {
	c := GetSession().DB(DatabaseName).C(EventSignalCollection)

	m := bson.M{}
	timeparam := bson.M{}
	timeparam["$gte"] = begintime
	timeparam["$lt"] = endtime

	m["time"] = timeparam

	err0 := c.Find(&m).Sort("-time").One(&signal)
	if err0 != nil {
		return models.EventSignal{}, err0
	}
	return signal, nil
}

//保存事件
func EventAdd(event *models.Event) (err error) {
	c := GetSession().DB(DatabaseName).C(EventCollection)
	err = c.Insert(event)
	if err != nil {
		return err
	}
	return nil
}

//保存更新事件
func EventUpsert(event *models.Event) (err error) {
	c := GetSession().DB(DatabaseName).C(EventCollection)
	query := bson.M{"eventid": event.EventId}
	_, err = c.Upsert(query, event)
	if err != nil {
		return err
	}
	return nil
}

//获取事件
func GetEventById(sid string) (event models.Event, err error) {
	c := GetSession().DB(DatabaseName).C(EventCollection)

	err0 := c.Find(bson.M{"eventid": sid}).One(&event)
	if err0 != nil {
		return models.Event{}, err0
	}
	return event, nil
}

//更新事件
func EventUpdate(event *models.Event) (err error) {
	c := GetSession().DB(DatabaseName).C(EventCollection)
	err = c.Update(bson.M{"eventid": event.EventId}, event)
	if err != nil {
		return err
	}
	return nil
}

//事件列表
func EventList(n int) (*[]models.Event, error) {
	c := GetSession().DB(DatabaseName).C(EventCollection)

	events := []models.Event{}

	err := c.Find(bson.M{}).Sort("-eventtime").Limit(n).All(&events)
	if err != nil {
		return nil, err
	}
	return &events, nil
}

//获取最近的一个事件
func GetLastEvent() (event models.Event, err error) {
	c := GetSession().DB(DatabaseName).C(EventCollection)

	err0 := c.Find(bson.M{}).Sort("-eventtime").One(&event)
	if err0 != nil {
		return models.Event{}, err0
	}
	return event, nil
}

//最后一个事件的报警数据
func GetAlarmsByEvent(event *models.Event) (*[]models.AlarmInfo, error) {
	c := GetSession().DB(DatabaseName).C(EventCollection)
	m := bson.M{"eventid": event.EventId}
	alist := []models.AlarmInfo{}
	err0 := c.Find(&m).All(&alist)
	if err0 != nil {
		return nil, err0
	}
	return &alist, nil
}

//根据事件ID查找报警数据
func GetAlarmsByEventId(eventid string) (*[]models.AlarmInfo, error) {
	c := GetSession().DB(DatabaseName).C(DataCollection)
	m := bson.M{"eventid": eventid}
	alist := []models.AlarmInfo{}
	err0 := c.Find(&m).All(&alist)
	if err0 != nil {
		return nil, err0
	}
	return &alist, nil
}

//分页查询
func EventPageList(p *util.Pagination) error {
	c := GetSession().DB(DatabaseName).C(EventCollection)
	events := []models.Event{}

	//构造查询参数
	m := bson.M{}
	eventid := p.QueryParams["eventid"]
	begintime := p.QueryParams["begintime"]
	endtime := p.QueryParams["endtime"]

	if eventid != nil {
		m["eventid"] = eventid
	}

	timeparam := bson.M{}
	hasTime := false
	if begintime != nil {
		sbtime, ok := begintime.(string)
		if ok {
			btime, _ := time.ParseInLocation(EventTimeLayout, sbtime, Local)
			timeparam["$gte"] = btime
			hasTime = true
		}
	}
	if endtime != nil {
		setime, ok := endtime.(string)
		if ok {
			etime, _ := time.ParseInLocation(EventTimeLayout, setime, Local)
			etime = etime.Add(time.Hour * 24)
			timeparam["$lt"] = etime
			hasTime = true
		}
	}
	if hasTime {
		m["eventtime"] = timeparam
	}

	//查询总数
	query := c.Find(&m).Sort("-eventtime")
	count, err := query.Count()
	if err != nil {
		return err
	}
	p.Count = count

	//查找列表
	err = query.Skip(p.SkipNum()).Limit(p.PageSize).All(&events)
	if err != nil {
		return err
	}
	p.Data = events
	return nil
}

//分页查询
func EventSignalPageList(p *util.Pagination) error {
	c := GetSession().DB(DatabaseName).C(EventSignalCollection)
	signals := []models.EventSignal{}

	//构造查询参数
	m := bson.M{}
	signalid := p.QueryParams["signalid"]
	level := p.QueryParams["level"]
	begintime := p.QueryParams["begintime"]
	endtime := p.QueryParams["endtime"]

	if signalid != nil {
		m["id"] = signalid
	}
	if level != nil {
		m["level"] = level
	}

	timeparam := bson.M{}
	hasTime := false
	if begintime != nil {
		sbtime, ok := begintime.(string)
		if ok {
			btime, _ := time.ParseInLocation(EventTimeLayout, sbtime, Local)
			timeparam["$gte"] = btime
			hasTime = true
		}
	}
	if endtime != nil {
		setime, ok := endtime.(string)
		if ok {
			etime, _ := time.ParseInLocation(EventTimeLayout, setime, Local)
			etime = etime.Add(time.Hour * 24)
			timeparam["$lt"] = etime
			hasTime = true
		}
	}
	if hasTime {
		m["time"] = timeparam
	}

	//查询总数
	query := c.Find(&m).Sort("-time")
	count, err := query.Count()
	if err != nil {
		return err
	}
	p.Count = count

	//查找列表
	err = query.Skip(p.SkipNum()).Limit(p.PageSize).All(&signals)
	if err != nil {
		return err
	}
	p.Data = signals
	return nil
}
