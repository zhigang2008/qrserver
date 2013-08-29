package quickserver

import (
	//"bytes"
	//"encoding/binary"
	"fmt"
	log "github.com/cihub/seelog"
	"syscall"
)

//数据处理器结构
//包含调用的dll以及 其中的function句柄
type DataProcessor struct {
	dll                 *syscall.DLL
	p_sendStr           *syscall.Proc
	p_GenerateReadParam *syscall.Proc
}

//初始化数据处理器
func NewDataProcessor() *DataProcessor {
	var dp = new(DataProcessor)
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
func (dp *DataProcessor) DataProcess(content []byte, datamgr *DataManager) (err error) {
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
	case I_Alert, I_AlertUp:
		dp.ProcessAlert(&content, datamgr)
	case I_Status:
		fmt.Println(I_Status)
	case I_RecordData:
		fmt.Println(I_RecordData)
	default:
		fmt.Println("无效数据")
	}
	return
}

func (dp *DataProcessor) ProcessAlert(content *[]byte, datamgr *DataManager) (err error) {
	data := new(DataAlert)
	data.SeqNo = string((*content)[0:10])
	datamgr.DataSave(data)
	return
}
