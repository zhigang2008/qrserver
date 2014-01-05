package quickserver

import (
	"encoding/xml"
	log "github.com/cihub/seelog"
	"sort"
)

var (
	GlobalConfig      DatabaseConfig //保存在数据库中的配置信息
	GlobalDataMapping DataMapping    //烈度对照表
	SystemConfigs     SystemConfig   //基本配置
)

//运行参数配置
type DatabaseConfig struct {
	CRC                 bool            //是否进行CRC校验
	ReadWaveAfterAlarm  bool            //是否在收到警报数据后立即发送波形记录读取命令
	IntensityMapingData string          //使用PGA/SI计算烈度. 默认PGA
	EventParams         EventParameters //时间控制参数设置
	FileConfig          FilesConfig     //
	ReportCfg           ReportParameter
}

//震情事件判断的参数
type EventParameters struct {
	XMLName              xml.Name `xml:"EventParams"`
	SignalTimeSpan       int      //有效震情信号判断用到的时间宽度设定
	ValidEventAlarmCount int      //确认一个震情事件是否有效,报警站点最低数量
	NewEventTimeGap      int      //一个报警消息是否属于新的震情事件,其与上个事件的时间间隔
	NewEventGapMultiple  float64  //报警信息与上个事件报警信息平均量的间隔时间倍数
	MinEventRecordLevel  int      //报警信息达到该级别才进行事件记录
}

//数据文件保存设置
type FilesConfig struct {
	XMLName       xml.Name `xml:"FileConfig"`
	WriteFile     bool
	FileDir       string
	ReportFileDir string
}

//报表处理参数
type ReportParameter struct {
	SleepTime          int  //延后时间,单位分钟
	ReportLevel        int  //最低的报警级别
	AuditBeforeSend    bool //发送前是否进行审核
	MinDirectSendLevel int  //该级别以上自动发送,无需审核

}

