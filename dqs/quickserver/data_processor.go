package quickserver

import (
	"dqs/util"
	"errors"
	//"fmt"
	log "github.com/cihub/seelog"
	"net"
	"time"
)

//数据处理器结构
var dataProcessor *DataProcessor

//包含调用的dll以及 其中的function句柄
type DataProcessor struct {
	dataManager *DataManager
	alarmMap    *util.SafeMap
	analyzer    *EarthquakeAnalyzer
}

//初始化全局数据处理器
func InitDataProcessor(dm *DataManager) {
	dataProcessor = new(DataProcessor)
	dataProcessor.dataManager = dm
	dataProcessor.alarmMap = util.NewSafeMap()
	dataProcessor.analyzer = NewEarthquakeAnalyzer("default", dm)

}

//初始化数据处理器
func NewDataProcessor(dm *DataManager) *DataProcessor {
	var dp = new(DataProcessor)
	dp.dataManager = dm
	dataProcessor.alarmMap = util.NewSafeMap()
	dataProcessor.analyzer = NewEarthquakeAnalyzer("default", dm)
	return dp
}

//解析数据
func (dp *DataProcessor) DataProcess(content []byte, remote string, conn *net.Conn) (err error) {
	log.Info("Begin process data")

	log.Infof("设备:%s\n", content[0:10])
	log.Infof("功能:%c\n", content[10])
	//判断接收的数据类型
	datatype := content[10]

	switch datatype {
	case 'a', 'A': //突发数据
		c, ok := hasCommand(remote, "R")
		if ok == true {
			c <- content
		} else {
			dp.ProcessFlashData(content)
			//是否报警后,发送波形数据
			if GlobalConfig.ReadWaveAfterAlarm {
				dp.sendFlashReadCommand(string(content[0:10]), remote, conn)
			}
		}
	case 'z', 'Z': //设备注册
		dp.DeviceRegister(&content, remote)
		dp.sendStatusReadCommand(string(content[0:10]), conn)
		//发送r命令读取波形图
		//dp.sendFlashReadCommand(string(content[0:10]), remote, conn)
	case 'g', 'G': //状态及参数设定
		//判断是否有控制命令等待返回数据
		c, ok := hasCommand(remote, "G")
		if ok == true {
			c <- content
		} else {
			dp.ProcessStatusData(content)
		}
	case 'r', 'R': //地形波读取
		//fmt.Println("波形信息读取")
		dp.ProcessWaveData(content)
	case 's', 'S': //设置参数
		//判断是否有控制命令等待返回数据
		c, ok := hasCommand(remote, "S")
		if ok == true {
			c <- content
		}
	default:
		log.Info("无效数据")
	}
	return
}

