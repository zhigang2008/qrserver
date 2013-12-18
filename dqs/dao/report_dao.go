package dao

import (
	"dqs/models"
	"dqs/util"
	"labix.org/v2/mgo/bson"
	"time"
)

const (
	ReportTimeLayout = "2006-01-02"
)

//保存事件
func ReportAdd(report *models.Report) (err error) {
	c := GetSession().DB(DatabaseName).C(ReportCollection)
	err = c.Insert(report)
	if err != nil {
		return err
	}
	return nil
}

//获取事件
func GetReportById(sid string) (report models.Report, err error) {
	c := GetSession().DB(DatabaseName).C(ReportCollection)

	err0 := c.Find(bson.M{"reportid": sid}).One(&report)
	if err0 != nil {
		return models.Report{}, err0
	}
	return report, nil
}

//更新事件
func ReportUpdate(report *models.Report) (err error) {
	c := GetSession().DB(DatabaseName).C(ReportCollection)
	err = c.Update(bson.M{"reportid": report.ReportId}, report)
	if err != nil {
		return err
	}
	return nil
}

//事件列表
func ReportList(n int) (*[]models.Report, error) {
	c := GetSession().DB(DatabaseName).C(ReportCollection)

	reports := []models.Report{}

	err := c.Find(bson.M{}).Sort("-generatetime").Limit(n).All(&reports)
	if err != nil {
		return nil, err
	}
	return &reports, nil
}

//获取最近的一个事件
func GetLastReport() (report models.Report, err error) {
	c := GetSession().DB(DatabaseName).C(ReportCollection)

	err0 := c.Find(bson.M{}).Sort("-generatetime").One(&report)
	if err0 != nil {
		return models.Report{}, err0
	}
	return report, nil
}

//分页查询
func ReportPageList(p *util.Pagination) error {
	c := GetSession().DB(DatabaseName).C(ReportCollection)
	reports := []models.Report{}

	//构造查询参数
	m := bson.M{}
	reportid := p.QueryParams["reportid"]
	begintime := p.QueryParams["begintime"]
	endtime := p.QueryParams["endtime"]

	if reportid != nil {
		m["reportid"] = reportid
	}

	timeparam := bson.M{}
	hasTime := false
	if begintime != nil {
		sbtime, ok := begintime.(string)
		if ok {
			btime, _ := time.ParseInLocation(ReportTimeLayout, sbtime, Local)
			timeparam["$gte"] = btime
			hasTime = true
		}
	}
	if endtime != nil {
		setime, ok := endtime.(string)
		if ok {
			etime, _ := time.ParseInLocation(ReportTimeLayout, setime, Local)
			etime = etime.Add(time.Hour * 24)
			timeparam["$lt"] = etime
			hasTime = true
		}
	}
	if hasTime {
		m["generatetime"] = timeparam
	}

	//查询总数
	query := c.Find(&m).Sort("-generatetime")
	count, err := query.Count()
	if err != nil {
		return err
	}
	p.Count = count

	//查找列表
	err = query.Skip(p.SkipNum()).Limit(p.PageSize).All(&reports)
	if err != nil {
		return err
	}
	p.Data = reports
	return nil
}
