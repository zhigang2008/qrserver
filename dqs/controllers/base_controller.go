package controllers

import (
	"dqs/models"
	"github.com/astaxie/beego"
)

const (
	CURRENTUSER = "StevenCurrentSessionUser"
	PAGE_INDEX  = "/"
	PAGE_LOGIN  = "/login"
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

//验证权限
func (this *BaseController) AuthRoles(roles ...string) {
	check := false

	if beego.SessionOn {
		u, ok := this.GetSession(CURRENTUSER).(models.User)
		if ok {
			for _, cr := range u.Roles {
				for _, r := range roles {
					if cr == r {
						check = true
						break
					}
				}
			}

		}
	}

	if check == false {
		this.Abort("401")
	}

}

//验证是否登录
func (this *BaseController) Authentication() {
	check := false

	if beego.SessionOn {
		u, ok := this.GetSession(CURRENTUSER).(models.User)
		if ok {
			if u.UserId != "" {
				check = true

			}
		}
	}

	if check == false {
		this.Abort("401")
	}

}
