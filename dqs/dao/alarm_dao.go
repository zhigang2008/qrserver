package dao

import (
	"dqs/models"
	"dqs/util"
	"fmt"
	log "github.com/cihub/seelog"
	"labix.org/v2/mgo/bson"
	"time"
)

const (
	alarmTimeLayout = "200601021504"
)

var local *time.Location

func init() {
	l, err := time.LoadLocation("Local")
	if err == nil {
		local = l
		//fmt.Println(local)
	} else {
		fmt.Println("init:" + err.Error())
	}
}

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

//实时报警信息列表
func GetRealtimeAlarm(timestep int64) ([]models.AlarmInfo, error) {
	c := GetSession().DB(DatabaseName).C(DataCollection)
	alarms := []models.AlarmInfo{}

	//构造查询参数
	m := bson.M{}

	timeparam := bson.M{}
	now := time.Now()
	duration := time.Duration(timestep) * time.Minute
	btime := now.Add(-duration)
	timeparam["$gte"] = btime
	m["createtime"] = timeparam

	//查找数据
	err := c.Find(&m).All(&alarms)
	if err != nil {
		return alarms, err
	}
	return alarms, nil
}

//导出报警信息列表
func ExportAlarms(sid, begintime, endtime string) ([]models.AlarmInfo, error) {
	c := GetSession().DB(DatabaseName).C(DataCollection)
	alarms := []models.AlarmInfo{}

	//构造查询参数
	m := bson.M{}

	if sid != "" {
		m["sensorid"] = sid
	}

	timeparam := bson.M{}
	hasTime := false

	if begintime != "" {
		btime, err := time.ParseInLocation(alarmTimeLayout, begintime, Local)
		if err == nil {
			timeparam["$gte"] = btime
			hasTime = true
		} else {
			log.Warnf("开始日期解析错误[%s]:%s", begintime, err.Error())
		}
	}
	if endtime != "" {
		etime, err := time.ParseInLocation(alarmTimeLayout, endtime, Local)
		if err == nil {
			timeparam["$lte"] = etime
			hasTime = true
		} else {
			log.Warnf("结束日期解析错误[%s]:%s", endtime, err.Error())
		}
	}

	if hasTime {
		m["createtime"] = timeparam
	}

	//查找数据
	err := c.Find(&m).All(&alarms)
	if err != nil {
		return alarms, err
	}
	return alarms, nil
}

//获取波形图
func GetWaveInfoById(oid string) (models.WaveInfo, error) {
	c := GetSession().DB(DatabaseName).C(WaveCollection)
	wave := models.WaveInfo{}

	//查找数据
	err := c.Find(&bson.M{"_id": bson.ObjectIdHex(oid)}).One(&wave)
	if err != nil {
		return wave, err
	}
	return wave, nil
}

//根据设备和序号查找波形数据
func GetWaveInfo(SensorId string, SeqNo string) (models.WaveInfo, error) {
	c := GetSession().DB(DatabaseName).C(WaveCollection)
	wave := models.WaveInfo{}

	//构造查询参数
	m := bson.M{}
	m["sensorid"] = SensorId
	m["SeqNo"] = SeqNo

	//查找数据
	err := c.Find(&m).One(&wave)
	if err != nil {
		return wave, err
	}
	return wave, nil
}
