package quickserver

import (
	"fmt"
	log "github.com/cihub/seelog"
)

//解析数据
func DataProcess(content []byte, datamgr *DataManager) (err error) {
	log.Info("Begin process data")

	fmt.Printf("设备:%d\n", content[0])
	fmt.Printf("功能:%c\n", content[1])
	//判断接收的数据类型
	datatype := content[1]

	switch datatype {
	case I_Alert, I_AlertUp:
		processAlert(&content, datamgr)
	case I_Status:
		fmt.Println(I_Status)
	case I_RecordData:
		fmt.Println(I_RecordData)
	default:
		fmt.Println("无效数据")
	}
	return
}

func processAlert(content *[]byte, datamgr *DataManager) (err error) {
	data := new(DataAlert)
	data.SeqNo = string((*content)[0])
	datamgr.DataSave(data)
	return
}
