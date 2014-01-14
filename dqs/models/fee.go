package models

import (
	//"labix.org/v2/mgo/bson"
	"time"
)

//缴费设备清单
type DevicePayment struct {
	SensorId    string
	NetOperator string    //3G运营商
	NetNo       string    //用户号码
	NetTariff   string    //资费标准
	NetTraffic  float32   //包月流量
	StartDate   string    //开卡日期
	NetQixian   int       //有效期限
	ValidDate   time.Time //有效期
	DateLeave   int       //剩余有效天数
}

//缴费历史
type PaymentHistory struct {
	SensorId    string
	UserName    string    //操作用户
	NetNo       string    //用户号码
	AddQixian   int       //有效期限
	ValidData   time.Time //有效期
	OperateTime time.Time //操作时间
}
