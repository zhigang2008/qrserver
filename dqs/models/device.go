package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

//服务端后台使用
type DeviceInfo struct {
	Id           bson.ObjectId "_id"
	SensorId     string
	RegisterTime time.Time
	OffTime      time.Time
	Online       bool
	UpdateTime   time.Time
	RemoteAddr   string
	SetParams    SensorInfo
	CustomParams CustomDefineParams
}

//传感器参数结构
//服务端后台使用
type SensorInfo struct {
	//--基本参数--
	SensorId      string  //传感器编号
	SiteName      string  //站点名称
	Longitude     float32 //台站经度
	Latitude      float32 //台站纬度
	SiteType      int     //场地类型
	ObserveObject int     //观测对象
	Accelerometer int     //加速度计型号
	Direction     int     //安装方向
	RangeType     int     //量程选择
	Period        float32 //采样周期
	RegionCode    string  //行政区划代码
	Custom1       string  //预留
	Custom2       string  //预留
	//--触发参数--
	PGATrigger          int     //PGA触发
	PGATrgThreshold     float32 //PGA阀值
	SITrigger           int     //SI触发
	SITrgThreshold      float32 //SI阀值
	CombTrigger         int     //组合触发
	ReserveTrigger      int     //预留触发
	ReserveTrgThreshold float32 //预留阀值
	//--报警参数--
	PGAAlert              int     //PGA报警
	PGAAlertThreshold     float32 //PGA报警阀值
	SIAlert               int     //SI报警
	SIAlertThreshold      float32 //SI报警阀值
	CombAlert             int     //组合报警
	ReserveAlert          int     //预留报警
	ReserveAlertThreshold float32 //预留报警阀值
	//--输出参数--
	DA1 int //DA输出1
	DA2 int //DA输出2
	IO1 int //IO输出1
	IO2 int //IO输出2
}

//用户自定义数据
type CustomDefineParams struct {
	InstallDate      string       //安装日期
	Installer        string       //安装人员
	Contact          string       //联系人
	Telphone         string       //固话
	Mobile           string       //移动电话
	MaintainDate     string       //维护日期
	MaintainPeole    string       //维护人
	NetType          string       //网络类型
	WirelessTypeInfo WirelessInfo //3G类型信息
	WireTypeInfo     WireInfo     //有线类型信息
	NotUse           bool
}

//3G类型的信息
type WirelessInfo struct {
	NetOperator string  //3G运营商
	NetNo       string  //用户号码
	NetTariff   string  //资费标准
	NetTraffic  float32 //包月流量
	StartDate   string  //开卡日期
	NetQixian   float32 //有效期
}

//有线类型信息
type WireInfo struct {
	IPType    string //动态IP-D;静态IP-S
	IPAddress string //IP地址
	IPMask    string //子网掩码
	GateWay   string //网关
	DNS       string //DNS
}
