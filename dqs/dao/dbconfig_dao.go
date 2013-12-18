package dao

import (
	"dqs/models"
	//"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	//"time"
)

//获取数据库配置信息
func GetDBConfig() (models.DatabaseConfig, error) {
	c := GetSession().DB(DatabaseName).C(ConfigCollection)
	configs := models.DatabaseConfig{}

	err := c.Find(nil).One(&configs)
	if err != nil {
		configs = models.DatabaseConfig{}
		return configs, err
	}
	return configs, nil
}

//更新配置信息
func UpdateDBConfig(configs *models.DatabaseConfig) error {
	c := GetSession().DB(DatabaseName).C(ConfigCollection)
	_, err := c.Upsert(nil, configs)
	if err != nil {
		return err
	}
	return nil
}
