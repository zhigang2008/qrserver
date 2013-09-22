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
	defaultDatabase         = "dqs"    //默认数据库名称
	defaultDataCollection   = "data"   //默认数据Collection
	defaultDeviceCollection = "device" //默认设备Collection
	defaultUserCollection   = "user"   //默认用户Collection
)

//全局信息
var (
	mux              sync.Mutex
	dbsession        *mgo.Session
	DatabaseName     string
	DataCollection   string
	DeviceCollection string
	UserCollection   string
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
