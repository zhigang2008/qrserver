package models

import (
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
