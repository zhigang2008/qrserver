package httpserver

import (
	"dqs/models"
	"dqs/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

const (
	PAGE_INDEX = "/"
	PAGE_LOGIN = "/login"
)

func configInit() {
	beego.SessionOn = true
	//错误页面处理
	beego.Errorhandler("404", page_not_found)
	beego.Errorhandler("401", page_unauth)
}

//添加模板函数
func addTemplateFuncs() {
	//beego.AddFuncMap("eq", util.Equals)
	beego.AddFuncMap("seqno", util.GenerateSeqNo)
	beego.AddFuncMap("purl", util.GenerateParamUrl)
	beego.AddFuncMap("contain", util.Contain)

}

var checkUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("user").(models.User)
	if ok {
		//ctx.Output.res.
	}
}

func addFilter() {
	beego.AddFilter("*", "AfterExec", checkUser)
}
