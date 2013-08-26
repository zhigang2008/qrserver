package quickserver

import (
	//"fmt"
	log "github.com/cihub/seelog"
)

//解析数据
func DataProcess(content []byte, datamgr *DataManager) (err error) {
	log.Info("Begin process data")
	data := new(DataAlert)
	data.SeqNo = string(content[0:1])
	datamgr.DataSave(data)
	return
}