//系统配置数据
type SystemConfig struct {
	UserDefaultPassword string
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

//烈度对照数据
type DataMapping struct {
	PGAMap PGAMapArray
	SIMap  SIMapArray
}
type PGAMapping struct {
	PGA       float32
	Intensity int //仪器烈度值
}
type SIMapping struct {
	SI        float32
	Intensity int //仪器烈度值
}

//排序参数
type PGAMapArray []PGAMapping

func (p PGAMapArray) Len() int           { return len(p) }
func (p PGAMapArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PGAMapArray) Less(i, j int) bool { return p[i].PGA < p[j].PGA }

//排序参数
type SIMapArray []SIMapping

func (p SIMapArray) Len() int           { return len(p) }
func (p SIMapArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p SIMapArray) Less(i, j int) bool { return p[i].SI < p[j].SI }

//全部初始化
func initGlobalConfigs() {
	initRuntimeConfigs()
	initDataMapping()
	InitSystemConfigs()
}

//初始化运行参数
func initRuntimeConfigs() {

	configs, err := server.dataManager.GetGlobalConfigs()
	if err != nil {
		log.Warnf("从数据库中获取运行参数失败:%s", err.Error())
		if err == ErrNotFound {
			log.Warnf("系统将自动初始化运行参数")
			GlobalConfig = DatabaseConfig{}
			GlobalConfig.CRC = false
			GlobalConfig.ReadWaveAfterAlarm = true
			GlobalConfig.IntensityMapingData = "PGA"
			GlobalConfig.EventParams.SignalTimeSpan = 5
			GlobalConfig.EventParams.ValidEventAlarmCount = 3
			GlobalConfig.EventParams.NewEventTimeGap = 15
			GlobalConfig.EventParams.NewEventGapMultiple = 2.2
			GlobalConfig.EventParams.MinEventRecordLevel = 3
			GlobalConfig.FileConfig.WriteFile = true
			GlobalConfig.FileConfig.FileDir = "./output/alarms"
			GlobalConfig.FileConfig.ReportFileDir = "./output/reports"
			GlobalConfig.ReportCfg.SleepTime = 3
			GlobalConfig.ReportCfg.ReportLevel = 5
			GlobalConfig.ReportCfg.AuditBeforeSend = true
			GlobalConfig.ReportCfg.MinDirectSendLevel = 7

			errc := server.dataManager.CreateGlobalConfigs(&GlobalConfig)
			if errc != nil {
				log.Warnf("初始化运行参数失败:%s", errc.Error())
				return
			}
		} else {
			return
		}
	} else {
		GlobalConfig = configs
	}
}

//初始化烈度对照表
func initDataMapping() {

	dataMap, err := server.dataManager.GetDataMapping()
	if err != nil {
		log.Warnf("从数据库中获取烈度对照表失败:%s", err.Error())
		if err == ErrNotFound {
			log.Warnf("系统将自动初始化烈度对照表")
			GlobalDataMapping = DataMapping{}
			pgamaps := PGAMapArray{}
			pgamaps = append(pgamaps, PGAMapping{22.0, 5})
			pgamaps = append(pgamaps, PGAMapping{45.0, 6})
			pgamaps = append(pgamaps, PGAMapping{90.0, 7})
			pgamaps = append(pgamaps, PGAMapping{178.0, 8})
			pgamaps = append(pgamaps, PGAMapping{354.0, 9})
			pgamaps = append(pgamaps, PGAMapping{708.0, 10})
			pgamaps = append(pgamaps, PGAMapping{1414.0, 11})
			sort.Sort(pgamaps)
			GlobalDataMapping.PGAMap = pgamaps

			simaps := SIMapArray{}
			simaps = append(simaps, SIMapping{3.2, 5})
			simaps = append(simaps, SIMapping{6.7, 6})
			simaps = append(simaps, SIMapping{14.0, 7})
			simaps = append(simaps, SIMapping{29.1, 8})
			simaps = append(simaps, SIMapping{107.8, 9})
			sort.Sort(simaps)
			GlobalDataMapping.SIMap = simaps

			errc := server.dataManager.CreateDataMapping(&GlobalDataMapping)
			if errc != nil {
				log.Warnf("初始化烈度对照表失败:%s", errc.Error())
				return
			}
		} else {
			return
		}
	} else {
		GlobalDataMapping = dataMap
	}
}

//初始化系统配置参数
func InitSystemConfigs() {
	//初始化系统参数
	newCfg, err0 := server.dataManager.GetSystemConfig()
	if err0 != nil {
		log.Warnf("从数据库中获取系统配置参数失败:%s", err0.Error())
		if err0 == ErrNotFound {
			log.Warnf("系统将自动初始化系统基础配置")
			SystemConfigs = SystemConfig{}
			SystemConfigs.UserDefaultPassword = "12345678"
			SystemConfigs.UseGis = true
			SystemConfigs.GisServiceUrl = "http://localhost:8080/geoserver/dqs/wms"
			SystemConfigs.GisServiceParams = ""
			SystemConfigs.GisLayerBasic = "dqs_layers"
			SystemConfigs.GisLayerChina = "china_layer"

			giscfg := GisImageConfig{}
			giscfg.Format = "image/png"
			giscfg.SRS = "EPSG:4326"
			giscfg.Height = "200"
			giscfg.Width = "200"
			giscfg.BBOX = "80.0,20.0,120.0,45.0"
			SystemConfigs.GisImageCfg = giscfg

			//彩信服务
			mms := MMSConfig{}
			mms.MmsEnable = true
			mms.ServiceUrl = "http://sdk3.entinfo.cn:8060/webservice.asmx/mdMmsSend"
			mms.UserNo = ""
			mms.Password = ""
			SystemConfigs.MmsCfg = mms

			mailcfg := MailConfig{}
			mailcfg.MailPort = "25"
			SystemConfigs.MailCfg = mailcfg

			err2 := server.dataManager.AddSystemConfig(&SystemConfigs)
			if err2 != nil {
				log.Warnf("初始化系统参数失败.")
			}
		} else {
			return
		}
	} else {
		SystemConfigs = newCfg
	}
}

//获取烈度值
func getMappingIntensity(a *AlarmInfo) int {
	if GlobalConfig.IntensityMapingData == "PGA" {
		return getMappingIntensityByPGA(a.PGA)
	} else if GlobalConfig.IntensityMapingData == "SI" {
		return getMappingIntensityBySI(a.SI)
	} else {
		return 0
	}
}

//根据PGA获取烈度值
func getMappingIntensityByPGA(pga float32) int {

	lowValue := 0
	dataArray := GlobalDataMapping.PGAMap
	size := len(dataArray)
	if size > 0 {
		lowValue = dataArray[0].Intensity

		for _, v := range dataArray {
			if pga < v.PGA {
				return lowValue
			}
			lowValue = v.Intensity
		}
	}
	return lowValue

}

//根据SI获取烈度值
func getMappingIntensityBySI(si float32) int {
	lowValue := 0
	dataArray := GlobalDataMapping.SIMap
	size := len(dataArray)
	if size > 0 {
		lowValue = dataArray[0].Intensity

		for _, v := range dataArray {
			if si < v.SI {
				return lowValue
			}
			lowValue = v.Intensity
		}
	}
	return lowValue
}
