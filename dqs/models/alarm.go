package models

import (
	//"labix.org/v2/mgo/bson"
	"encoding/xml"
	"time"
)

//报警信息
//服务端后台使用
type AlarmInfo struct {
	SeqNo         string    `xml:"-"`
	SensorId      string    `xml:"SensorId"`
	Longitude     float32   `xml:"Longitude"`
	Latitude      float32   `xml:"Latitude"`
	SiteType      int       `xml:"SiteType"`
	ObserveObject int       `xml:"ObserveObject"`
	Direction     int       `xml:"Direction"`
	RegionCode    string    `xml:"RegionCode"`
	InitTime      string    `xml:"InitTime"`
	InitRealTime  time.Time `xml:"-"`
	Period        float32   `xml:"Period"`
	PGA           float32   `xml:"PGA"`
	SI            float32   `xml:"SI"`
	Intensity     int       `xml:"Intensity"` //仪器烈度值
	Length        float32   `xml:"-"`
	CreateTime    time.Time `xml:"-"`
	HasWaveInfo   bool      `xml:"-"`
	EventId       string    `xml:"-"` //事件编号
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

//回送的报警数据
type AlarmDataList struct {
	XMLName xml.Name `xml:"DataList"`
	EventId string
	Alarms  []AlarmInfo `xml:"Data"`
}
