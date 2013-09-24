package httpserver

import (
	"dqs/dao"
	"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
)

//Http Server 结构
//需要数据库操作.

type HttpServer struct {
	Name string
}

//启动 http Server
func StartHttp() {
	s := HttpServer{}
	s.Name = "Http Server"

	host := beego.AppConfig.String("database.host")
	dataBaseName := beego.AppConfig.String("database.dbname")
	dataCollection := beego.AppConfig.String("database.datacollection")
	deviceCollection := beego.AppConfig.String("database.devicecollection")
	userCollection := beego.AppConfig.String("database.usercollection")
	port, err := beego.AppConfig.Int("database.port")

	if err != nil {
		log.Warnf("Http Server 的配置的数据库端口参数应是整型格式.")
		return
	}

	//创建数据库连接
	err = dao.Init(host, port, dataBaseName, dataCollection, deviceCollection, userCollection)
	if err != nil {
		log.Warnf("Http Server 数据库连接不能创建:%s", err.Error())
		return
	}

	log.Info("启动 Http Server...")

	//初始化配置
	configInit()
	//配置自定义模板方法
	addTemplateFuncs()

	//配置路由信息
	routeConfig()
	//启动Beego服务
	beego.Run()
}

//关闭http的数据库连接
func (s *HttpServer) Close() {
	dao.Close()
	beego.CloseSelf()
}
