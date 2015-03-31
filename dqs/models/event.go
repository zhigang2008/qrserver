package models

import (
	"encoding/xml"
	"time"
)

//地震事件
type Event struct {
	EventId      string
	EventTime    time.Time
	EventTimeStr string
	AlarmCount   int
	MaxLevel     int //最高震级
	CreateTime   time.Time
	IsConfirm    bool
	SignalId     string
	Signal       EventSignal
}

//地震事件确认信号
type EventSignal struct {
	Id             string
	Time           time.Time //震情时间
	Longitude      float32   //震中位置
	Latitude       float32
	Level          float32   //震级
	ReceiveTime    time.Time //信号接收时间
	EventId        string    //地震事件Id
	CODE           string    //台网编号
	CNAME          string    //台网名称
	DEPTH          float32   //深度
	LOCATION_CNAME string    //地点名称

}

//传入的地震事件
type EarthQuake struct {
	XMLName        xml.Name `xml:"ROOT"`
	EVENT_ID       string   `xml:"EVENT_ID"`
	Time           string   `xml:"O_TIME"` //震情时间
	Longitude      float32  `xml:"LON"`    //震中位置
	Latitude       float32  `xml:"LAT"`
	Level          float32  `xml:"M"` //震级
	CODE           string   `xml:"CODE"`
	CNAME          string   `xml:"CNAME"`
	DEPTH          float32  `xml:"DEPTH"`
	LOCATION_CNAME string   `xml:"LOCATION_CNAME"`
}

//反馈的信息
type Feedback struct {
	XMLName xml.Name `xml:"xml"`
	Ok      bool     `xml:"ok"`
	Message string   `xml:"message"`
}
