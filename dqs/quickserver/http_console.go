package quickserver

import (
	log "github.com/cihub/seelog"
	//"net/http"
	//"strconv"
	"dqs/controllers"
	"github.com/astaxie/beego"
)

//Http Server 结构
//需要数据库操作.

type HttpServer struct {
	Name        string
	dataManager *DataManager
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
	var err error
	dbConfig := DataServerConfig{
		Host:             beego.AppConfig.String("database.host"),
		DataBaseName:     beego.AppConfig.String("database.dbname"),
		DataCollection:   beego.AppConfig.String("database.datacollection"),
		DeviceCollection: beego.AppConfig.String("database.devicecollection"),
	}
	port, err2 := beego.AppConfig.Int("database.port")

	if err2 != nil {
		log.Warnf("Http Server 的配置的数据库端口参数应是整型格式.")
		return
	}
	dbConfig.Port = port

	s.dataManager, err = InitDatabase(dbConfig)
	if err != nil {
		log.Warnf("Http Server 数据库连接不能创建:%s", err.Error())
		return
	}

	log.Info("启动 Http Server...")
	beego.Router("/", &controllers.MainController{})
	beego.SetStaticPath("/logs", "logs")
	beego.
		beego.Run()
}

//关闭http的数据库连接
func (s *HttpServer) Close() {
	s.dataManager.DataClose()
	beego.CloseSelf()
}
