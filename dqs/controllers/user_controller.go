package controllers

import (
	"dqs/dao"
	"dqs/models"
	"dqs/util"
	//	"fmt"
	"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
	"time"
)

type UserController struct {
	BaseController
}

//用户信息列表
func (this *UserController) Get() {
	sid := this.GetString(":objectId")

	this.CheckUser()
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
		user.Roles = this.GetStrings("Roles")
		user.SetPassword(beego.AppConfig.String("user.default.password"))
		user.CreateTime = time.Now()

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

//更改用户信息
func (this *UserController) Put() {
	answer := JsonAnswer{}

	user := models.User{}
	err := this.ParseForm(&user)

	if err != nil {
		answer.Ok = false
		answer.Msg = "数据传递失败:" + err.Error()
	} else {
		user.Roles = this.GetStrings("Roles")
		user.UpdateTime = time.Now()

		err = dao.UpdateUser(&user)
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

//删除用户
func (this *UserController) Delete() {
	answer := JsonAnswer{}
	oid := this.GetString(":objectId")

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

	this.CheckUser()
	this.TplNames = "useradd.html"
}

//添加用户页面
func (this *UserController) ToResetPassword() {
	uid := this.GetString(":uid")
	this.CheckUser()

	user := dao.GetUser(uid)

	this.Data["title"] = "重置用户密码"
	this.Data["author"] = "wangzhigang"
	this.Data["user"] = user

	this.TplNames = "resetpwd.html"
}

//重置用户密码
func (this *UserController) ResetPassword() {
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
			}
		}
	}

	this.Data["json"] = &answer
	this.ServeJson()
}
