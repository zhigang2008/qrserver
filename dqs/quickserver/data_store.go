package quickserver

import (
	//	"fmt"
	"errors"
	log "github.com/cihub/seelog"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"sort"
	"strconv"
	"sync"
	"time"
)

const (
	defaultDatabase              = "dqs"            //默认数据库名称
	defaultDataCollection        = "data"           //默认数据Collection
	defaultDeviceCollection      = "device"         //默认设备Collection
	defaultWaveCollection        = "wavedata"       //默认波形记录Collection
	defaultEventCollection       = "event"          //默认事件Collection
	defaultEventSignalCollection = "eventsignal"    //默认事件信号Collection
	defaultIntensityCollection   = "intensitymap"   //默认事件信号Collection
	defaultConfigCollection      = "runtimeConfigs" //默认配置信息表
	defaultReportCollection      = "reports"        //默认配置信息表
)

var (
	mux         sync.Mutex
	ErrNotFound = errors.New("not found")
)

//数据库连接服务
type DataManager struct {
	session                    *mgo.Session
	databaseName               string
	dataCollection             string
	deviceCollection           string
	waveCollection             string
	eventCollection            string
	eventSignalCollection      string
	intensityMappingCollection string
	configCollection           string
	reportCollection           string
}

//初始化数据库连接
func InitDatabase(conf DataServerConfig) (dm *DataManager, err error) {
	mux.Lock()
	defer mux.Unlock()

	session1, err := mgo.Dial(conf.Host + ":" + strconv.Itoa(conf.Port))
	if err != nil {
		log.Criticalf("不能创建数据库连接[%s:%d]:%s", conf.Host, conf.Port, err.Error())
		return nil, err
	}
	session1.SetMode(mgo.Monotonic, true)
	log.Info("创建了数据连接")
	dataManager := &DataManager{
		session:                    session1,
		databaseName:               defaultDatabase,
		dataCollection:             defaultDataCollection,
		deviceCollection:           defaultDeviceCollection,
		waveCollection:             defaultWaveCollection,
		eventCollection:            defaultEventCollection,
		eventSignalCollection:      defaultEventSignalCollection,
		intensityMappingCollection: defaultIntensityCollection,
		configCollection:           defaultConfigCollection,
		reportCollection:           defaultReportCollection,
	}
	//设置配置文件指定值
	if conf.DataBaseName != "" {
		dataManager.databaseName = conf.DataBaseName
	}
	if conf.DataCollection != "" {
		dataManager.dataCollection = conf.DataCollection
	}
	if conf.DeviceCollection != "" {
		dataManager.deviceCollection = conf.DeviceCollection
	}

	//log.Infof("DataManager初始化完成.")
	return dataManager, nil
}

//获取session
func (dm *DataManager) getSession() *mgo.Session {
	return dm.session
}

//保存报警信息
func (dm *DataManager) FlashDataSave(data *AlarmInfo) (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.dataCollection)
	err = c.Insert(data)
	if err != nil {

		return err
	}
	return nil
}

//保存或更改报警信息
func (dm *DataManager) AlarmUpsert(data *AlarmInfo) (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.dataCollection)
	colQuerier := bson.M{"sensorid": data.SensorId, "seqno": data.SeqNo}

	_, err = c.Upsert(colQuerier, data)
	if err != nil {
		return err
	}
	return nil
}

//保存或更改报警信息
func (dm *DataManager) GetAlarmByIdAndSeqno(sid, seqno string) (a AlarmInfo, err error) {
	c := dm.session.DB(dm.databaseName).C(dm.dataCollection)
	colQuerier := bson.M{"sensorid": sid, "seqno": seqno}
	alarm := AlarmInfo{}
	err0 := c.Find(colQuerier).One(&alarm)
	if err0 != nil {
		if err0 == mgo.ErrNotFound {
			return alarm, ErrNotFound
		}
		return alarm, err0
	}
	return alarm, nil
}

//更新报警信息的事件号
func (dm *DataManager) updateAlarmEvent(data *AlarmInfo) (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.dataCollection)
	colQuerier := bson.M{"sensorid": data.SensorId, "seqno": data.SeqNo}

	err = c.Update(colQuerier, bson.M{"$set": bson.M{"eventid": data.EventId}})
	if err != nil {
		return err
	}
	return nil
}

