package models

import (
	"crypto/md5"
	"encoding/hex"
	"labix.org/v2/mgo/bson"
	"time"
)

const (
	salt = "@Steven@"
)

var h = md5.New()

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
	h.Reset()
	h.Write([]byte(pwd + salt))
	this.Password = hex.EncodeToString(h.Sum(nil))
}

//检查密码是否一致
func (this *User) CheckPwd(pwd string) bool {
	h.Reset()
	h.Write([]byte(pwd + salt))
	ckpwd := hex.EncodeToString(h.Sum(nil))
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
