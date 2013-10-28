package quickserver

import (
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
	"net"
	"time"
)

//数据处理器结构
var dataProcessor *DataProcessor

//包含调用的dll以及 其中的function句柄
type DataProcessor struct {
	dataManager *DataManager
}

//初始化全局数据处理器
func InitDataProcessor(dm *DataManager) {
	dataProcessor = new(DataProcessor)
	dataProcessor.dataManager = dm
}

//初始化数据处理器
func NewDataProcessor(dm *DataManager) *DataProcessor {
	var dp = new(DataProcessor)
	dp.dataManager = dm
	return dp
}

//解析数据
func (dp *DataProcessor) DataProcess(content []byte, remote string, conn *net.Conn) (err error) {
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
		dp.sendStatusReadCommand(string(content[0:10]), conn)
	case 'g', 'G': //状态及参数设定
		//判断是否有控制命令等待返回数据
		c, ok := hasCommand(remote, "G")
		if ok == true {
			c <- content
		} else {
			dp.ProcessStatusData(content)
			//发送r命令读取波形图
			dp.sendFlashReadCommand(string(content[0:10]), conn)
		}
	case 'r', 'R': //地形波读取
		fmt.Println("波形信息读取")
	case 's', 'S': //设置参数
		//判断是否有控制命令等待返回数据
		c, ok := hasCommand(remote, "S")
		if ok == true {
			c <- content
		}
	default:
		fmt.Println("无效数据")
	}
	return
}

//处理突发数据
func (dp *DataProcessor) ProcessFlashData(content []byte) (err error) {

	//进行数据处理
	id := string(content[0:10])

	if ServerConfigs.CRC {
		//先进行CRC校验.无效数据直接抛弃.
		if DllUtil.CheckCRCCode(content) != true {
			log.Warnf("[%s]设备报警态数据CRC校验失败", id)
			return errors.New("CRC校验失败,数据非法")
		}
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
	return nil
}

//处理状态数据
func (dp *DataProcessor) ProcessStatusData(content []byte) error {
	id := string(content[0:10])

	if ServerConfigs.CRC {
		//先进行CRC校验.无效数据直接抛弃.
		if DllUtil.CheckCRCCode(content) != true {
			log.Warnf("[%s]设备状态数据CRC校验失败", id)
			return errors.New("CRC校验失败,数据非法")
		}
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
	return nil
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

//所有设备下线
func (dp *DataProcessor) ResetAllDeviceStatus() {
	err := dp.dataManager.ResetAllDeviceStatus()
	if err != nil {
		log.Warnf("所有设备状态重置失败:[%s]", err.Error())
	}
	log.Infof("所有设备状态重置为offline")
}

//发送状态读取命令
func (dp *DataProcessor) sendStatusReadCommand(deviceid string, connP *net.Conn) {
	command, err := DllUtil.GenerateReadParam(deviceid)
	if err == nil {
		//发送控制命令
		_, err0 := (*connP).Write(command)

		if err0 != nil {
			log.Warnf("向[%s]设备发送状态读取指令失败:%s", deviceid, err0.Error())
		} else {
			log.Infof("向[%s]设备发送状态读取指令成功:%d", deviceid)
		}
	}
}

//发送波形图读取命令
func (dp *DataProcessor) sendFlashReadCommand(deviceid string, connP *net.Conn) {
	command, err := DllUtil.GenerateFlashReadParam(deviceid)
	if err == nil {
		//发送控制命令
		_, err0 := (*connP).Write(command)

		if err0 != nil {
			log.Warnf("向[%s]设备发送波形图读取指令失败:%s", deviceid, err0.Error())
		} else {
			log.Infof("向[%s]设备发送波形图读取指令成功:%d", deviceid)
		}
	}
}
