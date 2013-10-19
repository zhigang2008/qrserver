package main

import (
	"bitbucket.org/kardianos/service"
	httpServer "dqs/httpserver"
	server "dqs/quickserver"
	"fmt"
	log "github.com/cihub/seelog"
	"os"
	"path/filepath"
	"runtime"
)

var serviceLog service.Logger
var currentPath string

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var name = "DQS_Server"
	var displayName = "DQS 地壳所地震速报平台"
	var desc = "这是地震谱烈度速报平台的系统服务"

	var s, err = service.NewService(name, displayName, desc)
	serviceLog = s

	if err != nil {
		fmt.Printf("[%s] 系统服务启动失败: %s", displayName, err)
		return
	}
	currentPath, _ = service.GetExePath()

	if len(os.Args) > 1 {
		var err error
		verb := os.Args[1]
		switch verb {
		case "install":
			err = s.Install()
			if err != nil {
				fmt.Printf("Failed to install: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" installed.\n", displayName)
		case "remove":
			err = s.Remove()
			if err != nil {
				fmt.Printf("Failed to remove: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" removed.\n", displayName)
		case "run":
			startWork()
		case "start":
			err = s.Start()
			if err != nil {
				fmt.Printf("Failed to start: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" started.\n", displayName)
		case "stop":
			err = s.Stop()
			if err != nil {
				fmt.Printf("Failed to stop: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" stopped.\n", displayName)
		}
		return
	}

	err = s.Run(func() error {
		// start
		go startWork()
		return nil

	}, func() error {
		// stop
		go stopWork()
		return nil
	})

	if err != nil {
		s.Error(err.Error())
	}
}

func startWork() {
	dir, err00 := filepath.Split(currentPath)
	fmt.Printf("dir=%s;file=%s\n", dir, err00)
	os.Chdir(dir)
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
	}

	//启动监听服务
	log.Info("启动监听服务...")
	go server.InitAndStart(conf)
	/*if err != nil {
		//fmt.Printf("服务启动失败 : %s\n", err.Error())
		log.Errorf("[服务启动失败]: %s", err.Error())
		log.Flush()
	}
	*/
	log.Info("服务启动正常")

	//是否启动HTTP console
	if conf.HttpServerEnable == true {
		go httpServer.StartHttp(dir)
	}
	select {}
}

//停止服务
func stopWork() {
	httpServer.Close()
	server.Stop()
	log.Info()
	serviceLog.Info("I'm Stopping!")
}
