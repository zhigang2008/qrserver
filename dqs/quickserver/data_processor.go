package quickserver

import (
	//"bytes"
	//"encoding/binary"
	"fmt"
	log "github.com/cihub/seelog"
	"syscall"
	"time"
)

//数据处理器结构
//包含调用的dll以及 其中的function句柄
type DataProcessor struct {
	dataManager         *DataManager
	dll                 *syscall.DLL
	p_sendStr           *syscall.Proc
	p_GenerateReadParam *syscall.Proc
}

//初始化数据处理器
func NewDataProcessor(dm *DataManager) *DataProcessor {
	var dp = new(DataProcessor)
	dp.dataManager = dm
	dp.dll = syscall.MustLoadDLL("socket1.dll")
	dp.p_sendStr = dp.dll.MustFindProc("sendStr")
	dp.p_GenerateReadParam = dp.dll.MustFindProc("GenerateReadParam")
	return dp
}

//释放Dll资源
func (dp *DataProcessor) FreeDLL() {
	dp.dll.Release()
}

//读取参数设置
//TODO
func (dp *DataProcessor) GenerateReadParam(strParam string) []byte {

	var ret []byte
	dp.p_GenerateReadParam.Call()
	return ret
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
	case 'a', 'A':
		dp.ProcessAlert(&content)
	case 'z', 'Z':
		dp.DeviceRegister(&content)
	case 'g', 'G':
		fmt.Println("设置信息或者状态信息处理")
	case 'r', 'R':
		fmt.Println("波形信息读取")
	default:
		fmt.Println("无效数据")
	}
	return
}

func (dp *DataProcessor) ProcessAlert(content *[]byte) (err error) {
	data := new(DataAlert)
	data.SensorId = string((*content)[0:10])
	err = dp.dataManager.DataSave(data)
	if err != nil {
		log.Warnf("报警信息处理失败:%s", err.Error())
	}
	log.Infof("报警信息保存成功")
	return
}
func (dp *DataProcessor) DeviceRegister(content *[]byte) (err error) {
	device := new(DeviceInfo)
	device.SensorId = string((*content)[0:10])
	device.Online = true
	device.RegisterTime = time.Now()
	err = dp.dataManager.DeviceRegister(device)
	if err != nil {
		log.Warnf("设备注册失败:%s", err.Error())
	}
	log.Infof("设备注册成功")
	return
}
