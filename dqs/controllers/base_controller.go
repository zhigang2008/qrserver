package controllers

import (
	"dqs/models"
	"github.com/astaxie/beego"
)

const (
	CURRENTUSER = "StevenCurrentSessionUser"
)

type BaseController struct {
	beego.Controller
}

//获取用户
func (this *BaseController) CheckUser() {
	if beego.SessionOn {

		u, ok := this.GetSession(CURRENTUSER).(models.User)
		if ok {
			this.Data["isLogin"] = true
			this.Data["CurrentUser"] = &u
			this.Data["CurrentUserName"] = u.UserName
			this.Data["CurrentUserId"] = u.UserId
		} else {
			this.Data["isLogin"] = false
		}
	}

}

//获取Session里的当前用户
func (this *BaseController) GetCurrentUser() models.User {
	if beego.SessionOn {
		u, ok := this.GetSession(CURRENTUSER).(models.User)
		if ok {
			return u
		} else {
			return models.User{}
		}
	}
	return models.User{}

}
