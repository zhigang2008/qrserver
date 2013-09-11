// quickServer.go project main.go
package main

import (
	httpServer "dqs/httpserver"
	server "dqs/quickserver"
	"fmt"
	log "github.com/cihub/seelog"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//配置日志
	logger, e := log.LoggerFromConfigAsFile("seelog.xml")
	if e != nil {
		fmt.Println("读取日志配置出错:" + e.Error())
	}
	log.ReplaceLogger(logger)
	defer log.Flush()

	fmt.Println("Server starting ......")
	log.Info("Server starting ......")

	//读取配置文件
	conf, err0 := server.ReadConfigFromFile()
	if err0 != nil {
		//fmt.Printf("配置文件读取失败 : %s\n", err0.Error())
		log.Errorf("[配置文件读取失败]: %s", err0.Error())
		log.Flush()
		os.Exit(1)
	}

	//是否启动HTTP console
	if conf.HttpServerEnable == true {
		go httpServer.StartHttp()
	}

	//启动监听服务
	log.Info("启动监听服务...")
	err := server.InitAndStart(conf)
	if err != nil {
		//fmt.Printf("服务启动失败 : %s\n", err.Error())
		log.Errorf("[服务启动失败]: %s", err.Error())
		log.Flush()
		os.Exit(1)
	}
	log.Info("服务启动正常")

	defer server.Stop()
}
