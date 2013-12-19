package controllers

import (
	"dqs/dao"
	"dqs/models"
	log "github.com/cihub/seelog"
)

type SysConfigController struct {
	BaseController
}

//获取数据库配置信息
func (this *SysConfigController) Get() {
	this.Data["title"] = "系统配置"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()
	//权限检查
	this.AuthRoles("role_admin")

	//执行查询
	configs, err := dao.GetSystemConfig()
	if err != nil {
		log.Warnf("查询系统配置失败:%s", err.Error())
	}
	this.Data["configs"] = configs
	this.TplNames = "systemconfigs.html"
	this.Render()
}

//保存系统配置
func (this *SysConfigController) Put() {
	//权限检查
	this.AuthRoles("role_admin")

	answer := JsonAnswer{}
	sysConfigs := models.SystemConfig{}

	err0 := this.ParseForm(&sysConfigs)

	if err0 != nil {
		answer.Ok = false
		answer.Msg = "数据传递失败"
		log.Warnf("更新系统配置-解析参数失败")
	} else {
		err := dao.UpdateSystemConfig(&sysConfigs)
		if err != nil {
			answer.Ok = false
			answer.Msg = "更新系统配置失败:" + err.Error()
			log.Warnf("更新系统配置失败:%s", err.Error())
		} else {
			answer.Ok = true
			answer.Msg = "更新系统配置成功"
			log.Infof("更新系统配置成功")
			this.AuditLog("更新系统配置", true)
			//刷新缓存
			SystemConfigs = sysConfigs
		}
	}

	this.Data["json"] = &answer
	this.ServeJson()
}
