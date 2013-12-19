package controllers

import (
	"dqs/dao"
	"dqs/models"
	"dqs/quickserver"
	log "github.com/cihub/seelog"
)

type RuntimeConfigController struct {
	BaseController
}

//获取数据库配置信息
func (this *RuntimeConfigController) Get() {
	this.Data["title"] = "运行参数"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()
	//权限检查
	this.AuthRoles("role_admin")

	//执行查询
	configs, err := dao.GetDBConfig()
	if err != nil {
		log.Warnf("查询运行参数失败:%s", err.Error())
	}
	this.Data["configs"] = configs
	this.TplNames = "runtimeconfigs.html"
	this.Render()
}

//保存运行参数
func (this *RuntimeConfigController) Put() {
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
		log.Warnf("更新运行参数-解析参数失败")
	} else {
		dbConfigs.EventParams = eventParams
		dbConfigs.FileConfig = fileCfg
		dbConfigs.ReportCfg = reportParams

		err := dao.UpdateDBConfig(&dbConfigs)
		if err != nil {
			answer.Ok = false
			answer.Msg = "更新运行参数失败:" + err.Error()
			log.Warnf("更新运行参数失败:%s", err.Error())
		} else {
			answer.Ok = true
			answer.Msg = "更新运行参数成功"
			log.Infof("更新运行参数成功")
			this.AuditLog("更新运行参数", true)

			//更新服务端配置内容
			reerr := quickserver.CommandRefreshConfig()
			if reerr != nil {
				answer.Ok = false
				answer.Msg = "运行参数更改成功,但服务端刷新失败,请重新刷新运行参数,或重启服务."
				log.Errorf("运行参数已更改,但服务端更新失败,请重新刷新运行参数,或重启服务")
			}
		}
	}

	this.Data["json"] = &answer
	this.ServeJson()
}
