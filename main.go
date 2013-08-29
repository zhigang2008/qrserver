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
type SensorParameter struct {
	//--基本参数--
	SensorID      string  //传感器编号
	SiteName      string  //站点名称
	Longitude     float32 //台站经度
	Latitude      float32 //台站纬度
	SiteType      byte    //场地类型
	ObserveObject byte    //观测对象
	Accelerometer byte    //加速度计型号
	Direction     byte    //安装方向
	RangeType     byte    //量程选择
	RegionCode    string  //行政区划代码
	Custom1       string  //预留
	Custom2       string  //预留
	//--触发参数--
	PGATrigger          byte    //PGA触发
	PGATrgThreshold     float32 //PGA阀值
	SITrigger           byte    //SI触发
	SITrgThreshold      float32 //SI阀值
	CombTrigger         byte    //组合触发
	ReserveTrigger      byte    //预留触发
	ReserveTrgThreshold float32 //预留阀值
	//--报警参数--
	PGAAlert              byte    //PGA报警
	PGAAlertThreshold     float32 //PGA报警阀值
	SIAlert               byte    //SI报警
	SIAlertThreshold      float32 //SI报警阀值
	CombAlert             byte    //组合报警
	ReserveAlert          byte    //预留报警
	ReserveAlertThreshold float32 //预留报警阀值
	//--输出参数--
	DA1 byte //DA输出1
	DA2 byte //DA输出2
	IO1 byte //IO输出1
	IO2 byte //IO输出2
}

func main() {
	fmt.Println("begin....")
	p := SensorParameter{}
	p.SensorID = "SI30001001"
	p.SiteName = "1234567890"
	p.Latitude = 1.1
	p.Longitude = 1.1
	p.SiteType = 1
	p.ObserveObject = 1
	p.Accelerometer = 1
	p.Direction = 1

	dll := syscall.MustLoadDLL("socket1.dll")
	proc := dll.MustFindProc("GenerateReadParam")
	sid := []byte("SI30001001s")
	ret := make([]byte, 1024)
	ok, _, _ := proc.Call(
		uintptr(unsafe.Pointer(&sid[0])),
		uintptr(unsafe.Pointer(&ret[0])))
	if ok != 0 {
		fmt.Println(ok)
	}
	fmt.Println(string(ret))

}
