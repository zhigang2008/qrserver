package controllers

import (
//"dqs/dao"
//"github.com/astaxie/beego"
//log "github.com/cihub/seelog"
)

type TestController struct {
	BaseController
}

func (this *TestController) TestGis() {
	this.Data["title"] = "测试Gis"
	this.Data["author"] = "wangzhigang"

	this.TplNames = "testgis.html"
	this.Render()
}
func (this *TestController) Test() {
	this.Data["title"] = "测试页面"
	this.Data["author"] = "wangzhigang"

	this.TplNames = "test.html"
	this.Render()
}
