package dao

import (
	"dqs/models"
	"dqs/util"
	//"labix.org/v2/mgo"
	"fmt"
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
	userid := p.QueryParams["userid"]
	actcontent := p.QueryParams["actcontent"]
	begintime := p.QueryParams["begintime"]
	endtime := p.QueryParams["endtime"]

	fmt.Println(userid)
	fmt.Println(userid)

	if userid != nil {
		m["userid"] = userid
	}
	if actcontent != nil {
		v, ok := actcontent.(string)
		if ok {
			content := "/" + v + "/"
			m["actcontent"] = bson.M{"$regex": content}
		}
	}
	if begintime != nil {
		m["acttime"] = bson.M{"$gte": begintime}
	}
	if endtime != nil {
		m["acttime"] = bson.M{"$lte": endtime}
	}
	if begintime != nil && endtime != nil {
		m["acttime"] = bson.M{"$gte": begintime, "$lte": endtime}
	}

	for k, v := range m {
		fmt.Printf("[%s]=[%s]", k, v)
	}
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
