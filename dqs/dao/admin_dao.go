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
	query := c.Find(&m).Sort("-updatetime")
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

//根据Objectid
func GetUserByObjectId(oid string) models.User {
	c := GetSession().DB(DatabaseName).C(UserCollection)
	user := models.User{}
	//查找用户
	err := c.Find(&bson.M{"_id": bson.ObjectIdHex(oid)}).One(&user)
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
	u.UpdateTime = time.Now()

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

	ouser := models.User{}
	//查找用户
	err := c.Find(&bson.M{"userid": user.UserId}).One(&ouser)
	if err != nil {
		return err
	}
	//更改内容o
	ouser.UserName = user.UserName
	ouser.NickName = user.NickName
	ouser.Email = user.Email
	ouser.Gender = user.Gender
	ouser.Phone = user.Phone
	ouser.Mobile = user.Mobile
	ouser.Addr = user.Addr
	ouser.UserTitle = user.UserTitle
	ouser.Blocked = user.Blocked
	ouser.Roles = user.Roles
	ouser.UpdateTime = user.UpdateTime

	err = c.Update(&bson.M{"_id": ouser.ObjectId}, ouser)
	if err != nil {
		return err
	}
	return nil
}
