package quickserver

import (
	log "github.com/cihub/seelog"
	"net/http"
	"strconv"
)

//Http Server 结构
//需要数据库操作.

type HttpServer struct {
	Name        string
	dataManager *DataManager
}

//启动 http Server
func StartHttp(conf ServerConfig) {
	s := HttpServer{}
	var err error
	s.dataManager, err = InitDatabase(conf.Database)
	defer s.dataManager.DataClose()

	if err == nil {
		log.Info("启动HttpSever...")
		http.Handle("/css/", http.FileServer(http.Dir("web")))
		http.Handle("/js/", http.FileServer(http.Dir("web")))

		http.HandleFunc("/device/", s.DeviceHandler)
		http.HandleFunc("/alarm/", s.AlarmHandler)
		http.HandleFunc("/", s.IndexHandler)
		http.ListenAndServe(conf.HttpServer.Host+":"+strconv.Itoa(conf.HttpServer.Port), nil)
	}

}
