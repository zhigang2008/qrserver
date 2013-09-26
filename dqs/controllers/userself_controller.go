package controllers

import (
	"dqs/dao"
	"dqs/models"
	//"fmt"
	//"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
)

type UserSelfController struct {
	BaseController
}

//用户信息列表
func (this *UserSelfController) View() {
	sid := this.GetString(":objectid")

	this.CheckUser()
	this.Data["title"] = "个人信息"

	u := models.User{}
	if sid != "" {
		u = dao.GetUserByObjectId(sid)
	}
	sesUser := this.GetCurrentUser()

	if sesUser.UserId == u.UserId {
		this.Data["user"] = u
		this.TplNames = "userself.html"
	} else {
		log.Warnf("用户[%s]试图修改[%s]信息,被禁止", sesUser.UserId, u.UserId)
		this.Redirect("/login", 302)
	}
}
