// quickServer.go project main.go
package main

import (
	"fmt"
	log "github.com/cihub/seelog"
	"os"
	server "quickserver"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer log.Flush()
	//defer server.DataClose()

	fmt.Println("Server Start !")
	log.Info("Hello World!")

	conf, err := server.ReadConfigFromFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		log.Errorf("[创建服务失败]: %s", err.Error())
		os.Exit(1)
	}
	server.InitDatabase(conf)
	server.DataProcess2()
	//server.InitServer(conf)
	//server.Start()
}
