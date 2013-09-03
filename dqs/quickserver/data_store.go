package quickserver

import (
	//	"fmt"
	log "github.com/cihub/seelog"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strconv"
	"sync"
	"time"
)

const (
	defaultDatabase         = "dqs"    //默认数据库名称
	defaultDataCollection   = "data"   //默认数据Collection
	defaultDeviceCollection = "device" //默认设备Collection
)

var (
	mux sync.Mutex
)

//数据库连接服务
type DataManager struct {
	session          *mgo.Session
	databaseName     string
	dataCollection   string
	deviceCollection string
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
		session:          session1,
		databaseName:     defaultDatabase,
		dataCollection:   defaultDataCollection,
		deviceCollection: defaultDeviceCollection,
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

	log.Infof("datamanager %v \n", dataManager)
	return dataManager, nil
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
	err0 := c.Update(&bson.M{"sensorid": device.SensorId}, &device)
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
	change := bson.M{"$set": bson.M{"registertime": time.Now(), "online": true, "updatetime": time.Now()}}
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
	change := bson.M{"$set": bson.M{"offlinetime": time.Now(), "online": false, "updatetime": time.Now()}}
	err0 := c.Update(colQuerier, change)
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
	err := c.Find(&bson.M{}).Sort("-registertime").Limit(n).All(&devices)
	if err != nil {
		return nil, err
	}
	return &devices, nil
}

//查找报警信息
func (dm *DataManager) AlarmList(n int) (*[]AlarmInfo, error) {
	c := dm.session.DB(dm.databaseName).C(dm.dataCollection)

	alarms := []AlarmInfo{}

	err := c.Find(&bson.M{}).Sort("-createtime").Limit(n).All(&alarms)
	if err != nil {
		return nil, err
	}
	return &alarms, nil
}

//关闭数据库连接
func (dm *DataManager) DataClose() {
	dm.session.Close()

}
