package quickserver

import (
	"fmt"
	log "github.com/cihub/seelog"
	"labix.org/v2/mgo"
	"strconv"

//	"labix.org/v2/mgo/bson"
)

const (
	defaultDatabase         = "dqs"
	defaultDataCollection   = "data"
	defaultDeviceCollection = "device"
)

var (
	session        *mgo.Session
	databaseName   = defaultDatabase
	dataCollection = defaultDataCollection
)

//报警信息
type DataAlert struct {
	SeqNo         string
	SensorId      string
	Longitude     float32
	Latitude      float32
	SiteType      byte
	ObserveObject byte
	Direction     byte
	RegionCode    string
	InitTime      string
	Period        float32
	PGA           float32
	SI            float32
	Length        float32
}

//初始化数据库连接
func InitDatabase(conf ServerConfig) {
	session, err := mgo.Dial(conf.Database.Host + ":" + strconv.Itoa(conf.Database.Port))
	if err != nil {
		log.Criticalf("不能创建数据库连接:%s", err.Error())
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	fmt.Println("创建了数据连接:")
}

//保存报警信息
func DataSave(data *DataAlert) (err error) {
	c := session.DB(databaseName).C(dataCollection)
	err = c.Insert(data)
	if err != nil {
		panic(err)
		return err
	}
	fmt.Println("数据保存成功")
	return nil
}

//关闭数据库连接
func DataClose() {
	session.Close()

}
