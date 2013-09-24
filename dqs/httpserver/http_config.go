package httpserver

import (
	"dqs/util"
	"github.com/astaxie/beego"
	"net/http"
)

func configInit() {
	beego.SessionOn = true
}

//添加模板函数
func addTemplateFuncs() {
	//beego.AddFuncMap("eq", util.Equals)
	beego.AddFuncMap("seqno", util.GenerateSeqNo)
	beego.AddFuncMap("purl", util.GenerateParamUrl)
	beego.AddFuncMap("contain", util.Contain)

}

var FilterUser = func(w http.ResponseWriter, r *http.Request) {
	session := beego.GlobalSessions.SessionStart()
}

func addFilter() {
	beego.FilterAfter()
}
