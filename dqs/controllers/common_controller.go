package controllers

import (
	"dqs/dao"
	"dqs/models"
	"dqs/util"
	//"fmt"
	log "github.com/cihub/seelog"
	"net/url"
	"strings"
)

type CommonController struct {
	BaseController
}

//用户登录
func (this *CommonController) Sign() {
	sid := this.GetString("SignId")
	pwd := this.GetString("Password")
	toUrl := this.GetString("tourl")
	encodePwd := util.EncodePwd(pwd)

	params := url.Values{}
	user := models.User{}

	if toUrl == "" {
		toUrl = this.Ctx.Request.Referer()
	}
	params.Set("tourl", toUrl)
	//fmt.Printf("Sign in refer:%s \n", toUrl)
	if sid != "" && pwd != "" {
		user = dao.GetSignUser(sid)
	}

	if user.UserId == "" {
		params.Set("msg", "用户不存在")
		this.Redirect("/login?"+params.Encode(), 302)
	} else {
		if user.Password == encodePwd {
			log.Infof("[%s(%s)]成功登录", user.UserId, sid)
			this.SetSession(CURRENTUSER, user)
			this.SetSession("userName", user.UserName)

			if strings.Contains(toUrl, PAGE_LOGIN) {
				this.Redirect(PAGE_INDEX, 302)
			}
			this.Redirect(toUrl, 302)
		} else {
			params.Set("msg", "用户ID/密码不匹配,登录失败")
			log.Infof("[%s]试图登录,失败", sid)
			this.Redirect("/login?"+params.Encode(), 302)
		}
	}

}

//系统登出
func (this *CommonController) SignOut() {
	//url := this.Ctx.Request.Referer()
	this.DelSession(CURRENTUSER)
	this.Redirect("/", 302)
}

//登录页面
func (this *CommonController) Login() {
	this.Data["title"] = "用户登录"
	this.Data["msg"] = this.GetString("msg")
	this.Data["tourl"] = this.GetString("tourl")

	this.TplNames = "login.html"
}
