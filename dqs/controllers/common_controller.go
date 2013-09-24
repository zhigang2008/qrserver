package controllers

import (
	"dqs/dao"
	"dqs/models"
	"dqs/util"
	//	"fmt"
	"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
)

type CommonController struct {
	beego.Controller
}

//用户登录
func (this *CommonController) Sign() {
	sid := this.GetString("SignId")
	pwd := this.GetString("Password")

	answer := JsonAnswer{}
	encodePwd := util.EncodePwd(pwd)

	user := models.User{}

	if sid != "" && pwd != "" {
		user = dao.GetSignUser(sid)
	}

	if user.UserId == "" {
		answer.Ok = false
		answer.Msg = "用户不存在"
	} else {
		if user.Password == encodePwd {
			answer.Ok = true
			answer.Msg = "登录成功"
			log.Infof("[%s(%s)]成功登录", user.UserId, sid)
			this.SetSession("user", user)
			this.SetSession("userName", user.UserName)
		} else {
			answer.Ok = false
			answer.Msg = "用户ID/密码不匹配,登录失败"
			log.Infof("[%s]试图登录,失败", sid)
		}
	}

	this.Data["json"] = &answer
	this.ServeJson()

}
