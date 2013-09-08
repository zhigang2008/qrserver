package dao

import (
	log "github.com/cihub/seelog"
	"labix.org/v2/mgo"
	//	"labix.org/v2/mgo/bson"
	"strconv"
	"sync"
)

const (
	defaultDatabase         = "dqs"    //默认数据库名称
	defaultDataCollection   = "data"   //默认数据Collection
	defaultDeviceCollection = "device" //默认设备Collection
)

var (
	mux              sync.Mutex
	dbsession        *mgo.Session
	DatabaseName     string
	DataCollection   string
	DeviceCollection string
)

func Init(host string, port int, dbname, datacol, devicecol string) (err error) {
	dbsession, err = mgo.Dial(host + ":" + strconv.Itoa(port))
	if err != nil {
		log.Criticalf("不能创建数据库连接[%s:%d]:%s", host, port, err.Error())
		return err
	}
	dbsession.SetMode(mgo.Monotonic, true)

	DatabaseName = defaultDatabase
	DataCollection = defaultDataCollection
	DeviceCollection = defaultDeviceCollection

	if dbname != "" {
		DatabaseName = dbname
	}
	if datacol != "" {
		DataCollection = datacol
	}
	if devicecol != "" {
		DeviceCollection = devicecol
	}
	return nil
}

func GetSession() *mgo.Session {

	return dbsession
}
func Close() {
	dbsession.Close()
}
