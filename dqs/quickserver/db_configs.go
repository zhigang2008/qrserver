package quickserver

import (
	"encoding/xml"
)

//保存在数据库中的配置信息
var GlobalConfig DatabaseConfig

//数据库配置文件
type DatabaseConfig struct {
	CRC                bool            //是否进行CRC校验
	ReadWaveAfterAlarm bool            //是否在收到警报数据后立即发送波形记录读取命令
	EventParams        EventParameters //时间控制参数设置
	FileConfig         FilesConfig     //
	ReportCfg          ReportParameter
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

type ReportParameter struct {
	SleepTime   int //延后时间,单位分钟
	ReportLevel int //最低的报警级别

}
