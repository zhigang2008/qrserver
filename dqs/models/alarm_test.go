package models

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"
)

const (
	CommonTimeLayout = "2006-01-02 15:04:05"
)

func TestXml(t *testing.T) {
	a := AlarmInfo{}
	a.CreateTime = time.Now()
	a.Direction = 1
	a.EventId = "100003424"
	a.HasWaveInfo = false
	a.InitRealTime = time.Now()
	a.InitTime = time.Now().Format(CommonTimeLayout)
	a.Intensity = 9
	a.Latitude = 30.8872
	a.Longitude = 103.2344
	a.Length = 102
	a.ObserveObject = 1
	a.Period = 0.2
	a.PGA = 233.2
	a.SI = 34.2
	a.RegionCode = "100201"
	a.SensorId = "haidian001"
	a.SiteType = 2

	b := a
	b.SensorId = "sichuan002"
	b.Latitude = 31.333
	b.Longitude = 105.234
	b.Intensity = 8

	datas := []AlarmInfo{a, b}
	list := AlarmDataList{}
	list.EventId = "20150323080000.00"
	list.Alarms = datas

	data, err := xml.MarshalIndent(list, "", " ")

	if err == nil {
		fmt.Printf("%s\n", data)
		t.Log("ok")

	} else {
		fmt.Printf("%s", err)
		t.Error("fail:")
	}
}
