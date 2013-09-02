package quickserver

import (
	log "github.com/cihub/seelog"
	"html/template"
	"net/http"
)

func DeviceHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/html/device.html")
	if err != nil {
		log.Warn(err)
	}
	t.Execute(w, nil)
}

func AlarmHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/html/alarm.html")
	if err != nil {
		log.Warn(err)
	}
	t.Execute(w, nil)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/html/index.html")
	if err != nil {
		log.Warn(err)
	}
	t.Execute(w, nil)
}
