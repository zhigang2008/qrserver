package quickserver

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"time"
)

const (
	I_Register    = 'z'
	O_ConfRead    = 'g'
	O_ConfSet     = 's'
	I_Alert       = 'a'
	I_AlertUp     = 'A'
	O_Record      = 'r'
	I_RecordAlert = 'a'
	I_RecordData  = 'r'
	I_Status      = 'g'
)

/*
//modbus 协议结构
type ModBus struct {
	Addr       byte    //地址
	FunCode    byte    //功能码
	Datalength [2]byte //数据长度
	Data       []byte  //数据
	CRC        [2]byte //CRC校验
}
*/
//传感器参数结构
//主要用来与DLL交互
type RetData struct {
	//--基本参数--
	SensorId      [11]byte //传感器编号
	SiteName      [11]byte //站点名称
	Longitude     float32  //台站经度
	Latitude      float32  //台站纬度
	SiteType      int      //场地类型
	ObserveObject int      //观测对象
	Accelerometer int      //加速度计型号
	Direction     int      //安装方向
	RangeType     int      //量程选择
	Period        float32  //采样周期
	RegionCode    [7]byte  //行政区划代码
	Custom1       [9]byte  //预留
	Custom2       [9]byte  //预留
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

//突发数据
//用来与DLL交互
type FlashData struct {
	SeqNo         [11]byte //记录编号
	SensorId      [11]byte //传感器编号
	Longitude     float32  //经度
	Latitude      float32  //纬度
	SiteType      int      //场地类型
	ObserveObject int      //观测对象
	Direction     int      //安装方向
	RegionCode    [7]byte  //行政区域编码
	InitTime      [6]byte  //初始时刻
	Period        float32  //采用周期
	PGA           float32  //PGA值
	SI            float32  //SI值
	Length        float32  //记录长度
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

//波形记录数据
type WaveInfo struct {
	Id       bson.ObjectId "_id"
	SensorId string        //设备ID
	SeqNo    string
	Alarm    AlarmInfo
	X_data   [6000]int16
	Y_data   [6000]int16
	Z_data   [6000]int16
	LUD      time.Time //最后更新时间
}

//设备信息
//服务端后台使用
type DeviceInfo struct {
	SensorId     string
	RegisterTime time.Time
	OffTime      time.Time
	Online       bool
	UpdateTime   time.Time
	RemoteAddr   string
	SetParams    SensorInfo
}

//传感器参数转为服务端对象
func RetData2SensorInfo(ret *RetData) *SensorInfo {
	s := new(SensorInfo)
	s.SensorId = string(ret.SensorId[0:10])
	s.SiteName = string(ret.SiteName[0:10])
	s.Longitude = ret.Longitude
	s.Latitude = ret.Latitude
	s.SiteType = ret.SiteType
	s.ObserveObject = ret.ObserveObject
	s.Accelerometer = ret.Accelerometer
	s.Direction = ret.Direction
	s.RangeType = ret.RangeType
	s.Period = ret.Period
	s.RegionCode = string(ret.RegionCode[0:6])
	s.Custom1 = string(ret.Custom1[0:8])
	s.Custom2 = string(ret.Custom2[0:8])
	s.PGATrigger = ret.PGATrigger
	s.PGATrgThreshold = ret.PGATrgThreshold
	s.SITrigger = ret.SITrigger
	s.SITrgThreshold = ret.SITrgThreshold
	s.CombTrigger = ret.CombTrigger
	s.ReserveTrigger = ret.ReserveTrigger
	s.ReserveTrgThreshold = ret.ReserveTrgThreshold
	s.PGAAlert = ret.PGAAlert
	s.PGAAlertThreshold = ret.PGAAlertThreshold
	s.SIAlert = ret.SIAlert
	s.SIAlertThreshold = ret.SIAlertThreshold
	s.CombAlert = ret.CombAlert
	s.ReserveAlert = ret.ReserveAlert
	s.ReserveAlertThreshold = ret.ReserveAlertThreshold
	s.DA1 = ret.DA1
	s.DA2 = ret.DA2
	s.IO1 = ret.IO1
	s.IO2 = ret.IO2
	return s
}

//服务器后端传感器对象转化为前端传感器数据
func SensorInfo2RetData(sensor *SensorInfo) *RetData {
	r := new(RetData)
	for i, b := 0, []byte(sensor.SensorId); i < len(sensor.SensorId); i++ {
		r.SensorId[i] = b[i]
	}
	r.SensorId[10] = byte(0)

	for i, b := 0, []byte(sensor.SiteName); i < len(sensor.SiteName); i++ {
		r.SiteName[i] = b[i]
	}
	r.SiteName[10] = byte(0)

	r.Longitude = sensor.Longitude
	r.Latitude = sensor.Latitude
	r.SiteType = sensor.SiteType
	r.ObserveObject = sensor.ObserveObject
	r.Accelerometer = sensor.Accelerometer
	r.Direction = sensor.Direction
	r.RangeType = sensor.RangeType
	r.Period = sensor.Period
	for i, b := 0, []byte(sensor.RegionCode); i < len(sensor.RegionCode); i++ {
		r.RegionCode[i] = b[i]
	}
	r.RegionCode[6] = byte(0)

	for i, b := 0, []byte(sensor.Custom1); i < len(sensor.Custom1); i++ {
		r.Custom1[i] = b[i]
	}
	r.Custom1[8] = byte(0)

	for i, b := 0, []byte(sensor.Custom2); i < len(sensor.Custom2); i++ {
		r.Custom2[i] = b[i]
	}
	r.Custom2[8] = byte(0)
	r.PGATrigger = sensor.PGATrigger
	r.PGATrgThreshold = sensor.PGATrgThreshold
	r.SITrigger = sensor.SITrigger
	r.SITrgThreshold = sensor.SITrgThreshold
	r.CombTrigger = sensor.CombTrigger
	r.ReserveTrigger = sensor.ReserveTrigger
	r.ReserveTrgThreshold = sensor.ReserveTrgThreshold
	r.PGAAlert = sensor.PGAAlert
	r.PGAAlertThreshold = sensor.PGAAlertThreshold
	r.SIAlert = sensor.SIAlert
	r.SIAlertThreshold = sensor.SIAlertThreshold
	r.CombAlert = sensor.CombAlert
	r.ReserveAlert = sensor.ReserveAlert
	r.ReserveAlertThreshold = sensor.ReserveAlertThreshold
	r.DA1 = sensor.DA1
	r.DA2 = sensor.DA2
	r.IO1 = sensor.IO1
	r.IO2 = sensor.IO2
	//fmt.Printf("[SensorInfo2RetData]sensor=%x;\nretdata=%x", sensor, r)
	return r
}

//突发数据转化为报警信息
func FlashData2AlarmInfo(f *FlashData) *AlarmInfo {
	a := AlarmInfo{}
	a.SeqNo = string(f.SeqNo[0:10])
	a.SensorId = string(f.SensorId[0:10])
	a.Longitude = f.Longitude
	a.Latitude = f.Latitude
	a.SiteType = f.SiteType
	a.ObserveObject = f.ObserveObject
	a.Direction = f.Direction
	a.RegionCode = string(f.RegionCode[0:6])
	//采用16进值数字,6字节表示年月日时分秒
	a.InitTime = fmt.Sprintf("%x", f.InitTime[0:6])
	a.Period = f.Period
	a.PGA = f.PGA
	a.SI = f.SI
	a.Length = f.Length
	a.CreateTime = time.Now()
	return &a
}
