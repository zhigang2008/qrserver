package controllers

import (
	//"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"sort"
)

const (
	LogDir = "./logs"
)

type LogsController struct {
	BaseController
}

//用来排序的结构体及方法
type Files []os.FileInfo
type ByName struct{ Files }

func (f Files) Len() int            { return len(f) }
func (f Files) Swap(i, j int)       { f[i], f[j] = f[j], f[i] }
func (s ByName) Less(i, j int) bool { return s.Files[i].Name() > s.Files[j].Name() }

//获取日志文件
func (this *LogsController) Get() {
	//权限检查
	this.AuthRoles("role_admin")

	this.Data["title"] = "运行日志"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()

	var num int = 20

	num64, err0 := this.GetInt("num")
	if err0 == nil {
		num = int(num64)
	}

	logfiles, err := ioutil.ReadDir(LogDir)
	if err != nil {
		this.Data["errormsg"] = "读取日志文件列表失败"
	}
	//按文件名进行排序
	sort.Sort(ByName{logfiles})
	if len(logfiles) >= num {
		this.Data["LogFiles"] = logfiles[:num]
	} else {
		this.Data["LogFiles"] = logfiles
	}

	this.TplNames = "logs.html"
}
