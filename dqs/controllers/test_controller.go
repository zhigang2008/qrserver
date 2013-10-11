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