//更新设备状态及参数
func (dm *DataManager) UpdateDeviceStatus(status *SensorInfo) (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.deviceCollection)

	device := DeviceInfo{}
	//先查找设备
	err = c.Find(bson.M{"sensorid": status.SensorId}).One(&device)
	if err != nil {
		return err
	}
	device.SetParams = *status
	device.UpdateTime = time.Now()

	//更新数据
	//err0 := c.Update(&bson.M{"sensorid": device.SensorId}, &device)
	//只更新部分信息
	colQuerier := bson.M{"sensorid": device.SensorId}
	change := bson.M{"$set": bson.M{"setparams": device.SetParams, "updatetime": time.Now()}}
	err0 := c.Update(colQuerier, change)
	if err0 != nil {
		log.Infof("数据库更新设备参数失败:%s", err0.Error())
		return err0
	}

	return nil
}

//设备注册
func (dm *DataManager) DeviceRegister(device *DeviceInfo) (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.deviceCollection)

	/*
		changeInfo, err0 := c.Upsert(&bson.M{"sensorid": device.SensorId}, &device)
		if err0 != nil {
			log.Infof("数据库更新设备注册信息失败:%d", changeInfo.Updated)
			return err0
		}
		return nil
	*/

	devicetemp := DeviceInfo{}
	//先查找设备
	err = c.Find(bson.M{"sensorid": device.SensorId}).One(&devicetemp)
	if err != nil {
		if err == mgo.ErrNotFound { //不存在,插入设备信息
			err0 := c.Insert(device)
			if err0 != nil {
				return err0
			}
			return nil

		} else {
			return err
		}
	}

	//更新设备信息
	colQuerier := bson.M{"sensorid": device.SensorId}
	change := bson.M{"$set": bson.M{"registertime": time.Now(), "online": true, "updatetime": time.Now(), "remoteaddr": device.RemoteAddr}}
	err0 := c.Update(colQuerier, change)
	if err0 != nil {
		log.Infof("数据库更新设备参数失败:%s", err0.Error())
		return err0
	}

	return nil
}

//设备下线
func (dm *DataManager) DeviceOffline(deviceid string) (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.deviceCollection)
	//更新设备信息
	colQuerier := bson.M{"sensorid": deviceid}
	change := bson.M{"$set": bson.M{"offtime": time.Now(), "online": false, "updatetime": time.Now()}}
	err0 := c.Update(colQuerier, change)
	if err0 != nil {
		return err0
	}
	return nil
}

//所有设备下线
func (dm *DataManager) ResetAllDeviceStatus() (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.deviceCollection)
	//更新设备信息
	colQuerier := bson.M{}
	change := bson.M{"$set": bson.M{"online": false}}
	_, err0 := c.UpdateAll(colQuerier, change)
	if err0 != nil {
		return err0
	}
	return nil
}

//查找设备信息
func (dm *DataManager) DeviceList(n int) (*[]DeviceInfo, error) {
	c := dm.session.DB(dm.databaseName).C(dm.deviceCollection)

	devices := []DeviceInfo{}
	//先查找设备
	err := c.Find(bson.M{}).Sort("-registertime").Limit(n).All(&devices)
	if err != nil {
		return nil, err
	}
	return &devices, nil
}

//查找报警信息
func (dm *DataManager) AlarmList(n int) (*[]AlarmInfo, error) {
	c := dm.session.DB(dm.databaseName).C(dm.dataCollection)

	alarms := []AlarmInfo{}

	err := c.Find(bson.M{}).Sort("-createtime").Limit(n).All(&alarms)
	if err != nil {
		return nil, err
	}
	return &alarms, nil
}

//添加波形记录信息
func (dm *DataManager) WaveDataAdd(data *WaveInfo) (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.waveCollection)
	//添加objectid
	data.Id = bson.NewObjectId()
	data.LUD = time.Now()

	err = c.Insert(data)
	if err != nil {
		return err
	}
	return nil
}

