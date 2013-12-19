package models

//系统配置数据
type SystemConfig struct {
	UserDefaultPassword string
	UseGis              bool
	GisServiceUrl       string
	GisServiceParams    string
	GisLayerBasic       string
	GisLayerChina       string
}
