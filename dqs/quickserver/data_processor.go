package quickserver

import (
	"fmt"
	log "github.com/cihub/seelog"
	"time"
)

//数据处理器结构
//包含调用的dll以及 其中的function句柄
type DataProcessor struct {
	dataManager *DataManager
}

//初始化数据处理器
func NewDataProcessor(dm *DataManager) *DataProcessor {
	var dp = new(DataProcessor)
	dp.dataManager = dm
	return dp
}

//解析数据
func (dp *DataProcessor) DataProcess(content []byte, remote string) (err error) {
	log.Info("Begin process data")

	fmt.Printf("设备:%s\n", content[0:10])
	fmt.Printf("功能:%c\n", content[10])
	//判断接收的数据类型
	datatype := content[10]

	switch datatype {
	case 'a', 'A': //突发数据
		dp.ProcessFlashData(content)
	case 'z', 'Z': //设备注册
		dp.DeviceRegister(&content, remote)
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

	//进行数据处理
	id := string(content[0:10])

	//先进行CRC校验.无效数据直接抛弃.
	if DllUtil.CheckCRCCode(content) != true {
		log.Warnf("[%s]设备报警态数据CRC校验失败:%s", id, err.Error())
		return
	}

	//调用dll解析
	data, err := DllUtil.ParseReadFlashParam(content[0 : len(content)-4])
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
	//先进行CRC校验.无效数据直接抛弃.
	if DllUtil.CheckCRCCode(content) != true {
		log.Warnf("[%s]设备状态数据CRC校验失败:%s", id, err.Error())
		return
	}

	//调用dll解析
	data, err := DllUtil.ParseReadSetParam(content[0 : len(content)-4])
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
func (dp *DataProcessor) DeviceRegister(content *[]byte, remote string) (err error) {
	device := new(DeviceInfo)
	device.SensorId = string((*content)[0:10])
	device.Online = true
	device.RegisterTime = time.Now()
	device.UpdateTime = time.Now()
	device.RemoteAddr = remote

	err = dp.dataManager.DeviceRegister(device)
	if err != nil {
		log.Warnf("设备注册失败:%s", err.Error())
		return err
	}
	log.Infof("设备注册成功")
	return
}

//设备下线
func (dp *DataProcessor) DeviceOffline(deviceid string) {
	err := dp.dataManager.DeviceOffline(deviceid)
	if err != nil {
		log.Warnf("设备[%s]下线失败:%s", deviceid, err.Error())
	}
	log.Infof("设备[%s]下线", deviceid)
}
