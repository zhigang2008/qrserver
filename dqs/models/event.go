package models

import (
	"time"
)

//地震事件
type Event struct {
	EventId    string
	EventTime  time.Time
	AlarmCount int
	CreateTime time.Time
	IsConfirm  bool
	SignalId   string
	Signal     EventSignal
}

//地震事件确认信号
type EventSignal struct {
	Id          string
	time        time.Time //震情时间
	Longitude   float32   //震中位置
	Latitude    float32
	Level       int       //震级
	ReceiveTime time.Time //信号接收时间
}
