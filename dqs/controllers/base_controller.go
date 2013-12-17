package controllers

import (
	"dqs/dao"
	"dqs/models"
	"github.com/astaxie/beego"
	"strings"
	"time"
)

const (
	CURRENTUSER      = "StevenCurrentSessionUser"
	PAGE_INDEX       = "/"
	PAGE_LOGIN       = "/login"
	CommonTimeLayout = "2006-01-02 15:04:05"
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

//进行审计记录
func (this *BaseController) AuditLog(cont string, status bool) {
	audit := models.AuditLog{}
	audit.ActTime = time.Now()
	audit.ActType = models.ActOperation
	audit.ActContent = cont
	audit.Status = status
	remote := this.Ctx.Request.RemoteAddr
	pos := strings.Index(remote, ":")
	audit.RemoteAddr = remote
	if pos > 0 {
		audit.RemoteAddr = remote[0:pos]
	}

	if beego.SessionOn {
		u, ok := this.GetSession(CURRENTUSER).(models.User)
		if ok {
			audit.UserId = u.UserId
			audit.UserName = u.UserName
			//增加审计记录
			dao.AddAuditLog(audit)
		}
	}

}
