package models

//系统配置数据
type SystemConfig struct {
	UserDefaultPassword string
	UseGis              bool
	GisServiceUrl       string
	GisServiceParams    string
	GisLayerBasic       string
	GisLayerChina       string
	GisImageCfg         GisImageConfig
}

//gis图片设置
type GisImageConfig struct {
	SRS    string
	BBOX   string
	Height string
	Width  string
	Format string
}
