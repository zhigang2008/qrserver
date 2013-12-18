package models

import (
	"time"
)

//报表数据
type Report struct {
	ReportId     string
	EventId      string
	GenerateTime time.Time
	Summary      map[string]interface{}
	ImageFile    string
	Sended       bool
	Verify       bool
	Event        Event
	SendNumber   int //发送数
}
