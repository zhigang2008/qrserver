package controllers

import (
	"dqs/dao"
	"dqs/models"
	"dqs/util"
	//"fmt"
	log "github.com/cihub/seelog"
	"net/url"
	"strings"
	"time"
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
		//audit
		this.auditLogIn(sid, sid, "试图用不存在的用户["+sid+"]登录", false)
		log.Warnf("试图用不存在的用户[%s]登录", sid)

		this.Redirect("/login?"+params.Encode(), 302)
	} else {
		if user.Password == encodePwd {
			log.Infof("[%s(%s)]成功登录", user.UserName, user.UserId)
			this.SetSession(CURRENTUSER, user)
			this.SetSession("userName", user.UserName)
			//audit
			this.auditLogIn(user.UserId, user.UserName, "登录成功", true)

			if strings.Contains(toUrl, PAGE_LOGIN) {
				this.Redirect(PAGE_INDEX, 302)
			}
			this.Redirect(toUrl, 302)
		} else {
			params.Set("msg", "用户ID/密码不匹配,登录失败")
			log.Warnf("[%s]试图登录,失败", sid)
			//audit
			this.auditLogIn(user.UserId, user.UserName, "用户密码不匹配", false)
			this.Redirect("/login?"+params.Encode(), 302)
		}
	}

}

//系统登出
func (this *CommonController) SignOut() {
	//url := this.Ctx.Request.Referer()
	//audit
	u, ok := this.GetSession(CURRENTUSER).(models.User)
	if ok {
		log.Infof("[%s(%s)]注销", u.UserName, u.UserId)
		this.auditLogOut(u)
	}

	this.DelSession(CURRENTUSER)
	this.Redirect(PAGE_INDEX, 302)
}

//登录页面
func (this *CommonController) Login() {
	this.Data["title"] = "用户登录"
	this.Data["msg"] = this.GetString("msg")
	this.Data["tourl"] = this.GetString("tourl")

	this.TplNames = "login.html"
	this.Render()
}

//进行审计记录
func (this *CommonController) auditLogIn(id, uname, cont string, status bool) {
	audit := models.AuditLog{}
	audit.ActTime = time.Now()
	audit.ActType = models.ActLogin
	audit.ActContent = cont
	audit.Status = status
	remote := this.Ctx.Request.RemoteAddr
	pos := strings.Index(remote, ":")
	audit.RemoteAddr = remote
	if pos > 0 {
		audit.RemoteAddr = remote[0:pos]
	}
	audit.UserId = id
	audit.UserName = uname
	//增加审计记录
	dao.AddAuditLog(audit)
}

//进行审计记录
func (this *CommonController) auditLogOut(u models.User) {
	audit := models.AuditLog{}
	audit.ActTime = time.Now()
	audit.ActType = models.ActLogout

	audit.Status = true
	remote := this.Ctx.Request.RemoteAddr
	pos := strings.Index(remote, ":")
	audit.RemoteAddr = remote
	if pos > 0 {
		audit.RemoteAddr = remote[0:pos]
	}

	audit.UserId = u.UserId
	audit.UserName = u.UserName
	audit.ActContent = "用户注销"
	//增加审计记录
	dao.AddAuditLog(audit)

}
