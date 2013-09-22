package controllers

import (
	"dqs/dao"
	"dqs/models"
	"dqs/util"
	"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
	"time"
)

type UserController struct {
	beego.Controller
}

//报警信息列表
func (this *UserController) Get() {
	sid := this.Ctx.Params[":objectId"]
	//单个用户查询
	if sid != "" {
		this.Data["title"] = "用户详细信息"
		this.Data["author"] = "wangzhigang"
		this.Data["user"] = dao.GetUser(sid)

		this.TplNames = "user.html"

	} else { //列表
		this.Data["title"] = "用户列表"
		this.Data["author"] = "wangzhigang"

		pagination := util.Pagination{}
		page, err := this.GetInt("page")
		if err != nil {
			pagination.CurrentPage = 1
		} else {
			pagination.CurrentPage = int(page)
		}
		pagesize, err2 := this.GetInt("pagesize")
		if err2 != nil {
			pagination.PageSize = 10
		} else {
			pagination.PageSize = int(pagesize)
		}

		//查询参数
		sid := this.GetString("id")
		if sid != "" {
			pagination.AddParams("id", sid)
		}

		//执行查询
		err = dao.UserList(&pagination)
		if err != nil {
			log.Warnf("查询所有用户信息失败:%s", err.Error())
		}
		pagination.Compute()

		this.Data["pagedata"] = pagination
		this.TplNames = "userlist.html"
	}
}

//添加用户
func (this *UserController) Post() {
	answer := JsonAnswer{}

	user := models.User{}
	err := this.ParseForm(&user)

	if err != nil {
		answer.Ok = false
		answer.Msg = "数据传递失败:" + err.Error()
	} else {
		user.CreateTime = time.Now()
		user.SetPassword(user.Password)

		err = dao.AddUser(&user)
		if err != nil {
			answer.Ok = false
			answer.Msg = "用户添加失败:" + err.Error()
		} else {
			answer.Ok = true
			answer.Msg = "保存成功"
		}
	}

	this.Data["json"] = &answer
	this.ServeJson()
}

//删除用户
func (this *UserController) Delete() {
	answer := JsonAnswer{}
	oid := this.Ctx.Params[":objectId"]

	if oid != "" {
		err := dao.DeleteUser(oid)
		if err != nil {
			answer.Ok = false
			answer.Msg = "用户删除失败:" + err.Error()
		} else {
			answer.Ok = true
			answer.Msg = "删除成功"
		}
	} else {
		answer.Ok = false
		answer.Msg = "没有用户需要删除"
	}

	this.Data["json"] = &answer
	this.ServeJson()
}

//添加用户页面
func (this *UserController) ToUserAddPage() {
	this.Data["title"] = "添加用户"
	this.Data["author"] = "wangzhigang"

	this.TplNames = "useradd.html"
}
