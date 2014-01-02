package controllers

import (
	"dqs/dao"
	"dqs/models"
	"dqs/quickserver"
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
	gisImagecfg := models.GisImageConfig{}
	mms := models.MMSConfig{}
	mailcfg:=models.MailConfig{}

	err0 := this.ParseForm(&sysConfigs)
	err1 := this.ParseForm(&gisImagecfg)
	err2 := this.ParseForm(&mms)
	err3 := this.ParseForm(&mailcfg)

	if err0 != nil || err1 != nil || err2 != nil || err3 != nil{
		answer.Ok = false
		answer.Msg = "数据传递失败"
		log.Warnf("更新系统配置-解析参数失败")
	} else {
		sysConfigs.GisImageCfg = gisImagecfg
		sysConfigs.MmsCfg = mms
		sysConfigs.MailCfg=mailcfg

		err := dao.UpdateSystemConfig(&sysConfigs)
		if err != nil {
			answer.Ok = false
			answer.Msg = "更新系统配置失败:" + err.Error()
			log.Warnf("更新系统配置失败:%s", err.Error())
		} else {
			//刷新缓存
			SystemConfigs = sysConfigs
			//更新服务端配置内容
			reerr := quickserver.CommandRefreshSystemConfig()
			if reerr != nil {
				answer.Ok = false
				answer.Msg = "系统配置更改成功,但服务端刷新失败,请重新刷新系统配置,或重启服务."
				log.Errorf("系统配置已更改,但服务端更新失败,请重新刷新系统配置,或重启服务")
			} else {
				answer.Ok = true
				answer.Msg = "更新系统配置成功"
				log.Infof("更新系统配置成功")
				this.AuditLog("更新系统配置", true)
			}
		}
	}

	this.Data["json"] = &answer
	this.ServeJson()
}
