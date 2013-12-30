package quickserver

import (
	//	"bytes"
	"dqs/util"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	log "github.com/cihub/seelog"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

type MmsServiceCode struct {
	XMLName xml.Name `xml:"xml"`
	Code    int      `xml:"string"`
}

func CheckAndSendMms(report Report) {
	//发送条件判断
	needAudit := GlobalConfig.ReportCfg.AuditBeforeSend
	directSendLevel := GlobalConfig.ReportCfg.MinDirectSendLevel

	//不需要审核,则直接发送
	if needAudit == false {
		PrepareMms(report)
	} else {
		//如果级别大于审核级别,则直接发送
		if report.Event.MaxLevel >= directSendLevel {
			PrepareMms(report)
		} else {
			//如果已经通过审核.则发送.
			if report.Verify {
				PrepareMms(report)
			}
		}
	}

}

//准备发送彩信
func PrepareMms(report Report) {

	//存在彩信服务才发送
	mmscfg := SystemConfigs.MmsCfg
	if mmscfg.MmsEnable == true && mmscfg.ServiceUrl != "" && mmscfg.UserNo != "" && mmscfg.Password != "" {
		num, userlist := getMMSReceiver()
		//有多于1个的接收人才发送
		if num > 0 {
			sendMms(report, userlist)
		}
	}
}

//发送彩信
func sendMms(report Report, users string) {
	//彩信服务账户信息
	sn := SystemConfigs.MmsCfg.UserNo
	password := SystemConfigs.MmsCfg.Password
	pwd := util.GetMd5Hex(sn + password)
	//彩信图片地址
	dir := GlobalConfig.FileConfig.ReportFileDir

	title := fmt.Sprintf("地震速报报警%s", report.Summary.EventTime)
	mmsText := fmt.Sprintf("事件时间:%s", report.Summary.EventTime)
	mmsText += fmt.Sprintf("\n报警数量:%d", report.Summary.AlarmCount)
	mmsText += fmt.Sprintf("\n报警统计:%s", report.Summary.Brief)
	if report.Event.IsConfirm {
		mmsText += fmt.Sprintf("\n地震数据:%s", report.Summary.QuakeInfo)
	}

	//发送内容
	ct := ""
	ct += "1_1.txt," + base64.StdEncoding.EncodeToString([]byte(mmsText)) + ";"
	ct += "1_2.jpg,"
	fc, err := ioutil.ReadFile(filepath.Join(dir, report.ImageFile))
	if err == nil {
		ct += base64.StdEncoding.EncodeToString(fc)
	}

	v := url.Values{}
	v.Add("sn", sn)
	v.Add("pwd", strings.ToUpper(pwd))
	v.Add("title", title)
	v.Add("mobile", users)
	v.Add("content", ct)
	v.Add("stime", "")

	//发送
	r, err := http.PostForm("http://sdk3.entinfo.cn:8060/webservice.asmx/mdMmsSend", v)
	if err != nil {
		log.Warnf("调用彩信接口出错:%s\n", err.Error())
		return
	}
	response, err2 := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err2 != nil {
		log.Warnf("解析彩信返回数据出错:%s\n", err2.Error())
		return
	}
	backcode := MmsServiceCode{}
	err = xml.Unmarshal(response, backcode)
	if err == nil {
		if backcode.Code < 0 || backcode.Code == 1 {
			log.Errorf("彩信发送失败.code=%d", backcode.Code)
		} else {
			log.Infof("无法解析彩信返回信息:%s", err.Error())
		}
	}
}

//获取系统内可接收彩信的账号
func getMMSReceiver() (int, string) {
	receivers := ""
	ulist := []string{}

	users, err := server.dataManager.GetValidUsers()
	if err != nil {
		log.Errorf("获取彩信接收用户列表失败,停止发送彩信.%s", err.Error())
	}

	for _, v := range users {
		if v.ReportSet.ReportPhone && v.Mobile != "" {
			ulist = append(ulist, v.Mobile)
		}
	}

	ucount := len(ulist)
	for k, v := range ulist {
		receivers += fmt.Sprintf("%s", v)
		if k < (ucount - 1) {
			receivers += ","
		}
	}

	return ucount, receivers
}
