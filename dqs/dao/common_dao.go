package dao

import (
	"dqs/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

//用户登录查询
func GetSignUser(sid string) models.User {
	c := GetSession().DB(DatabaseName).C(UserCollection)
	user := models.User{}
	//查找用户
	err := c.Find(&bson.M{"userid": sid}).One(&user)
	if err != mgo.ErrNotFound {
		return user
	}
	err = c.Find(&bson.M{"mobile": sid}).One(&user)
	if err != mgo.ErrNotFound {
		return user
	}
	err = c.Find(&bson.M{"email": sid}).One(&user)
	if err != mgo.ErrNotFound {
		return user
	}
	return user
}
