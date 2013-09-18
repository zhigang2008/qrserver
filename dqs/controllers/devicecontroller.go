package controllers

import (
	"dqs/models"
	"dqs/quickserver"
	"dqs/util"
	"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
	"time"
)

type DeviceController struct {
	beego.Controller
}

//获取设备列表或者单个设备信息
func (this *DeviceController) Get() {
	sid := this.Ctx.Params[":objectId"]
	//单个设备查询
	if sid != "" {
		this.Data["title"] = "设备详细信息"
		this.Data["author"] = "wangzhigang"
		this.Data["device"] = models.GetDevice(sid)

		this.TplNames = "device.html"

	} else { //列表
		this.Data["title"] = "设备列表"
		this.Data["author"] = "wangzhigang"

		pagination := util.Pagination{}
		page, err := this.GetInt("page")
		if err != nil {
			pagination.CurrentPage = 1
		} else {
			pagination.CurrentPage = int(page)
		}
		pagesize, err2 := this.GetInt("pagesize")
		if err2 != nil {
			pagination.PageSize = 10
		} else {
			pagination.PageSize = int(pagesize)
		}

		//查询参数
		sid := this.GetString("sensorid")
		if sid != "" {
			pagination.AddParams("sensorid", sid)
		}
		sonline := this.GetString("online")
		if sonline != "" {
			online, err := this.GetBool("online")
			if err == nil {
				pagination.AddParams("online", online)
			}
		}

		//执行查询
		err = models.DeviceList(&pagination)
		if err != nil {
			log.Warnf("查询所有设备信息失败:%s", err.Error())
		}
		pagination.Compute()

		this.Data["pagedata"] = pagination
		this.TplNames = "devicelist.html"
	}
}

//添加设备
func (this *DeviceController) Post() {
	answer := JsonAnswer{}

	device := models.DeviceInfo{}
	setParams := models.SensorInfo{}
	customParams := models.CustomDefineParams{}

	this.ParseForm(&setParams)
	this.ParseForm(&customParams)
	err := this.ParseForm(&device)
	if err != nil {
		answer.Ok = false
		answer.Msg = "数据传递失败:" + err.Error()
	} else {
		device.SetParams = setParams
		device.CustomParams = customParams
		device.UpdateTime = time.Now()
		device.RegisterTime = time.Now()

		err = models.AddDevice(&device)
		if err != nil {
			answer.Ok = false
			answer.Msg = "设备添加失败:" + err.Error()
		} else {
			answer.Ok = true
			answer.Msg = "保存成功"
		}
	}

	this.Data["json"] = &answer
	this.ServeJson()
}

//重新获取设备参数
func (this *DeviceController) RefreshDeviceParams() {
	sid := this.Ctx.Params[":id"]
	err := quickserver.CommandRead(sid)
	answer := JsonAnswer{}
	if err != nil {
		answer.Ok = false
		answer.Msg = "读取失败:" + err.Error()

	} else {
		answer.Ok = true
		answer.Msg = "成功"
	}
	this.Data["json"] = &answer
	this.ServeJson()

}

//更新设备参数
func (this *DeviceController) UpdateDeviceParams() {
	sid := this.Ctx.Params[":id"]
	answer := JsonAnswer{}

	//判断设备编号
	if sid != "" {
		params := quickserver.SensorInfo{}
		err := this.ParseForm(&params)
		if err != nil {
			answer.Ok = false
			answer.Msg = "读取失败:" + err.Error()
		} else {
			//发送控制指令
			err = quickserver.CommandSet(sid, quickserver.SensorInfo2RetData(&params))
			if err != nil {
				answer.Ok = false
				answer.Msg = "控制命令执行失败:" + err.Error()
			} else {
				answer.Ok = true
				answer.Msg = "成功"

				//更新数据库数据
				params2 := models.SensorInfo{}
				this.ParseForm(&params2)
				device := models.DeviceInfo{}
				device.SetParams = params2
				device.SensorId = sid

				err = models.UpdateDeviceSetParams(&device)
				if err != nil {
					answer.Ok = false
					answer.Msg = "设备数据已更新,数据库保存未成功.请等待设备上报数据"
				} else {
					answer.Ok = true
					answer.Msg = "成功"
				}
			}
		}
	}
	this.Data["json"] = &answer
	this.ServeJson()

}

//更新自定义参数
func (this *DeviceController) UpdateCustomParams() {
	sid := this.Ctx.Params[":id"]
	answer := JsonAnswer{}

	//判断设备编号
	if sid != "" {
		params := models.CustomDefineParams{}
		err := this.ParseForm(&params)
		if err != nil {
			answer.Ok = false
			answer.Msg = "读取失败:" + err.Error()
		} else {
			//执行保存操作
			device := models.DeviceInfo{}
			device.SensorId = sid
			device.CustomParams = params

			err = models.UpdateDeviceCustomeParams(&device)
			if err != nil {
				answer.Ok = false
				answer.Msg = "数据保存失败:" + err.Error()
			} else {
				answer.Ok = true
				answer.Msg = "成功"
			}
		}
	}
	this.Data["json"] = &answer
	this.ServeJson()

}

//添加设备页面
func (this *DeviceController) ToDeviceAddPage() {
	this.Data["title"] = "添加设备"
	this.Data["author"] = "wangzhigang"

	this.TplNames = "deviceadd.html"
}
