package models

import (
	//"labix.org/v2/mgo/bson"
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
	Intensity     float32 //仪器烈度值
	Length        float32
	CreateTime    time.Time
	HasWaveInfo   bool
}

//报警信息列表
type AlarmList struct {
	SensorId  string
	BeginTime string
	EndTime   string
	Alarms    []AlarmInfo
}

//波形记录数据
type WaveInfo struct {
	//Id       bson.ObjectId "_id"
	SensorId string //设备ID
	SeqNo    string
	Alarm    AlarmInfo
	X_data   [6000]int16
	Y_data   [6000]int16
	Z_data   [6000]int16
	LUD      time.Time //最后更新时间
}
