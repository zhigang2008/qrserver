package quickserver

import (
	//"errors"
	"fmt"
	log "github.com/cihub/seelog"
	//"net"
	"dqs/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

const (
	ErrorImage = "notfound.jpg"
)

//延时进行速报构建
func DelayGenerateReport(event *Event) {
	//延时处理.
	delaytime := GlobalConfig.ReportCfg.SleepTime
	time.Sleep(time.Duration(delaytime) * time.Minute)

	newEvent, err := server.dataManager.GetEventById(event.EventId)
	if err != nil {
		log.Warnf("速报处理-查询最新的时间信息失败:%s", err.Error())
		newEvent = *event
	}

	//log.Infof("级别判断%d-%d", event.MaxLevel, GlobalConfig.ReportCfg.ReportLevel)
	if event.MaxLevel >= GlobalConfig.ReportCfg.ReportLevel {
		//制作速报信息
		GenerateReport(&newEvent)

	}

}

//生成速报
func GenerateReport(event *Event) {
	log.Infof("准备生成速报...")

	summary, valStr, dataSize := generateReportSummary(event)
	report := new(Report)
	report.ReportId = util.GUID()
	report.GenerateTime = time.Now()
	report.EventId = event.EventId
	report.Event = *event
	report.Summary = summary
	report.Verify = false
	report.Sended = false
	report.Valid = true

	imgfile := generateReportMap(report.ReportId, valStr, dataSize)
	report.ImageFile = imgfile

	//保存速报
	err := server.dataManager.ReportSave(report)
	if err != nil {
		log.Warnf("速报保存失败:%s", err.Error())
	} else {
		log.Info("速报已生成.")
	}

	//发送彩信
	CheckAndSendNotify(*report)
}

//生成概要信息
func generateReportSummary(event *Event) (ReportSummary, string, int) {

	sumary := ReportSummary{}
	//基本信息
	sumary.EventTime = event.EventTimeStr
	sumary.AlarmCount = event.AlarmCount

	valStr := ""
	dataSize := 0
	maxLevel := 0
	maxPointAlarm := AlarmInfo{}

	cs := make(map[int]int)
	alarms, err := server.dataManager.GetAlarmsByEvent(event)
	if err == nil {
		dataSize = len(*alarms)
		for k, v := range *alarms {
			i := v.Intensity

			val, ok := cs[i]
			if ok {
				cs[i] = val + 1
			} else {
				cs[i] = 1
			}
			//判断最高地点
			if i > maxLevel {
				maxLevel = i
				maxPointAlarm = v
			}
			//构造value string
			if k < (dataSize - 1) {
				valStr += fmt.Sprintf("%f-%f-%d,", v.Longitude, v.Latitude, v.Intensity)
			} else {
				valStr += fmt.Sprintf("%f-%f-%d", v.Longitude, v.Latitude, v.Intensity)
			}
		}
	} else {
		log.Infof("获取震情事件的所有报警信息失败:%s", err.Error())
	}

	alarmSumary := ""
	for k, v := range cs {
		alarmSumary += fmt.Sprintf("%d度-%d; ", k, v)

	}
	//添加详细信息
	device, _ := server.dataManager.GetDeviceById(maxPointAlarm.SensorId)
	alarmSumary += fmt.Sprintf("\n最强仪器烈度(%d),来自%s-%s,经纬度[%f,%f]", maxLevel, device.SensorId, device.SetParams.SiteName, maxPointAlarm.Longitude, maxPointAlarm.Latitude)

	sumary.Brief = alarmSumary

	//实际地震信息
	if event.IsConfirm {
		signal := event.Signal
		stime := signal.Time.Format(CommonTimeLayout)
		sumary.QuakeInfo = fmt.Sprintf("%s 发生于[%f,%f] 震级%d级", stime, signal.Longitude, signal.Latitude, signal.Level)
	}
	return sumary, valStr, dataSize
}

//生成速报的图片信息
func generateReportMap(reportid string, valstr string, simplesize int) string {

	v := url.Values{}
	v.Add("SERVICE", "AMS")
	v.Add("REQUEST", "GetMap")
	v.Add("STYLES", "")
	v.Add("TRANSPARENT", "TRUE")
	v.Add("VERSION", "1.1.1")
	v.Add("LAYERS", SystemConfigs.GisLayerBasic)

	v.Add("SRS", SystemConfigs.GisImageCfg.SRS)
	v.Add("FORMAT", SystemConfigs.GisImageCfg.Format)
	v.Add("HEIGHT", SystemConfigs.GisImageCfg.Height)
	v.Add("WIDTH", SystemConfigs.GisImageCfg.Width)
	v.Add("BBOX", SystemConfigs.GisImageCfg.BBOX)

	v.Add("INTERPOLATION_STRATEGY", "2")
	v.Add("INTERVALS[]", "0,1,2,3,4,5,6,7,8,9,10,11,12")
	v.Add("INTERVALS_COLORS[]", "0xffffff01,0xe9ffd255,0xe9ffd299,0xe9ffd2aa,0xe9ffd2bb,0xe9ffd2cc,0xe9ffbeff,0xa3ff73ff,0x98e600ff,0xffcc00ff,0xff6600ff,0xa80000ff,0x343434ff")
	v.Add("RADIUS", "50")
	v.Add("RENDERER_TYPE", "2")
	v.Add("SIMPLIFY_METHOD", "1")

	v.Add("DATA_ARRAY", valstr)
	v.Add("SIMPLIFY_SIZE", fmt.Sprintf("%d", simplesize))

	r, err := http.PostForm(SystemConfigs.GisInnerServiceUrl, v)
	if err != nil {
		log.Errorf("生成图片调用远程GIS接口出错:%s\n", err.Error())
		return ErrorImage
	}
	response, err2 := ioutil.ReadAll(r.Body)
	if err2 != nil {
		log.Errorf("解析GIS图片出错:%s\n", err2.Error())
		return ErrorImage
	}
	imagename := reportid + ".jpg"
	/*if SystemConfigs.GisImageCfg.Format == "image/jpeg" {
		imagename = reportid + ".jpg"
	}
	*/
	dir := GlobalConfig.FileConfig.ReportFileDir
	//先判断目录是否存在,不存在则创建
	if util.IsDirExist(dir) == false {
		os.MkdirAll(dir, 0777)
	}

	err = ioutil.WriteFile(filepath.Join(dir, imagename), response, 0777)
	if err != nil {
		log.Errorf("等值线图片保存出错:%s\n", err.Error())
		return ErrorImage
	}

	return imagename
}
