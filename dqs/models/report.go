package models

import (
	"time"
)

//报表数据
type Report struct {
	ReportId     string
	EventId      string
	GenerateTime time.Time
	Summary      ReportSummary
	ImageFile    string
	Sended       bool
	Verify       bool
	Event        Event
	SendNumber   int  //发送数
	Valid        bool //是否有效
}

//速报概要
type ReportSummary struct {
	EventTime  string
	AlarmCount int
	Brief      string
	QuakeInfo  string
}
