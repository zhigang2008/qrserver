package quickserver

import (
	//"fmt"
	log "github.com/cihub/seelog"
)

//解析数据
func DataProcess(content []byte) (err error) {
	data := new(DataAlert)
	log.Infof("cont:%s", string(content[0:]))
	data.SeqNo = string(content[0:1])
	DataSave2()
	return
}

//解析数据
func DataProcess2() (err error) {
	data := new(DataAlert)
	log.Infof("cont:%s", "hello")
	data.SeqNo = string("seq")
	DataSave(data)
	return
}
