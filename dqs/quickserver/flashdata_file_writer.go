package quickserver

import (
	"dqs/util"
	"encoding/xml"
	log "github.com/cihub/seelog"
	"os"
	"path"
	"time"
)

const (
	FileTimeLayout = "20060102150405.000000"
)

func writeAlarm(a *AlarmInfo) {

	xmlcontent, err1 := xml.MarshalIndent(a, "  ", "    ")
	if err1 != nil {
		log.Warnf("报警信息解析xml出错:%s", err1.Error())
		return
	}

	dir := ServerConfigs.FileConfig.FileDir
	//先判断目录是否存在,不存在则创建
	if util.IsDirExist(dir) == false {
		os.MkdirAll(dir, 0777)
	}

	filename := time.Now().Format(FileTimeLayout) + ".xml"
	file, err := os.Create(path.Join(dir, filename))
	if err != nil {
		log.Warnf("创建文件[%s]失败:%s", path.Join(dir, filename), err.Error())
		return
	}
	defer file.Close()

	file.Write([]byte(xml.Header))
	file.Write(xmlcontent)

}
