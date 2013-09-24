package httpserver

import (
	"dqs/models"
	"dqs/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
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

var checkUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("user").(models.User)
	if ok {
		//ctx.Output.res.
	}
}

func addFilter() {
	beego.AddFilter("*", "AfterExec", checkUser)
}
