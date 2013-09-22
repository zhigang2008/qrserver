package dao

import (
	"dqs/models"
	"dqs/util"
	"errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

//用户列表
func UserList(p *util.Pagination) error {
	c := GetSession().DB(DatabaseName).C(UserCollection)
	users := []models.User{}
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

	//查找用户
	err = query.Skip(p.SkipNum()).Limit(p.PageSize).All(&users)
	if err != nil {
		return err
	}
	p.Data = users
	return nil
}

//根据编号查找用户
func GetUser(sid string) models.User {
	c := GetSession().DB(DatabaseName).C(UserCollection)
	user := models.User{}
	//查找用户
	err := c.Find(&bson.M{"userid": sid}).One(&user)
	if err != nil {
		user = models.User{}
	}
	return user
}

//添加用户
func AddUser(u *models.User) error {
	c := GetSession().DB(DatabaseName).C(UserCollection)
	//先查找,是否存在
	user := models.User{}
	err := c.Find(&bson.M{"userid": u.UserId}).One(&user)
	if err != nil && err != mgo.ErrNotFound {
		return err
	}
	if user.UserId != "" {
		return errors.New("该用户名已被使用")
	}

	//添加objectid
	u.ObjectId = bson.NewObjectId()
	u.CreateTime = time.Now()

	err = c.Insert(u)
	if err != nil {
		return errors.New("添加失败:" + err.Error())
	}
	return nil
}

//删除设备信息
func DeleteUser(objectid string) error {
	c := GetSession().DB(DatabaseName).C(UserCollection)
	err := c.RemoveId(bson.ObjectIdHex(objectid))
	if err != nil {
		return err
	}
	return nil
}

//保存设备参数信息
func UpdateUser(user *models.User) error {
	c := GetSession().DB(DatabaseName).C(UserCollection)
	err := c.Update(&bson.M{"userid": user.UserId}, user)
	if err != nil {
		return err
	}
	return nil
}
