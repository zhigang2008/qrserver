//负责为http server 创建数据连接
package dao

import (
	log "github.com/cihub/seelog"
	"labix.org/v2/mgo"
	//	"labix.org/v2/mgo/bson"
	"strconv"
	"sync"
)

//默认数据库信息
const (
	defaultDatabase              = "dqs"          //默认数据库名称
	defaultDataCollection        = "data"         //默认数据Collection
	defaultWaveCollection        = "wavedata"     //默认波形记录Collection
	defaultDeviceCollection      = "device"       //默认设备Collection
	defaultUserCollection        = "user"         //默认用户Collection
	defaultAuditCollection       = "audit"        //默认审计Collection
	defaultEventCollection       = "event"        //默认事件Collection
	defaultEventSignalCollection = "eventsignal"  //默认事件信号Collection
	defaultIntensityCollection   = "intensitymap" //默认事件信号Collection
	defaultConfigCollection      = "configs"      //默认配置信息表
	defaultReportCollection      = "reports"      //默认配置信息表
)

//全局信息
var (
	mux                        sync.Mutex
	dbsession                  *mgo.Session
	DatabaseName               string
	DataCollection             string
	WaveCollection             string
	DeviceCollection           string
	UserCollection             string
	AuditCollection            string
	EventCollection            string
	EventSignalCollection      string
	IntensityMappingCollection string
	ConfigCollection           string
	ReportCollection           string
)

//初始化数据库连接
func Init(host string, port int, dbname, datacol, devicecol, usercol string) (err error) {
	mux.Lock()
	dbsession, err = mgo.Dial(host + ":" + strconv.Itoa(port))
	if err != nil {
		log.Criticalf("不能创建数据库连接[%s:%d]:%s", host, port, err.Error())
		return err
	}
	dbsession.SetMode(mgo.Monotonic, true)
	mux.Unlock()

	DatabaseName = defaultDatabase
	DataCollection = defaultDataCollection
	DeviceCollection = defaultDeviceCollection
	UserCollection = defaultUserCollection
	AuditCollection = defaultAuditCollection
	WaveCollection = defaultWaveCollection
	EventCollection = defaultEventCollection
	EventSignalCollection = defaultEventSignalCollection
	IntensityMappingCollection = defaultIntensityCollection
	ConfigCollection = defaultConfigCollection
	ReportCollection = defaultReportCollection

	if dbname != "" {
		DatabaseName = dbname
	}
	if datacol != "" {
		DataCollection = datacol
	}
	if devicecol != "" {
		DeviceCollection = devicecol
	}
	if usercol != "" {
		UserCollection = usercol
	}

	return nil
}

//获取数据session
func GetSession() *mgo.Session {
	return dbsession
}

//关闭数据连接
func Close() {
	dbsession.Close()
}
