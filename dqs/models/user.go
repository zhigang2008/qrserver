package models

import (
	"dqs/util"
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	ObjectId   bson.ObjectId "_id"
	UserId     string
	UserName   string
	Password   string
	NickName   string
	Email      string
	Gender     string
	Phone      string
	Mobile     string
	Addr       string
	UserTitle  string
	Blocked    bool
	CreateTime time.Time
	UpdateTime time.Time
	Roles      []string
}

//设置密码
func (this *User) SetPassword(pwd string) {
	this.Password = util.EncodePwd(pwd)
}

//检查密码是否一致
func (this *User) CheckPwd(pwd string) bool {
	ckpwd := util.EncodePwd(pwd)
	if ckpwd == this.Password {
		return true
	} else {
		return false
	}
}

//角色判断
func (this *User) HasRole(role string) bool {
	for _, v := range this.Roles {
		if v == role {
			return true
		}
	}
	return false
}