//处理突发数据
func (dp *DataProcessor) ProcessFlashData(content []byte) (err error) {

	//进行数据处理
	id := string(content[0:10])

	if GlobalConfig.CRC {
		//先进行CRC校验.无效数据直接抛弃.
		if DllUtil.CheckCRCCode(content) != true {
			log.Warnf("[%s]设备报警数据CRC校验失败", id)
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
	//烈度值查表

	//err = dp.dataManager.FlashDataSave(sData)
	err = dp.dataManager.AlarmUpsert(sData)
	if err != nil {
		log.Warnf("[%s]报警信息处理失败:%s", id, err.Error())
		return err
	}
	log.Infof("报警信息保存成功")

	//------写入文件-----
	if GlobalConfig.FileConfig.WriteFile {
		go writeAlarm(sData)
	}
	//-----进入震情分析过程------
	go dp.analyzer.analyze(sData)

	return nil
}

//处理状态数据
func (dp *DataProcessor) ProcessStatusData(content []byte) error {
	id := string(content[0:10])

	if GlobalConfig.CRC {
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
		n, err0 := (*connP).Write(command)

		if err0 != nil {
			log.Warnf("向[%s]设备发送状态读取指令失败:%s", deviceid, err0.Error())
		} else {
			log.Infof("向[%s]设备发送状态读取指令成功:%d", deviceid, n)
		}
	}
}

//发送波形图读取命令
func (dp *DataProcessor) sendFlashReadCommand(deviceid string, remote string, connP *net.Conn) {
	command, err := DllUtil.GenerateFlashReadParam(deviceid)
	if err == nil {
		//发送控制命令
		n, err0 := (*connP).Write(command)

		if err0 != nil {
			log.Warnf("向[%s]设备发送波形图读取指令失败:%s", deviceid, err0.Error())
		} else {
			log.Infof("向[%s]设备发送波形图读取指令成功:%d", deviceid, n)

			//remote := (*connP).RemoteAddr().String()
			//等着读取返回的首帧报警数据
			c := make(chan []byte)
			AddCommand(remote, "R", c)
			back := <-c
			//处理
			dp.ProcessWaveFlashData(back)
			//取消控制命令
			DeleteCommand(remote, "R")

		}
	}
}

//处理波形图的突发数据
func (dp *DataProcessor) ProcessWaveFlashData(content []byte) (err error) {

	//进行数据处理
	id := string(content[0:10])

	if GlobalConfig.CRC {
		//先进行CRC校验.无效数据直接抛弃.
		if DllUtil.CheckCRCCode(content) != true {
			log.Warnf("[%s]设备报警数据CRC校验失败", id)
			return errors.New("CRC校验失败,数据非法")
		}
	}

	//调用dll解析
	data, err := DllUtil.ParseReadFlashParam(content[0 : len(content)-4])
	if err != nil {
		log.Warnf("[%s]波形记录报警信息DLL解析失败:%s", id, err.Error())
		return err
	}
	//数据转换
	sData := FlashData2AlarmInfo(data)

	//更新报警信息
	sData.HasWaveInfo = true
	err = dp.dataManager.AlarmUpsert(sData)
	if err != nil {
		log.Warnf("[%s]波形记录报警信息更新失败:%s", id, err.Error())
	}

	//记录波形记录
	wData := WaveInfo{}
	wData.Alarm = *sData
	wData.SensorId = sData.SensorId
	wData.SeqNo = sData.SeqNo

	oData, _ := dp.dataManager.GetWaveData(sData.SensorId, sData.SeqNo)
	if oData.Id != "" {
		oData.Alarm = *sData
		oData.LUD = time.Now()
		err = dp.dataManager.WaveDataUpdate(&oData)
		if err != nil {
			log.Warnf("[%s]波形图之报警信息更新处理失败:%s", id, err.Error())
			return err
		}
	} else {
		err = dp.dataManager.WaveDataAdd(&wData)
		if err != nil {
			log.Warnf("[%s]波形图之报警信息处理失败:%s", id, err.Error())
			return err
		}
	}

	//更新设备波形标记
	dp.alarmMap.Set(sData.SensorId, sData.SeqNo)

	log.Infof("波形记录之报警信息保存成功")
	return nil
}

//波形图数据接收
func (dp *DataProcessor) ProcessWaveData(content []byte) (err error) {

	//进行数据处理
	id := string(content[0:10])

	if GlobalConfig.CRC {
		//先进行CRC校验.无效数据直接抛弃.
		if DllUtil.CheckCRCCode(content) != true {
			log.Warnf("[%s]设备波形图数据CRC校验失败", id)
			return errors.New("CRC校验失败,数据非法")
		}
	}

	//调用dll解析
	data, frame, err := DllUtil.ParseFlashData(content[0:len(content)-4], id)
	if err != nil {
		log.Warnf("[%s]波形图信息DLL解析失败:%s", id, err.Error())
		return err
	}
	log.Infof("波形图%d帧数据:%d", frame, data)

	//数据处理

	seqno := dp.alarmMap.Get(id)
	var wData = WaveInfo{}
	if v, ok := seqno.(string); ok == true {
		wData, err = dp.dataManager.GetWaveData(id, v)
		if err != nil {
			log.Warnf("获取[%s-%s]波形图失败,将获取最后时间的波形图记录:%s", id, seqno, err.Error())
		}
	}
	if wData.SeqNo == "" {
		wData, err = dp.dataManager.GetLastWave(id)
		if err != nil {
			log.Warnf("[%s]获取最新的波形图记录失败:%s", id, err.Error())
			return errors.New("获取最新波形图失败[" + err.Error() + "]")
		}
	}

	//var i int16
	//x分量
	if frame >= 0 && frame < 25 {
		//for i = 0; i < 240; i++ {
		//	wData.X_data[frame*240+i] = data[i]
		//}
		copy(wData.X_data[frame*240:(frame+1)*240-1], data[:])
	} else if frame >= 25 && frame < 50 {
		//for i = 0; i < 240; i++ {
		//	wData.Y_data[(frame-25)*240+i] = data[i]
		//}
		copy(wData.Y_data[(frame-25)*240:(frame-24)*240-1], data[:])

	} else if frame >= 50 && frame < 75 {
		//for i = 0; i < 240; i++ {
		//	wData.Z_data[(frame-50)*240+i] = data[i]
		//}
		copy(wData.Z_data[(frame-50)*240:(frame-49)*240-1], data[:])
	}

	err = dp.dataManager.WaveDataUpdate(&wData)
	if err != nil {
		log.Warnf("[%s]波形记录信息处理失败:%s", id, err.Error())
		return err
	}
	log.Infof("波形图信息保存成功")

	return nil
}
