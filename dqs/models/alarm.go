package models

import (
	"dqs/dao"
	"dqs/util"
	"labix.org/v2/mgo/bson"
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

//报警信息列表
func AlarmList(p *util.Pagination) error {
	c := dao.GetSession().DB(dao.DatabaseName).C(dao.DataCollection)
	alarms := []AlarmInfo{}

	//构造查询参数
	m := bson.M{}
	for k, v := range p.QueryParams {
		m[k] = v
	}

	//查询总数
	query := c.Find(&m).Sort("-createtime")
	count, err := query.Count()
	if err != nil {
		return err
	}
	p.Count = count

	//查找数据
	err = query.Skip(p.SkipNum()).Limit(p.PageSize).All(&alarms)
	if err != nil {
		return err
	}
	p.Data = alarms
	return nil
}
