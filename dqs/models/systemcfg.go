package models

//系统配置数据
type SystemConfig struct {
	UserDefaultPassword string
	DisableGoogleMap    bool
	UseGis              bool
	GisServiceUrl       string
	GisServiceParams    string
	GisLayerBasic       string
	GisLayerChina       string
	GisImageCfg         GisImageConfig
	MmsCfg              MMSConfig
	MailCfg             MailConfig
}

//gis图片设置
type GisImageConfig struct {
	SRS    string
	BBOX   string
	Height string
	Width  string
	Format string
}

//彩信服务
type MMSConfig struct {
	MmsEnable  bool
	ServiceUrl string
	UserNo     string
	Password   string
}

//邮件服务
type MailConfig struct {
	MailHost     string
	MailPort     string
	MailAddr     string
	NeedAuth     bool
	MailUser     string
	MailPassword string
}
