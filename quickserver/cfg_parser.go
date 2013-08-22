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
	defaultPort = 7777
	defaultType = "tcp4"
	TCP4        = "tcp4"
	TCP6        = "tcp6"
)

type ServerConfig struct {
	XMLName xml.Name `xml:"Server"`
	Host    string
	Port    int
	Type    string
}

func ReadConfigFromFile() ServerConfig {

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Error(err)
		//config := ServerConfig{defaultPort}
		//return config
	}

	var result ServerConfig
	result.Host = defaultHost
	result.Port = defaultPort
	result.Type = defaultType

	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Error(err)
	}
	log.Info(result)
	return result
}

func CheckConfig(conf ServerConfig) (err error) {
	if conf.Port < 1 || conf.Port > 65535 {
		err = fmt.Errorf("Port must be in (1 ~ 65535")
		return
	}
	if !(conf.Type == TCP4 || conf.Type == TCP6) {
		err = fmt.Errorf("TCP Type only be 'tcp4' or 'tcp6' ")
		return
	}
	return
}
