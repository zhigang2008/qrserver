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
	Id          string
	Time        time.Time //震情时间
	Longitude   float32   //震中位置
	Latitude    float32
	Level       int       //震级
	ReceiveTime time.Time //信号接收时间
}

//传入的地震事件
type EarthQuake struct {
	Id        string
	Time      string  //震情时间
	Longitude float32 //震中位置
	Latitude  float32
	Level     int //震级
}

//反馈的信息
type Feedback struct {
	XMLName xml.Name `xml:"xml"`
	Ok      bool     `xml:"ok"`
	Message string   `xml:"message"`
}
