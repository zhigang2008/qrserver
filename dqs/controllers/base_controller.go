package controllers

import (
	"dqs/dao"
	"dqs/models"
	"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
	"strings"
	"time"
)

const (
	CURRENTUSER      = "StevenCurrentSessionUser"
	PAGE_INDEX       = "/"
	PAGE_LOGIN       = "/login"
	CommonTimeLayout = "2006-01-02 15:04:05"
	CommonDateLayout = "2006-01-02"
)

var (
	SystemConfigs  models.SystemConfig
	RuntimeConfigs models.DatabaseConfig
)

type BaseController struct {
	beego.Controller
}

//获取用户
func (this *BaseController) CheckUser() {
	if beego.SessionOn {

		u, ok := this.GetSession(CURRENTUSER).(models.User)
		if ok {
			this.Data["isLogin"] = true
			this.Data["CurrentUser"] = &u
			this.Data["CurrentUserName"] = u.UserName
			this.Data["CurrentUserId"] = u.UserId
		} else {
			this.Data["isLogin"] = false
		}
	}

}

//获取Session里的当前用户
func (this *BaseController) GetCurrentUser() models.User {
	if beego.SessionOn {
		u, ok := this.GetSession(CURRENTUSER).(models.User)
		if ok {
			return u
		} else {
			return models.User{}
		}
	}
	return models.User{}

}

//验证权限
func (this *BaseController) AuthRoles(roles ...string) {
	check := false

	if beego.SessionOn {
		u, ok := this.GetSession(CURRENTUSER).(models.User)
		if ok {
			for _, cr := range u.Roles {
				for _, r := range roles {
					if cr == r {
						check = true
						break
					}
				}
			}

		}
	}

	if check == false {
		this.Abort("401")
	}

}

//验证是否登录
func (this *BaseController) Authentication() {
	check := false

	if beego.SessionOn {
		u, ok := this.GetSession(CURRENTUSER).(models.User)
		if ok {
			if u.UserId != "" {
				check = true

			}
		}
	}

	if check == false {
		this.Abort("401")
	}

}

//进行审计记录
func (this *BaseController) AuditLog(cont string, status bool) {
	audit := models.AuditLog{}
	audit.ActTime = time.Now()
	audit.ActType = models.ActOperation
	audit.ActContent = cont
	audit.Status = status
	remote := this.Ctx.Request.RemoteAddr
	pos := strings.Index(remote, ":")
	audit.RemoteAddr = remote
	if pos > 0 {
		audit.RemoteAddr = remote[0:pos]
	}

	if beego.SessionOn {
		u, ok := this.GetSession(CURRENTUSER).(models.User)
		if ok {
			audit.UserId = u.UserId
			audit.UserName = u.UserName
			//增加审计记录
			dao.AddAuditLog(audit)
		}
	}

}

//初始化系统配置参数
func InitSystemConfigs() {

	//初始化系统参数
	newSysCfg, err0 := dao.GetSystemConfig()
	newRuntimeCfg, err1 := dao.GetDBConfig()
	if err0 != nil || err1 != nil {
		//滞后再去读取
		go reInitSysteConfigs()
		return
		/*
			SystemConfigs = models.SystemConfig{}
			SystemConfigs.UserDefaultPassword = "12345678"
			SystemConfigs.UseGis = false
			SystemConfigs.GisServiceUrl = "http://localhost:8080/geoserver/dqs/wms"
			SystemConfigs.GisServiceParams = "?service=WMS&version=1.1.0&request=GetMap"
			SystemConfigs.GisLayerBasic = "dqs_layers"
			SystemConfigs.GisLayerChina = "china_layer"

			giscfg := models.GisImageConfig{}
			giscfg.Format = "image/png"
			giscfg.SRS = "EPSG:4326"
			giscfg.Height = "200"
			giscfg.Weight = "200"
			giscfg.BBOX = "80.0,20.0,120.0,45.0"
			SystemConfigs.GisImageCfg = giscfg
			err2 := dao.AddSystemConfig(&SystemConfigs)
			if err2 != nil {
				log.Warnf("初始化系统参数失败.")
			}
		*/
	} else {
		SystemConfigs = newSysCfg
		RuntimeConfigs = newRuntimeCfg
	}
}

//*滞后时间进行读取
func reInitSysteConfigs() {
	time.Sleep(2 * time.Second)
	newCfg3, err3 := dao.GetSystemConfig()
	newRuntimeCfg4, err4 := dao.GetDBConfig()
	if err3 != nil || err4 != nil {
		log.Criticalf("HttpServer未能读取到系统基本配置.请检查数据库数据")
	} else {
		log.Info("HttpServer读取系统基本配置成功.")
		SystemConfigs = newCfg3
		RuntimeConfigs = newRuntimeCfg4
	}
}
