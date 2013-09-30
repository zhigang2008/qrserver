package dao

import (
	"dqs/models"
	"dqs/util"
	//"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

//添加审计日志
func AddAuditLog(audit models.AuditLog) bool {
	c := GetSession().DB(DatabaseName).C(AuditCollection)
	err := c.Insert(&audit)
	if err != nil {
		return false
	}
	return true
}

//审计日志列表
func AuditList(p *util.Pagination) error {
	c := GetSession().DB(DatabaseName).C(AuditCollection)
	audits := []models.AuditLog{}
	//构造查询参数
	m := bson.M{}

	//查询总数
	query := c.Find(&m).Sort("acttime")
	count, err := query.Count()
	if err != nil {
		return err
	}
	p.Count = count

	//查找设备
	err = query.Skip(p.SkipNum()).Limit(p.PageSize).All(&audits)
	if err != nil {
		return err
	}
	p.Data = audits
	return nil
}
