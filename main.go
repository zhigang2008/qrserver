package main

import (
	//"bytes"
	//"encoding/binary"
	"fmt"
	//log "github.com/cihub/seelog"
	"syscall"
	//"time"
	"unsafe"
)

//传感器参数结构
type RetData struct {
	//--基本参数--
	SensorID      [11]byte //传感器编号
	SiteName      [11]byte //站点名称
	Longitude     float32  //台站经度
	Latitude      float32  //台站纬度
	SiteType      int      //场地类型
	ObserveObject int      //观测对象
	Accelerometer int      //加速度计型号
	Direction     int      //安装方向
	RangeType     int      //量程选择
	Period        float32  //采样周期
	RegionCode    [7]byte  //行政区划代码
	Custom1       [9]byte  //预留
	Custom2       [9]byte  //预留
	//--触发参数--
	PGATrigger          int     //PGA触发
	PGATrgThreshold     float32 //PGA阀值
	SITrigger           int     //SI触发
	SITrgThreshold      float32 //SI阀值
	CombTrigger         int     //组合触发
	ReserveTrigger      int     //预留触发
	ReserveTrgThreshold float32 //预留阀值
	//--报警参数--
	PGAAlert              int     //PGA报警
	PGAAlertThreshold     float32 //PGA报警阀值
	SIAlert               int     //SI报警
	SIAlertThreshold      float32 //SI报警阀值
	CombAlert             int     //组合报警
	ReserveAlert          int     //预留报警
	ReserveAlertThreshold float32 //预留报警阀值
	//--输出参数--
	DA1 int //DA输出1
	DA2 int //DA输出2
	IO1 int //IO输出1
	IO2 int //IO输出2
}

//突发数据
type FlashData struct {
	SeqNo         [11]byte //记录编号
	SensorId      [11]byte //传感器编号
	Longitude     float32  //经度
	Latitude      float32  //纬度
	SiteType      int      //场地类型
	ObserveObject int      //观测对象
	Direction     int      //安装方向
	RegionCode    [7]byte  //行政区域编码
	InitTime      [6]byte  //初始时刻
	Period        float32  //采用周期
	PGA           float32  //PGA值
	SI            float32  //SI值
	Length        float32  //记录长度
}

func main() {
	fmt.Println("begin....")
	fmt.Printf("%sssss\n", "SI30002345"+string(byte(0)))

	/*
		p := SensorParameter{}
		p.SensorID = "SI30001001"
		p.SiteName = "1234567890"
		p.Latitude = 1.1
		p.Longitude = 1.1
		p.SiteType = 1
		p.ObserveObject = 1
		p.Accelerometer = 1
		p.Direction = 1
	*/

	var ret RetData = RetData{}

	dll := syscall.MustLoadDLL("Socket1.dll")
	proc := dll.MustFindProc("parseReadSetParam")
	sid := []byte("SI30002345g005FSI30002345haidian00010584628032435441101100002000510802000000000000000010220000000032000000000000001045000000006700000000000000010030f6")
	ok, _, _ := proc.Call(
		uintptr(unsafe.Pointer(&sid[0])),
		uintptr(unsafe.Pointer(&ret)))
	if ok != 1 {
		fmt.Println(ok)

	}

	fmt.Printf("%s\n", string(ret.SiteName[:10]))
	fmt.Printf("%s\n", string(ret.SensorID[:10]))
	fmt.Printf("%f\n", ret.Longitude)
	fmt.Printf("%f\n", ret.Latitude)
	fmt.Printf("%d\n", ret.SiteType)
	fmt.Printf("%d\n", ret.Accelerometer)
	fmt.Printf("%d\n", ret.ObserveObject)
	fmt.Printf("%f\n", ret.Period)
	fmt.Printf("%s\n", string(ret.RegionCode[:]))
	fmt.Printf("%d\n", ret.DA2)
	fmt.Printf("%d\n", ret.IO1)
	fmt.Printf("%d\n", ret.IO2)

	fmt.Println("========================")
	var ret2 FlashData = FlashData{}
	proc2 := dll.MustFindProc("parseReadFlashParam")
	flashdata1 := []byte("SI30001051a003b1309010715SI3000105110375824030919431111326291309010715510000199904528266003049630000000ea46")
	ok1, _, _ := proc2.Call(
		uintptr(unsafe.Pointer(&flashdata1[0])),
		uintptr(unsafe.Pointer(&ret2)))
	if ok1 != 1 {
		fmt.Println(ok1)

	}
	s := fmt.Sprintf("%x", ret2.InitTime)
	fmt.Printf("%s\n", s)
}
