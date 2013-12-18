package models

//数据库配置文件
type DatabaseConfig struct {
	CRC                bool
	ReadWaveAfterAlarm bool            // 是否报警后立即读取波形记录
	EventParams        EventParameters //时间控制参数设置
	FileConfig         FilesConfig     //
	ReportCfg          ReportParameter
}

//震情事件判断的参数
type EventParameters struct {
	SignalTimeSpan       int     //有效震情信号判断用到的时间宽度设定
	ValidEventAlarmCount int     //确认一个震情事件是否有效,报警站点最低数量
	NewEventTimeGap      int     //一个报警消息是否属于新的震情事件,其与上个事件的时间间隔
	NewEventGapMultiple  float64 //报警信息与上个事件报警信息平均量的间隔时间倍数
}

//数据文件保存设置
type FilesConfig struct {
	WriteFile     bool
	FileDir       string
	ReportFileDir string
}

//速报相关的配置
type ReportParameter struct {
	SleepTime   int //延后时间,单位分钟
	ReportLevel int //最低的报警级别

}
