package quickserver

import (
	//"bytes"
	//"encoding/binary"
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
	"syscall"
	"time"
	"unsafe"
)

//数据处理器结构
//包含调用的dll以及 其中的function句柄
type DataProcessor struct {
	dataManager           *DataManager
	dll                   *syscall.DLL
	p_parseReadFlashParam *syscall.Proc
	p_parseReadSetParam   *syscall.Proc
	p_ParseDelParam       *syscall.Proc
	p_ParseSetParam       *syscall.Proc
	p_GenerateSetParam    *syscall.Proc
	p_parseFlashData      *syscall.Proc
}

//初始化数据处理器
func NewDataProcessor(dm *DataManager) *DataProcessor {
	var dp = new(DataProcessor)
	dp.dataManager = dm
	dp.dll = syscall.MustLoadDLL("socket1.dll")
	dp.p_parseReadFlashParam = dp.dll.MustFindProc("parseReadFlashParam")
	dp.p_parseReadSetParam = dp.dll.MustFindProc("parseReadSetParam")
	dp.p_ParseDelParam = dp.dll.MustFindProc("ParseDelParam")
	dp.p_ParseSetParam = dp.dll.MustFindProc("ParseSetParam")
	dp.p_GenerateSetParam = dp.dll.MustFindProc("GenerateSetParam")
	dp.p_parseFlashData = dp.dll.MustFindProc("parseFlashData")

	return dp
}

//释放Dll资源
func (dp *DataProcessor) FreeDLL() {
	dp.dll.Release()
}

//DLL解析接收的突发数据
func (dp *DataProcessor) parseReadFlashParam(rec []byte) (*FlashData, error) {
	flashData := FlashData{}

	ok, _, _ := dp.p_parseReadFlashParam.Call(
		uintptr(unsafe.Pointer(&rec[0])),
		uintptr(unsafe.Pointer(&flashData)))
	if ok != 1 {
		return nil, errors.New("DLL解析突发数据失败")
	}
	return &flashData, nil
}

//DLL解析接收的设置数据
func (dp *DataProcessor) parseReadSetParam(rec []byte) (*RetData, error) {
	retData := RetData{}

	ok, _, _ := dp.p_parseReadSetParam.Call(
		uintptr(unsafe.Pointer(&rec[0])),
		uintptr(unsafe.Pointer(&retData)))
	if ok != 1 {
		return nil, errors.New("DLL解析设备的设置参数失败")
	}
	return &retData, nil
}

//DLL解析删除设备参数是否成功
func (dp *DataProcessor) parseDelParam(rec []byte) bool {
	ok, _, _ := dp.p_ParseDelParam.Call(
		uintptr(unsafe.Pointer(&rec[0])))
	if ok == 0 {
		return true
	} else {
		return false
	}
}

//DLL解析删除设备参数是否成功
func (dp *DataProcessor) parseSetParam(rec []byte) bool {
	ok, _, _ := dp.p_ParseSetParam.Call(
		uintptr(unsafe.Pointer(&rec[0])))
	if ok == 0 {
		return true
	} else {
		return false
	}
}

//解析数据
func (dp *DataProcessor) DataProcess(content []byte) (err error) {
	log.Info("Begin process data")

	/*
		contbuf := bytes.NewBuffer(content)
		var modbus ModBus
		err0 := binary.Read(contbuf, binary.BigEndian, &modbus)
		if err0 != nil {
			fmt.Printf("转换错误 %s\n", err0.Error())
		}

		fmt.Printf("读取到的modbus数据 %x\n", modbus)
	*/
	fmt.Printf("设备:%s\n", content[0:10])
	fmt.Printf("功能:%c\n", content[10])
	//判断接收的数据类型
	datatype := content[10]

	switch datatype {
	case 'a', 'A': //突发数据
		dp.ProcessFlashData(content)
	case 'z', 'Z': //设备注册
		dp.DeviceRegister(&content)
	case 'g', 'G': //状态及参数设定
		dp.ProcessStatusData(content)
	case 'r', 'R': //地形波读取
		fmt.Println("波形信息读取")
	default:
		fmt.Println("无效数据")
	}
	return
}

//处理突发数据
func (dp *DataProcessor) ProcessFlashData(content []byte) (err error) {
	id := string(content[0:10])
	//调用dll解析
	data, err := dp.parseReadFlashParam(content)
	if err != nil {
		log.Warnf("[%s]报警信息DLL解析失败:%s", id, err.Error())
		return err
	}
	//数据转换
	sData := FlashData2AlarmInfo(data)
	err = dp.dataManager.FlashDataSave(sData)
	if err != nil {
		log.Warnf("[%s]报警信息处理失败:%s", id, err.Error())
		return err
	}
	log.Infof("报警信息保存成功")
	return
}

//处理状态数据
func (dp *DataProcessor) ProcessStatusData(content []byte) (err error) {
	id := string(content[0:10])
	//调用dll解析
	data, err := dp.parseReadSetParam(content)
	if err != nil {
		log.Warnf("[%s]状态信息DLL解析失败:%s", id, err.Error())
		return err
	}
	//数据转换
	sData := RetData2SensorInfo(data)
	err = dp.dataManager.UpdateDeviceStatus(sData)
	if err != nil {
		log.Warnf("[%s]状态信息处理失败:%s", id, err.Error())
		return err
	}
	log.Infof("设备状态及参数读取成功")
	return
}

//设备注册
func (dp *DataProcessor) DeviceRegister(content *[]byte) (err error) {
	device := new(DeviceInfo)
	device.SensorId = string((*content)[0:10])
	device.Online = true
	device.RegisterTime = time.Now()
	device.UpdateTime = time.Now()
	err = dp.dataManager.DeviceRegister(device)
	if err != nil {
		log.Warnf("设备注册失败:%s", err.Error())
		return err
	}
	log.Infof("设备注册成功")
	return
}

//设备下线
func (dp *DataProcessor) DeviceOffline(remote string) {
	err := dp.dataManager.DeviceOffline(remote)
	if err != nil {
		log.Warnf("设备[%s]下线失败:%s", remote, err.Error())
	}
	log.Infof("设备[%s]下线", remote)
}
