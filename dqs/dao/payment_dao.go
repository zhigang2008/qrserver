package dao

import (
	"dqs/models"
	"dqs/util"
	//"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

const (
	PaymentTimeLayout = "2006-01-02"
)

//获取待付费信息
func GetAllPayment() ([]models.DevicePayment, error) {
	c := GetSession().DB(DatabaseName).C(PaymentCollection)
	payments := []models.DevicePayment{}

	err := c.Find(nil).All(&payments)
	if err != nil {
		payments = []models.DevicePayment{}
		return payments, err
	}
	return payments, nil
}

//待付费信息
func GetPagePayments(p *util.Pagination) error {
	c := GetSession().DB(DatabaseName).C(PaymentCollection)
	fees := []models.DeviceFee{}
	//构造查询参数
	m := bson.M{}
	sensorid := p.QueryParams["SensorId"]
	netoperator := p.QueryParams["NetOperator"]
	netno := p.QueryParams["NetNo"]
	begintime := p.QueryParams["begintime"]
	endtime := p.QueryParams["endtime"]

	if sensorid != nil {
		m["sensorid"] = sensorid
	}
	if netoperator != nil {
		m["netoperator"] = netoperator
	}
	if netno != nil {
		m["netno"] = netno
	}

	timeparam := bson.M{}
	hasTime := false
	if begintime != nil {
		sbtime, ok := begintime.(string)
		if ok {
			btime, _ := time.ParseInLocation(PaymentTimeLayout, sbtime, Local)
			timeparam["$gte"] = btime
			hasTime = true
		}
	}
	if endtime != nil {
		setime, ok := endtime.(string)
		if ok {
			etime, _ := time.ParseInLocation(PaymentTimeLayout, setime, Local)
			etime = etime.Add(time.Hour * 24)
			timeparam["$lt"] = etime
			hasTime = true
		}
	}
	if hasTime {
		m["validdate"] = timeparam
	}
	//查询总数
	query := c.Find(&m).Sort("leftdate")
	count, err := query.Count()
	if err != nil {
		return err
	}
	p.Count = count

	//查找数据
	err = query.Skip(p.SkipNum()).Limit(p.PageSize).All(&fees)
	if err != nil {
		return err
	}
	p.Data = fees
	return nil
}

//清空待付费数据
func ClearPaymentInfo() error {
	err := GetSession().DB(DatabaseName).C(PaymentCollection).DropCollection()
	return err
}

//添加待付费信息
func AddPayment(fee *models.DeviceFee) error {
	c := GetSession().DB(DatabaseName).C(PaymentCollection)
	return c.Insert(fee)
}
