package models

import (
	"dqs/dao"
	"labix.org/v2/mgo/bson"
	"time"
)

//报警信息
//服务端后台使用
type AlarmInfo struct {
	SeqNo         string
	SensorId      string
	Longitude     float32
	Latitude      float32
	SiteType      int
	ObserveObject int
	Direction     int
	RegionCode    string
	InitTime      string
	Period        float32
	PGA           float32
	SI            float32
	Length        float32
	CreateTime    time.Time
}

func AlarmList(n int) []AlarmInfo {
	c := dao.GetSession().DB(dao.DatabaseName).C(dao.DataCollection)
	alarms := []AlarmInfo{}

	err := c.Find(&bson.M{}).Sort("-createtime").Limit(n).All(&alarms)
	if err != nil {
		panic(err)
		return nil
	}
	return alarms
}
