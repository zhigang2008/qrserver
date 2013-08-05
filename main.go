// quickServer.go project main.go
package main

import (
	"fmt"
	log "github.com/cihub/seelog"
	server "quickserver"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer log.Flush()
	fmt.Println("Server Start !")
	log.Info("Hello World!")
	server.Start()

}
