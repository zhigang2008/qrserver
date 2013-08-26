// quickServer.go project main.go
package main

import (
	server "dqs/quickServer"
	"fmt"
	log "github.com/cihub/seelog"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer log.Flush()

	fmt.Println("Server starting ......")
	log.Info("Server starting ......")

	conf, err0 := server.ReadConfigFromFile()
	if err0 != nil {
		//fmt.Printf("配置文件读取失败 : %s\n", err0.Error())
		log.Errorf("[配置文件读取失败]: %s", err0.Error())
		log.Flush()
		os.Exit(1)
	}

	err := server.InitAndStart(conf)
	if err != nil {
		//fmt.Printf("服务启动失败 : %s\n", err.Error())
		log.Errorf("[服务启动失败]: %s", err.Error())
		log.Flush()
		os.Exit(1)
	}

	log.Info("服务启动正常")
}
