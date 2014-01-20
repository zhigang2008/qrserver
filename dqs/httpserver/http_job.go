package httpserver

import (
	"dqs/dao"
	"dqs/models"
	"fmt"
	"time"
)

func FeeJob() {
	time.Sleep(time.Minute)
	timer := time.NewTicker(time.Hour * 24)
	for {
		select {
		case <-timer.C:
			fmt.Println("job begin")
			go doJob()

		}
	}
}

func doJob() {
	devices, err := dao.GetAllValidDevices()
	curTime := time.Now()

	if err == nil {
		//清除待付费信息
		dao.ClearPaymentInfo()

		for _, v := range devices {
			validDate := v.CustomParams.WirelessTypeInfo.ValidDate
			if validDate.IsZero() == false {
				leftHours := validDate.Sub(curTime).Hours()
				if leftHours < 30*24 {
					payment := new(models.DeviceFee)
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
	} else {
		fmt.Println(err.Error())
	}
}
