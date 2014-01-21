package controllers

import (
	"dqs/dao"
	"dqs/models"
	"dqs/util"
	log "github.com/cihub/seelog"
	"time"
)

type FeeController struct {
	BaseController
}

//用户登录
func (this *FeeController) GetPayments() {
	this.Data["title"] = "待续费设备信息"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()

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

	SensorId := this.GetString("SensorId")
	if SensorId != "" {
		pagination.AddParams("SensorId", SensorId)
	}
	NetOperator := this.GetString("NetOperator")
	if NetOperator != "" {
		pagination.AddParams("NetOperator", NetOperator)
	}
	NetNo := this.GetString("NetNo")
	if NetNo != "" {
		pagination.AddParams("NetNo", NetNo)
	}
	begintime := this.GetString("begintime")
	if begintime != "" {
		pagination.AddParams("begintime", begintime)
	}
	endtime := this.GetString("endtime")
	if endtime != "" {
		pagination.AddParams("endtime", endtime)
	}

	//执行查询
	err = dao.GetPagePayments(&pagination)
	if err != nil {
		log.Warnf("查询待充值信息失败:%s", err.Error())
	}
	pagination.Compute()

	this.Data["pagedata"] = pagination
	this.TplNames = "payment.html"
	this.Render()
}

//充值记录
func (this *FeeController) GetPaymentHistory() {
	//权限检查
	this.AuthRoles("role_admin", "role_device")

	this.Data["title"] = "充值记录"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()

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

	SensorId := this.GetString("SensorId")
	if SensorId != "" {
		pagination.AddParams("SensorId", SensorId)
	}
	UserId := this.GetString("UserId")
	if UserId != "" {
		pagination.AddParams("UserId", UserId)
	}
	NetOperator := this.GetString("NetOperator")
	if NetOperator != "" {
		pagination.AddParams("NetOperator", NetOperator)
	}
	NetNo := this.GetString("NetNo")
	if NetNo != "" {
		pagination.AddParams("NetNo", NetNo)
	}
	begintime := this.GetString("begintime")
	if begintime != "" {
		pagination.AddParams("begintime", begintime)
	}
	endtime := this.GetString("endtime")
	if endtime != "" {
		pagination.AddParams("endtime", endtime)
	}

	//执行查询
	err = dao.GetPagePaymentHistorys(&pagination)
	if err != nil {
		log.Warnf("查询充值记录信息失败:%s", err.Error())
	}
	pagination.Compute()

	this.Data["pagedata"] = pagination
	this.TplNames = "paymenthistory.html"
	this.Render()
}

//充值
func (this *FeeController) Recharge() {
	this.CheckUser()
	user := this.GetCurrentUser()

	answer := JsonAnswer{}
	hasErr := false
	hasSuccess := false

	//权限验证
	isAuth := this.IsAuthRoles("role_admin", "role_device")
	if isAuth == false {
		answer.Ok = false
		answer.Msg = "您无此操作权限,请登录或更改其他账户"
		this.Data["json"] = &answer
		this.ServeJson()
	}

	months, err := this.GetInt("feeMonths")
	if err != nil {
		answer.Ok = false
		answer.Msg = "充值月份数据不正确"
		this.Data["json"] = &answer
		this.ServeJson()
	}

	devices := this.GetStrings("feeDeviceids")
	for _, v := range devices {
		device := dao.GetDevice(v)
		if device.SensorId != "" {
			validDate := device.CustomParams.WirelessTypeInfo.ValidDate
			if validDate.IsZero() == false {
				device.CustomParams.WirelessTypeInfo.ValidDate = validDate.AddDate(0, int(months), 0)
				//更新有效时间
				err1 := dao.UpdateDeviceCustomeParams(&device)
				//更新充值记录
				if err1 == nil {

					history := new(models.PaymentHistory)
					history.SensorId = v
					history.UserId = user.UserId
					history.UserName = user.UserName
					history.NetOperator = device.CustomParams.WirelessTypeInfo.NetOperator
					history.NetNo = device.CustomParams.WirelessTypeInfo.NetNo
					history.AddQixian = int(months)
					history.ValidDate = device.CustomParams.WirelessTypeInfo.ValidDate
					history.OperateTime = time.Now()
					//添加记录
					err2 := dao.AddPaymentHistory(history)
					if err2 != nil {
						hasErr = true
					} else {
						hasSuccess = true
					}

				} else {
					hasErr = true
				}
			}
		} else {
			answer.Ok = false
			answer.Msg = "获取设备信息失败,无法完成充值操作"
			this.Data["json"] = &answer
			this.ServeJson()
		}
	}

	if hasSuccess == true && hasErr == false {
		answer.Ok = true
		answer.Msg = "充值成功"
	} else if hasErr == true {
		answer.Ok = true
		answer.Msg = "部分充值操作成功"
	} else {
		answer.Ok = true
		answer.Msg = "成功,未进行任何操作变更"
	}

	//数据清理
	if answer.Ok == true {
		//重新清理数据
		this.reUpdatePaymentData()
		//audit
		this.AuditLog("提交续费操作", true)
	}

	this.Data["json"] = &answer
	this.ServeJson()
}

//充值后刷新
func (this *FeeController) RefreshAfterRecharge() {
	this.Data["title"] = "待充值设备信息"
	this.Data["author"] = "wangzhigang"
	this.CheckUser()

	this.reUpdatePaymentData()

	//查询
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
	//执行查询
	err = dao.GetPagePayments(&pagination)
	if err != nil {
		log.Warnf("查询待充值信息失败:%s", err.Error())
	}
	pagination.Compute()

	this.Data["pagedata"] = pagination
	this.TplNames = "payment.html"
	this.Render()
}

//清理数据
func (this *FeeController) reUpdatePaymentData() {
	//获取所有设备
	devices, err0 := dao.GetAllValidDevices()
	curTime := time.Now()

	if err0 == nil {
		//清除待付费信息
		dao.ClearPaymentInfo()

		for _, v := range devices {
			if v.CustomParams.NetType == "3G" {
				validDate := v.CustomParams.WirelessTypeInfo.ValidDate
				if validDate.IsZero() == false {
					leftHours := validDate.Sub(curTime).Hours()
					if leftHours < 30*24 {
						payment := new(models.DevicePayment)
						payment.SensorId = v.SensorId
						payment.SiteAliasName = v.CustomParams.SiteAliasName
						payment.NetOperator = v.CustomParams.WirelessTypeInfo.NetOperator
						payment.NetNo = v.CustomParams.WirelessTypeInfo.NetNo
						payment.NetTariff = v.CustomParams.WirelessTypeInfo.NetTariff
						payment.NetTraffic = v.CustomParams.WirelessTypeInfo.NetTraffic
						payment.StartDate = v.CustomParams.WirelessTypeInfo.StartDate
						payment.ValidDate = v.CustomParams.WirelessTypeInfo.ValidDate
						payment.LeftDate = int(leftHours / 24)

						dao.AddPayment(payment)

					}
				}
			}
		}
	} else {
		log.Warnf("查询有效设备失败:%s", err0.Error())
	}
}
