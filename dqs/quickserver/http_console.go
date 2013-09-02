package quickserver

import (
	log "github.com/cihub/seelog"
	"net/http"
	"strconv"
)

func StartHttp(conf HttpServerConfig) error {
	log.Info("启动HttpSever...")
	http.HandleFunc("/device/", DeviceHandler)
	http.HandleFunc("/alarm/", AlarmHandler)
	http.HandleFunc("/", NotFoundHandler)
	http.ListenAndServe(conf.Host+":"+strconv.Itoa(conf.Port), nil)
	return nil
}
