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

//更改用户信息
func (this *UserSelfController) Update() {
	//sid := this.GetString(":objectid")

	answer := JsonAnswer{}

	user := models.User{}
	err := this.ParseForm(&user)

	if err != nil {
		answer.Ok = false
		answer.Msg = "数据传递失败:" + err.Error()
	} else {
		curUser := this.GetCurrentUser()
		if curUser.UserId != user.UserId {
			log.Warnf("用户[%s]试图修改用户[%s]的信息,被禁止", curUser.UserId, user.UserId)
			this.Abort("404")
		}

		err = dao.UpdateUserBySelf(&user)
		if err != nil {
			answer.Ok = false
			answer.Msg = "用户更改失败:" + err.Error()
		} else {
			answer.Ok = true
			answer.Msg = "用户更改成功"
		}
	}

	this.Data["json"] = &answer
	this.ServeJson()
}

//重置用户密码
func (this *UserSelfController) ResetPassword() {
	//检查是否登录
	this.Authentication()

	answer := JsonAnswer{}

	uid := this.GetString("UserId")
	oldPwd := this.GetString("oldPwd")
	newPwd := this.GetString("newPwd")

	user := dao.GetUser(uid)

	if user.ObjectId == "" {
		answer.Ok = false
		answer.Msg = "当前用户不存在"
	} else {
		if user.CheckPwd(oldPwd) == false {
			answer.Ok = false
			answer.Msg = "原始密码不正确"
		} else {
			err := dao.ResetUserPassword(uid, newPwd)
			if err != nil {
				answer.Ok = false
				answer.Msg = err.Error()
			} else {
				answer.Ok = true
				answer.Msg = "重置密码成功"
				//audit
				this.AuditLog("用户修改自身密码", true)
			}
		}
	}

	this.Data["json"] = &answer
	this.ServeJson()
}
