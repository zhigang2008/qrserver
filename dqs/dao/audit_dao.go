package dao

import (
	"dqs/models"
	"dqs/util"
	//"labix.org/v2/mgo"
	"fmt"
	"labix.org/v2/mgo/bson"
	"time"
)

const (
	AuditTimeLayout = "2006-01-02"
)

var Local *time.Location

func init() {
	l, err := time.LoadLocation("Local")
	if err == nil {
		Local = l
		fmt.Println(Local)
	} else {
		fmt.Println("init:" + err.Error())
	}
}

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

	if userid != nil {
		m["userid"] = userid
	}
	if actcontent != nil {
		v, ok := actcontent.(string)
		if ok {
			regex := bson.RegEx{v, "i"}
			m["actcontent"] = bson.M{"$regex": regex}
		}
	}

	timeparam := bson.M{}
	hasTime := false
	if begintime != nil {
		sbtime, ok := begintime.(string)
		if ok {
			btime, _ := time.ParseInLocation(AuditTimeLayout, sbtime, Local)
			timeparam["$gte"] = btime
			hasTime = true
		}
	}
	if endtime != nil {
		setime, ok := endtime.(string)
		if ok {
			etime, _ := time.ParseInLocation(AuditTimeLayout, setime, Local)
			etime = etime.Add(time.Hour * 24)
			timeparam["$lt"] = etime
			hasTime = true
		}
	}
	if hasTime {
		m["acttime"] = timeparam
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
