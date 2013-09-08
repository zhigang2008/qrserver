package quickserver

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

/*/启动 http Server
func StartHttp(conf ServerConfig) {
	s := HttpServer{}
	var err error
	s.dataManager, err = InitDatabase(conf.Database)
	defer s.dataManager.DataClose()

	if err == nil {
		log.Info("启动HttpSever...")
		http.Handle("/css/", http.FileServer(http.Dir("web")))
		http.Handle("/js/", http.FileServer(http.Dir("web")))
		//日志查看
		http.Handle("/logs/", http.StripPrefix("/logs/", http.FileServer(http.Dir("./logs"))))

		http.HandleFunc("/device/", s.DeviceHandler)
		http.HandleFunc("/alarm/", s.AlarmHandler)
		//http.HandleFunc("/log/", s.LogViewHandler)
		http.HandleFunc("/", s.IndexHandler)

		http.ListenAndServe(conf.HttpServer.Host+":"+strconv.Itoa(conf.HttpServer.Port), nil)

	}

}
*/

//启动 http Server
func StartHttp() {
	s := HttpServer{}
	s.Name = "Http Server"

	host := beego.AppConfig.String("database.host")
	dataBaseName := beego.AppConfig.String("database.dbname")
	dataCollection := beego.AppConfig.String("database.datacollection")
	deviceCollection := beego.AppConfig.String("database.devicecollection")
	port, err := beego.AppConfig.Int("database.port")

	if err != nil {
		log.Warnf("Http Server 的配置的数据库端口参数应是整型格式.")
		return
	}

	//创建数据库连接
	err = dao.Init(host, port, dataBaseName, dataCollection, deviceCollection)
	if err != nil {
		log.Warnf("Http Server 数据库连接不能创建:%s", err.Error())
		return
	}

	log.Info("启动 Http Server...")

	//配置路由信息
	RouteConfig()

	//启动Beego服务
	beego.Run()
}

//关闭http的数据库连接
func (s *HttpServer) Close() {
	dao.Close()
	beego.CloseSelf()
}
