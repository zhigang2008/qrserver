package dao

import (
	"dqs/models"
	"dqs/util"
	"errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

//获取所有设备列表
func GetAllDevices() ([]models.DeviceInfo, error) {
	c := GetSession().DB(DatabaseName).C(DeviceCollection)
	devices := []models.DeviceInfo{}
	//构造查询参数
	m := bson.M{}

	//查询总数
	err := c.Find(&m).Sort("-registertime").All(&devices)
	if err != nil {
		return devices, err
	}
	return devices, nil
}

//获取所有设备在用列表
func GetAllValidDevices() ([]models.DeviceInfo, error) {
	c := GetSession().DB(DatabaseName).C(DeviceCollection)
	devices := []models.DeviceInfo{}
	//构造查询参数
	m := bson.M{"customparams.notuse": false}
	//查询总数
	err := c.Find(&m).All(&devices)
	if err != nil {
		return devices, err
	}
	return devices, nil
}

//设备列表
func DeviceList(p *util.Pagination) error {
	c := GetSession().DB(DatabaseName).C(DeviceCollection)
	devices := []models.DeviceInfo{}
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

//根据编号查找设备
func GetDevice(sid string) models.DeviceInfo {
	c := GetSession().DB(DatabaseName).C(DeviceCollection)
	device := models.DeviceInfo{}
	//查找设备
	err := c.Find(&bson.M{"sensorid": sid}).One(&device)
	if err != nil {
		device = models.DeviceInfo{}
	}
	return device
}

//添加设备信息
func AddDevice(dev *models.DeviceInfo) error {
	c := GetSession().DB(DatabaseName).C(DeviceCollection)
	//先查找,是否存在
	device := models.DeviceInfo{}
	err := c.Find(&bson.M{"sensorid": dev.SensorId}).One(&device)
	if err != nil && err != mgo.ErrNotFound {
		return err
	}
	if device.SensorId != "" {
		return errors.New("已存在该设备")
	}
	//添加objectid
	dev.Id = bson.NewObjectId()
	err = c.Insert(dev)
	if err != nil {
		return errors.New("添加失败:" + err.Error())
	}
	return nil
}

//删除设备信息
func DeleteDevice(id string) error {
	c := GetSession().DB(DatabaseName).C(DeviceCollection)
	err := c.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}
	return nil
}

//保存设备参数信息
func UpdateDeviceSetParams(dev *models.DeviceInfo) error {
	c := GetSession().DB(DatabaseName).C(DeviceCollection)
	device := models.DeviceInfo{}
	//查找设备
	err := c.Find(&bson.M{"sensorid": dev.SensorId}).One(&device)
	if err != nil {
		return err
	}
	//更新信息
	device.SetParams = dev.SetParams
	device.UpdateTime = time.Now()
	err = c.Update(&bson.M{"sensorid": dev.SensorId}, &device)
	if err != nil {
		return err
	}
	return nil
}

//保存设备自定义信息
func UpdateDeviceCustomeParams(dev *models.DeviceInfo) error {
	c := GetSession().DB(DatabaseName).C(DeviceCollection)
	device := models.DeviceInfo{}
	//查找设备
	err := c.Find(&bson.M{"sensorid": dev.SensorId}).One(&device)
	if err != nil {
		return err
	}
	//更新信息
	device.CustomParams = dev.CustomParams
	device.UpdateTime = time.Now()
	err = c.Update(&bson.M{"sensorid": dev.SensorId}, &device)
	if err != nil {
		return err
	}
	return nil
}
