package dao

import (
	"dqs/models"
	//"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	//"time"
)

//获取系统参数
func GetSystemConfig() (models.SystemConfig, error) {
	c := GetSession().DB(DatabaseName).C(SystemCollection)
	configs := models.SystemConfig{}

	err := c.Find(nil).One(&configs)
	if err != nil {
		configs = models.SystemConfig{}
		return configs, err
	}
	return configs, nil
}

//更新系统参数信息
func UpdateSystemConfig(configs *models.SystemConfig) error {
	c := GetSession().DB(DatabaseName).C(SystemCollection)
	_, err := c.Upsert(nil, configs)
	if err != nil {
		return err
	}
	return nil
}

//添加系统参数信息
func AddSystemConfig(configs *models.SystemConfig) error {
	c := GetSession().DB(DatabaseName).C(SystemCollection)
	err := c.Insert(configs)
	if err != nil {
		return err
	}
	return nil
}
