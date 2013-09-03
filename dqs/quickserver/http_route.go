package quickserver

import (
	log "github.com/cihub/seelog"
	"html/template"
	"net/http"
	"strconv"
)

func HtmlType(w *http.ResponseWriter) {
	(*w).Header().Add("Content-Type", "text/html")
}

//展示设备列表
func (hs *HttpServer) DeviceHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/pages/device.html")
	if err != nil {
		log.Warn(err)
		http.NotFound(w, r)
		return
	}
	var cols int = 20
	colsstr := r.FormValue("num")
	if colsstr != "" {
		var e error
		cols, e = strconv.Atoi(colsstr)
		if e != nil {
			cols = 20
		}
	}
	HtmlType(&w)
	devices, err := hs.dataManager.DeviceList(cols)
	if err != nil {
		log.Warnf("获取设备信息失败:%s", err.Error())
		t.Execute(w, nil)
	}
	t.Execute(w, *devices)
}

//处理报警信息列表
func (hs *HttpServer) AlarmHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/pages/alarm.html")
	if err != nil {
		log.Warn(err)
		http.NotFound(w, r)
		return
	}
	var cols int = 50
	colsstr := r.FormValue("num")
	if colsstr != "" {
		var e error
		cols, e = strconv.Atoi(colsstr)
		if e != nil {
			cols = 50
		}
	}
	HtmlType(&w)
	dataList, err := hs.dataManager.AlarmList(cols)
	if err != nil {
		log.Warnf("获取报警信息失败:%s", err.Error())
		t.Execute(w, nil)
	}
	t.Execute(w, *dataList)
}

//处理主页面
func (hs *HttpServer) IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/pages/index.html")
	if err != nil {
		log.Warn(err)
		http.NotFound(w, r)
		return
	}
	HtmlType(&w)
	t.Execute(w, nil)
}

//日志查看页面
func (hs *HttpServer) LogViewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/pages/logs.html")
	if err != nil {
		log.Warn(err)
		http.NotFound(w, r)
		return
	}
	HtmlType(&w)
	t.Execute(w, nil)
}