//更新波形记录信息
func (dm *DataManager) WaveDataUpdate(data *WaveInfo) (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.waveCollection)

	//colQuerier := bson.M{"sensorid", data.SensorId}
	//change := bson.M{"$set": data}
	err0 := c.Update(bson.M{"_id": data.Id}, data)
	if err0 != nil {
		return err0
	}
	return nil
}

//获取Wavedata
func (dm *DataManager) GetWaveData(sid, seqno string) (wave WaveInfo, err error) {
	c := dm.session.DB(dm.databaseName).C(dm.waveCollection)

	//wave := WaveInfo{}
	m := bson.M{}
	m["sensorid"] = sid
	m["seqno"] = seqno

	err0 := c.Find(m).Sort("-lud").One(&wave)
	if err0 != nil {
		return WaveInfo{}, err0
	}
	return wave, nil
}

//获取最新的Wavedata
func (dm *DataManager) GetLastWave(sid string) (wave WaveInfo, err error) {
	c := dm.session.DB(dm.databaseName).C(dm.waveCollection)

	//wave := WaveInfo{}
	err0 := c.Find(bson.M{"sensorid": sid}).Sort("-lud").One(&wave)
	if err0 != nil {
		return WaveInfo{}, err0
	}
	return wave, nil
}

//关闭数据库连接
func (dm *DataManager) DataClose() {
	dm.session.Close()

}

//-------------事件处理-----
//保存确认信号
func (dm *DataManager) EventSignalAdd(signal *EventSignal) (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.eventSignalCollection)

	err = c.Insert(signal)
	if err != nil {
		return err
	}
	return nil
}

//事件列表
func (dm *DataManager) EventSignalList(n int) (*[]EventSignal, error) {
	c := dm.session.DB(dm.databaseName).C(dm.eventSignalCollection)

	eventSignals := []EventSignal{}
	//先查找设备
	err := c.Find(bson.M{}).Sort("-eventtime").Limit(n).All(&eventSignals)
	if err != nil {
		return nil, err
	}
	return &eventSignals, nil
}

//获取事件确认信号
func (dm *DataManager) GetEventSignalById(sid string) (signal EventSignal, err error) {
	c := dm.session.DB(dm.databaseName).C(dm.eventSignalCollection)

	err0 := c.Find(bson.M{"id": sid}).One(&signal)
	if err0 != nil {
		return EventSignal{}, err0
	}
	return signal, nil
}

//获取当前时间段内有效的信号
func (dm *DataManager) GetValidEventSignal(begintime, endtime time.Time) (signal EventSignal, err error) {
	c := dm.session.DB(dm.databaseName).C(dm.eventSignalCollection)

	m := bson.M{}
	timeparam := bson.M{}
	timeparam["$gte"] = begintime
	timeparam["$lt"] = endtime

	m["time"] = timeparam

	err0 := c.Find(&m).Sort("-time").One(&signal)
	if err0 != nil {
		return EventSignal{}, err0
	}
	return signal, nil
}

//保存事件
func (dm *DataManager) EventAdd(event *Event) (err error) {
	event.CreateTime = time.Now()
	c := dm.session.DB(dm.databaseName).C(dm.eventCollection)
	err = c.Insert(event)
	if err != nil {
		return err
	}
	return nil
}

//保存更新事件
func (dm *DataManager) EventUpsert(event *Event) (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.eventCollection)
	query := bson.M{"eventid": event.EventId}
	_, err = c.Upsert(query, event)
	if err != nil {
		return err
	}
	return nil
}

//获取事件
func (dm *DataManager) GetEventById(sid string) (event Event, err error) {
	c := dm.session.DB(dm.databaseName).C(dm.eventCollection)

	err0 := c.Find(bson.M{"eventid": sid}).One(&event)
	if err0 != nil {
		return Event{}, err0
	}
	return event, nil
}

//获取事件
func (dm *DataManager) GetEventBySingalId(sid string) (event Event, err error) {
	c := dm.session.DB(dm.databaseName).C(dm.eventCollection)

	err0 := c.Find(bson.M{"signalid": sid}).One(&event)
	if err0 != nil {
		return Event{}, err0
	}
	return event, nil
}

