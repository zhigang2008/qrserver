package quickserver

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

const (
	configFile  = "server.xml"
	defaultPort = "7777"
)

type ServerConfig struct {
	XMLName xml.Name `xml:"Server"`
	Port    string
}

func ReadConfigFromFile() ServerConfig {

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
		//config := ServerConfig{defaultPort}
		//return config
	}

	var result ServerConfig
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
	return result
}
