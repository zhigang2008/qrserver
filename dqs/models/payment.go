package models

import (
	//"labix.org/v2/mgo/bson"
	"time"
)

//待充值记录
type DevicePayment struct {
	SensorId      string
	SiteAliasName string
	NetOperator   string    //3G运营商
	NetNo         string    //用户号码
	NetTariff     string    //资费标准
	NetTraffic    float32   //包月流量
	StartDate     string    //开卡日期
	LeftDate      int       //剩余期限
	ValidDate     time.Time //有效期
}

//充值历史记录
type PaymentHistory struct {
	SensorId    string
	UserId      string    //操作用户ID
	UserName    string    //操作用户
	NetOperator string    //3G运营商
	NetNo       string    //用户号码
	AddQixian   int       //添加的期限
	ValidDate   time.Time //有效期
	OperateTime time.Time //操作时间
}
