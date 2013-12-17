package quickserver

import (
	"encoding/xml"
	"fmt"
	log "github.com/cihub/seelog"
	"io/ioutil"
)

const (
	configFile  = "server.xml"
	defaultHost = ""
	defaultPort = 8083
	defaultType = "tcp4"
	TCP4        = "tcp4"
	TCP6        = "tcp6"
)

//全局的服务器配置
var (
	ServerConfigs ServerConfig
)

//服务器配置信息结构
type ServerConfig struct {
	XMLName            xml.Name `xml:"Server"`
	Host               string
	Port               int
	Type               string
	Database           DataServerConfig
	HttpServerEnable   bool
	CRC                bool
	ReadWaveAfterAlarm bool // 是否报警后立即读取波形记录
	EventParams        EventParameters
	FileConfig         FilesConfig
}

//数据库配置文件
type DataServerConfig struct {
	XMLName          xml.Name `xml:"DataServer"`
	Host             string
	Port             int
	DataBaseName     string
	DataCollection   string
	DeviceCollection string
}

//震情事件判断的参数
type EventParameters struct {
	XMLName              xml.Name `xml:"EventParams"`
	SignalTimeSpan       int      //有效震情信号判断用到的时间宽度设定
	ValidEventAlarmCount int      //确认一个震情事件是否有效,报警站点最低数量
	NewEventTimeGap      int      //一个报警消息是否属于新的震情事件,其与上个事件的时间间隔
	NewEventGapMultiple  float64  //报警信息与上个事件报警信息平均量的间隔时间倍数
}

//数据文件保存设置
type FilesConfig struct {
	XMLName       xml.Name `xml:"FileConfig"`
	WriteFile     bool
	FileDir       string
	ReportFileDir string
}

//读取配置文件,并进行校验
func ReadConfigFromFile() (ServerConfig, error) {

	ServerConfigs.Host = defaultHost
	ServerConfigs.Port = defaultPort
	ServerConfigs.Type = defaultType

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Error(err)
		return ServerConfigs, err
	}

	err = xml.Unmarshal(content, &ServerConfigs)
	if err != nil {
		log.Error(err)
		return ServerConfigs, err
	}
	log.Info(ServerConfigs)

	err = CheckConfig(ServerConfigs)
	return ServerConfigs, err
}

/*检查读取到的配置文件*/
func CheckConfig(conf ServerConfig) (err error) {
	if conf.Port < 1 || conf.Port > 65535 {
		err = fmt.Errorf("Server Port must be in (1 ~ 65535)")
		return
	}
	if !(conf.Type == TCP4 || conf.Type == TCP6) {
		err = fmt.Errorf("TCP Type only be 'tcp4' or 'tcp6' ")
		return
	}
	if conf.Database.Host == "" {
		err = fmt.Errorf("必须设置数据库地址")
		return
	}
	return
}
