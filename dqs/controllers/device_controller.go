package controllers

import (
	"dqs/dao"
	"dqs/models"
	"dqs/quickserver"
	"dqs/util"
	"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
	"time"
)

type DeviceController struct {
	BaseController
}

//获取设备列表或者单个设备信息
func (this *DeviceController) Get() {
	sid := this.GetString(":objectId")
	this.CheckUser()

	//单个设备查询
	if sid != "" {
		this.Data["title"] = "设备详细信息"
		this.Data["author"] = "wangzhigang"
		this.Data["device"] = dao.GetDevice(sid)

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
		err = dao.DeviceList(&pagination)
		if err != nil {
			log.Warnf("查询所有设备信息失败:%s", err.Error())
		}
		pagination.Compute()

		this.Data["pagedata"] = pagination
		this.TplNames = "devicelist.html"
	}
	this.Render()
}

//添加设备
func (this *DeviceController) Post() {

	//权限检查
	this.AuthRoles("role_admin", "role_device")

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
		log.Warnf("添加设备-解析参数失败:%s", err.Error())
	} else {
		device.SetParams = setParams
		device.CustomParams = customParams
		device.UpdateTime = time.Now()
		device.RegisterTime = time.Now()

		err = dao.AddDevice(&device)
		if err != nil {
			answer.Ok = false
			answer.Msg = "设备添加失败:" + err.Error()
			log.Warnf("添加设备失败[%s]:%s", device.SensorId, err.Error())
		} else {
			answer.Ok = true
			answer.Msg = "保存成功"

			log.Infof("添加设备成功[%s]", device.SensorId)
			this.AuditLog("添加设备["+device.SensorId+"]", true)
		}
	}

	this.Data["json"] = &answer
	this.ServeJson()
}

//删除设备
func (this *DeviceController) Delete() {
	//权限检查
	this.AuthRoles("role_admin", "role_device")

	answer := JsonAnswer{}
	sid := this.GetString(":objectId")

	if sid != "" {
		err := dao.DeleteDevice(sid)
		if err != nil {
			answer.Ok = false
			answer.Msg = "设备删除失败:" + err.Error()
			log.Warnf("删除设备失败[%s]:%s", sid, err.Error())
		} else {
			answer.Ok = true
			answer.Msg = "删除成功"

			log.Infof("删除设备成功[%s]", sid)
			this.AuditLog("删除设备["+sid+"]", true)
		}
	} else {
		answer.Ok = false
		answer.Msg = "没有设备需要删除"
	}

	this.Data["json"] = &answer
	this.ServeJson()
}

//重新获取设备参数
func (this *DeviceController) RefreshDeviceParams() {
	sid := this.GetString("id")
	remote := this.GetString("remote")

	err := quickserver.CommandRead(sid, remote)
	answer := JsonAnswer{}
	if err != nil {
		answer.Ok = false
		answer.Msg = "读取失败:" + err.Error()
		log.Warnf("读取设备参数失败[%s]:%s", sid, err.Error())

	} else {
		answer.Ok = true
		answer.Msg = "成功"
		log.Infof("获取设备参数成功[%s]", sid)
	}
	this.Data["json"] = &answer
	this.ServeJson()

}

//更新设备参数
func (this *DeviceController) UpdateDeviceParams() {
	//权限检查
	this.AuthRoles("role_admin", "role_device")

	sid := this.GetString("id")
	remote := this.GetString("remote")
	answer := JsonAnswer{}

	//判断设备编号
	if sid != "" {
		params := quickserver.SensorInfo{}
		err := this.ParseForm(&params)
		if err != nil {
			answer.Ok = false
			answer.Msg = "读取失败:" + err.Error()
			log.Warnf("更新设备参数-读取参数失败[%s]:%s", sid, err.Error())
		} else {
			//发送控制指令
			err = quickserver.CommandSet(sid, remote, quickserver.SensorInfo2RetData(&params))
			if err != nil {
				answer.Ok = false
				answer.Msg = "控制命令执行失败:" + err.Error()
				log.Warnf("更新设备参数-发送控制命令失败[%s]:%s", sid, err.Error())
			} else {
				answer.Ok = true
				answer.Msg = "成功"

				//更新数据库数据
				params2 := models.SensorInfo{}
				this.ParseForm(&params2)
				device := models.DeviceInfo{}
				device.SetParams = params2
				device.SensorId = sid

				err = dao.UpdateDeviceSetParams(&device)
				if err != nil {
					answer.Ok = false
					answer.Msg = "设备数据已更新,数据库保存未成功.请等待设备上报数据"
					log.Warnf("更新设备参数成功-数据库保存失败[%s]:%s", sid, err.Error())
				} else {
					answer.Ok = true
					answer.Msg = "成功"
				}
				log.Infof("更新设备参数成功[%s]", sid)
				//audit
				this.AuditLog("更新设备参数["+device.SensorId+"]", true)
			}
		}
	}
	this.Data["json"] = &answer
	this.ServeJson()

}

//更新自定义参数
func (this *DeviceController) UpdateCustomParams() {
	//权限检查
	this.AuthRoles("role_admin", "role_device")

	sid := this.GetString(":id")
	answer := JsonAnswer{}

	//判断设备编号
	if sid != "" {
		params := models.CustomDefineParams{}
		wirelessParams := models.WirelessInfo{}
		wireParams := models.WireInfo{}

		this.ParseForm(&wirelessParams)
		this.ParseForm(&wireParams)
		err := this.ParseForm(&params)
		if err != nil {
			answer.Ok = false
			answer.Msg = "读取失败:" + err.Error()
			log.Warnf("更新设备额外参数-参数解析失败[%s]:%s", sid, err.Error())
		} else {
			//执行保存操作
			device := models.DeviceInfo{}
			device.SensorId = sid
			params.WirelessTypeInfo = wirelessParams
			params.WireTypeInfo = wireParams

			device.CustomParams = params

			err = dao.UpdateDeviceCustomeParams(&device)
			if err != nil {
				answer.Ok = false
				answer.Msg = "数据保存失败:" + err.Error()
				log.Warnf("更新设备额外参数-保存失败[%s]:%s", sid, err.Error())
			} else {
				answer.Ok = true
				answer.Msg = "成功"

				log.Infof("更新设备自定义参数成功[%s]", sid)
				//audit
				this.AuditLog("更新设备自定义参数["+device.SensorId+"]", true)
			}
		}
	}
	this.Data["json"] = &answer
	this.ServeJson()

}

//添加设备页面
func (this *DeviceController) ToDeviceAddPage() {
	//权限检查
	this.AuthRoles("role_admin", "role_device")

	this.Data["title"] = "添加设备"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()
	this.TplNames = "deviceadd.html"
	this.Render()
}

//设备定位
func (this *DeviceController) DeviceLocation() {
	sid := this.GetString(":id")
	this.Data["title"] = "设备定位"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()
	device := models.DeviceInfo{}
	//单个设备查询
	if sid != "" {
		device = dao.GetDevice(sid)
	}
	this.Data["device"] = device

	usegis, err := beego.AppConfig.Bool("map.gis")
	if err != nil {
		usegis = false
		log.Warnf("无法从配置文件中获取gis启用信息.将使用地图模式.")
	}
	if usegis {
		this.TplNames = "location-gis.html"
	} else {
		this.TplNames = "location.html"
	}
	this.Render()
}
