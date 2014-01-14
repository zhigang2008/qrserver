package models

import (
	//"labix.org/v2/mgo/bson"
	"time"
)

type DeviceFee struct {
	SensorId    string
	NetOperator string    //3G运营商
	NetNo       string    //用户号码
	NetTariff   string    //资费标准
	NetTraffic  float32   //包月流量
	StartDate   string    //开卡日期
	NetQixian   int       //有效期限
	ValidDate   time.Time //有效期
}

type FeeHistory struct {
	SensorId  string
	UserName  string    //操作用户
	NetNo     string    //用户号码
	AddQixian int       //有效期限
	ValidData time.Time //有效期
}
