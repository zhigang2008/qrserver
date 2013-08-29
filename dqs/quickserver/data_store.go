package quickserver

import (
	//	"fmt"
	log "github.com/cihub/seelog"
	"labix.org/v2/mgo"
	"strconv"

	"labix.org/v2/mgo/bson"
)

const (
	defaultDatabase         = "dqs"    //默认数据库名称
	defaultDataCollection   = "data"   //默认数据Collection
	defaultDeviceCollection = "device" //默认设备Collection
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
func (dm *DataManager) DataSave(data *DataAlert) (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.dataCollection)
	err = c.Insert(data)
	if err != nil {

		return err
	}
	return nil
}

//设备注册
func (dm *DataManager) DeviceRegister(device *DeviceInfo) (err error) {
	c := dm.session.DB(dm.databaseName).C(dm.deviceCollection)

	changeInfo, err0 := c.Upsert(&bson.M{"sensorid": device.SensorId}, &device)
	if err0 != nil {
		log.Infof("updated:%de", changeInfo.Updated)
		return err0
	}

	return nil
}

//关闭数据库连接
func (dm *DataManager) DataClose() {
	dm.session.Close()

}
