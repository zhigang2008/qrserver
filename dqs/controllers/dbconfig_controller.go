package controllers

import (
	"dqs/dao"
	"dqs/models"
	"dqs/quickserver"
	log "github.com/cihub/seelog"
)

type DBConfigController struct {
	BaseController
}

//获取数据库配置信息
func (this *DBConfigController) Get() {
	this.Data["title"] = "配置信息"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()
	//权限检查
	this.AuthRoles("role_admin")

	//执行查询
	configs, err := dao.GetDBConfig()
	if err != nil {
		log.Warnf("查询配置信息失败:%s", err.Error())
	}
	this.Data["configs"] = configs
	this.TplNames = "dbconfigs.html"
	this.Render()
}

//保存配置信息
func (this *DBConfigController) Put() {
	//权限检查
	this.AuthRoles("role_admin", "role_config")

	answer := JsonAnswer{}

	dbConfigs := models.DatabaseConfig{}
	eventParams := models.EventParameters{}
	fileCfg := models.FilesConfig{}
	reportParams := models.ReportParameter{}

	err0 := this.ParseForm(&dbConfigs)
	err1 := this.ParseForm(&eventParams)
	err2 := this.ParseForm(&fileCfg)
	err3 := this.ParseForm(&reportParams)

	if err0 != nil || err1 != nil || err2 != nil || err3 != nil {
		answer.Ok = false
		answer.Msg = "数据传递失败"
		log.Warnf("更新配置信息-解析参数失败")
	} else {
		dbConfigs.EventParams = eventParams
		dbConfigs.FileConfig = fileCfg
		dbConfigs.ReportCfg = reportParams

		err := dao.UpdateDBConfig(&dbConfigs)
		if err != nil {
			answer.Ok = false
			answer.Msg = "更新配置信息失败:" + err.Error()
			log.Warnf("更新配置信息失败:%s", err.Error())
		} else {
			answer.Ok = true
			answer.Msg = "更新配置信息成功"
			log.Infof("更新配置信息成功")
			this.AuditLog("更新配置信息", true)

			//更新服务端配置内容
			reerr := quickserver.CommandRefreshConfig()
			if reerr != nil {
				answer.Ok = false
				answer.Msg = "系统参数更改成功,但服务端刷新失败,请重新刷新服务端参数,或重启服务."
				log.Errorf("系统参数已更改,但服务端更新失败,请重新刷新服务端参数,或重启服务")
			}
		}
	}

	this.Data["json"] = &answer
	this.ServeJson()
}
