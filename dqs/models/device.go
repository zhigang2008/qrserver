package models

import (
	"dqs/dao"
	"dqs/util"
	"labix.org/v2/mgo/bson"
	"time"
)

//服务端后台使用
type DeviceInfo struct {
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

type CustomDefineParams struct {
	NetType int
}

func DeviceList(p *util.Pagination) error {

	c := dao.GetSession().DB(dao.DatabaseName).C(dao.DeviceCollection)
	devices := []DeviceInfo{}
	//构造查询参数
	m := bson.M{}
	for k, v := range p.QueryParams {
		m[k] = v
	}

	//查询总数
	query := c.Find(&m).Sort("-registertime")
	count, err := query.Count()
	if err != nil {
		return err
	}
	p.Count = count

	//查找设备
	err = query.Skip(p.SkipNum()).Limit(p.PageSize).All(&devices)
	if err != nil {
		return err
	}
	p.Data = devices
	return nil

}

func GetDevice(sid string) DeviceInfo {
	c := dao.GetSession().DB(dao.DatabaseName).C(dao.DeviceCollection)
	device := DeviceInfo{}
	//查找设备
	err := c.Find(&bson.M{"sensorid": sid}).One(&device)
	if err != nil {
		device = DeviceInfo{}
	}
	return device
}
