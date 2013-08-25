package quickserver

import (
	//"fmt"
	log "github.com/cihub/seelog"
)

//解析数据
func DataProcess(content []byte, datamgr *DataManager) (err error) {
	data := new(DataAlert)
	log.Infof("cont:%s", string(content[0:]))
	data.SeqNo = string(content[0:1])
	datamgr.DataSave(data)
	return
}
