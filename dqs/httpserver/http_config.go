package httpserver

import (
	//	"dqs/models"
	"dqs/util"
	//	"fmt"
	"github.com/astaxie/beego"
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
	beego.AddFuncMap("hasrole", util.HasRoles)

}

/*
var authMap map[string]string

func init() {
	authMap = make(map[string]string)
	authMap["post:/device"] = "role_admin"
	authMap["put:/device"] = "role_admin"
	authMap["delete:/device"] = "role_admin"
	authMap["get:/device/add"] = "role_admin"
	authMap["post:/device/add"] = "role_admin"
}

var checkAuth = func(ctx *context.Context) {

	fmt.Println(ctx.Input.Uri())
	fmt.Println(ctx.Input.Url())
	fmt.Println(ctx.Input.Method())

	u, ok := ctx.Input.Session("StevenCurrentSessionUser").(models.User)
	if ok {
		fmt.Printf("user roles:%s.\n", u.Roles)

	} else {
		fmt.Println("no user.")

	}
}

func addFilter() {
	beego.AddFilter("*", "BeforRouter", checkAuth)
}
*/
