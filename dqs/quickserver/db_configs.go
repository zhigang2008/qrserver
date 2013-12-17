package quickserver

//保存在数据库中的配置信息
var DBConfig DatabaseConfig

//数据库配置文件
type DatabaseConfig struct {
	EventParams EventParameters //时间控制参数设置
	FileConfig  FilesConfig     //
	ReportCfg   ReportConfig
}

type ReportConfig struct {
	SleepTime int //延后时间,单位分钟

}
