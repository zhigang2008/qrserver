package httpserver

import (
	"dqs/dao"
	"time"
)

func FeeJob() {
	timer := time.NewTicker(time.Hour * 24)
	for {
		select {
		case <-timer.C:

			go doJob()
		}
	}
}

func doJob() {
	devices, err := dao.GetAllDevices()
	if err != nil {
		for _, v := range devices {

		}
	}
}