//更新事件
func (dm *DataManager) EventUpdate(event *Event) (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.eventCollection)
	err = c.Update(bson.M{"eventid": event.EventId}, event)
	if err != nil {
		return err
	}
	return nil
}

//事件列表
func (dm *DataManager) EventList(n int) (*[]Event, error) {
	c := dm.session.DB(dm.databaseName).C(dm.eventCollection)

	events := []Event{}

	err := c.Find(bson.M{}).Sort("-eventtime").Limit(n).All(&events)
	if err != nil {
		return nil, err
	}
	return &events, nil
}

//获取最近的一个事件
func (dm *DataManager) GetLastEvent() (event Event, err error) {
	c := dm.session.DB(dm.databaseName).C(dm.eventCollection)

	err0 := c.Find(bson.M{}).Sort("-eventtime").One(&event)
	if err0 != nil {
		return Event{}, err0
	}
	return event, nil
}

//最后一个事件的报警数据
func (dm *DataManager) GetAlarmsByEvent(event *Event) (*[]AlarmInfo, error) {
	c := dm.session.DB(dm.databaseName).C(dm.dataCollection)
	m := bson.M{"eventid": event.EventId}
	alist := []AlarmInfo{}
	err0 := c.Find(&m).All(&alist)
	if err0 != nil {
		return nil, err0
	}
	return &alist, nil
}

//获取数据库中的烈度对照表
func (dm *DataManager) GetDataMapping() (DataMapping, error) {
	c := dm.session.DB(dm.databaseName).C(dm.intensityMappingCollection)
	datamap := DataMapping{}
	err0 := c.Find(&bson.M{}).One(&datamap)
	if err0 != nil {
		if err0 == mgo.ErrNotFound {
			return datamap, ErrNotFound
		}
		return datamap, err0
	}
	//排序
	sortedPGAMap := datamap.PGAMap
	sort.Sort(sortedPGAMap)
	datamap.PGAMap = sortedPGAMap
	sortedSIMap := datamap.SIMap
	sort.Sort(sortedSIMap)
	datamap.SIMap = sortedSIMap

	return datamap, nil
}

//添加烈度对照表
func (dm *DataManager) CreateDataMapping(cfg *DataMapping) error {
	c := dm.session.DB(dm.databaseName).C(dm.intensityMappingCollection)

	_, err0 := c.Upsert(nil, cfg)
	if err0 != nil {
		return err0
	}
	return nil
}

//获取数据库中的配置信息
func (dm *DataManager) GetGlobalConfigs() (DatabaseConfig, error) {
	c := dm.session.DB(dm.databaseName).C(dm.configCollection)
	dbcfg := DatabaseConfig{}
	err0 := c.Find(&bson.M{}).One(&dbcfg)
	if err0 != nil {
		if err0 == mgo.ErrNotFound {
			return dbcfg, ErrNotFound
		}
		return dbcfg, err0
	}
	return dbcfg, nil
}

//添加配置信息
func (dm *DataManager) CreateGlobalConfigs(cfg *DatabaseConfig) error {
	c := dm.session.DB(dm.databaseName).C(dm.configCollection)

	_, err0 := c.Upsert(nil, cfg)
	if err0 != nil {
		return err0
	}
	return nil
}

//添加速报信息
func (dm *DataManager) ReportAdd(r *Report) error {
	c := dm.session.DB(dm.databaseName).C(dm.reportCollection)
	err := c.Insert(r)
	if err != nil {
		return err
	}
	return nil
}

//保存速报信息
func (dm *DataManager) ReportSave(r *Report) error {
	c := dm.session.DB(dm.databaseName).C(dm.reportCollection)
	_, err := c.Upsert(&bson.M{"eventid": r.EventId}, r)
	if err != nil {
		return err
	}
	return nil
}

//获取速报
func (dm *DataManager) GetReportById(id string) (rep Report, err error) {
	c := dm.session.DB(dm.databaseName).C(dm.reportCollection)
	m := bson.M{"reportid": id}
	err = c.Find(&m).One(&rep)
	if err != nil {
		return Report{}, err
	}
	return rep, nil
}
