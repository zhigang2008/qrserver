package dao

import (
	"dqs/models"
	"dqs/util"
	"labix.org/v2/mgo/bson"
)

//报警信息列表
func AlarmList(p *util.Pagination) error {
	c := GetSession().DB(DatabaseName).C(DataCollection)
	alarms := []models.AlarmInfo{}

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
