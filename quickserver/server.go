// server.go
package quickserver

import (
	"fmt"
	log "github.com/cihub/seelog"
	"net"
	"os"
	"time"
)

func Start() {
	conf := ReadConfigFromFile()
	Server(conf)
}

func Server(conf ServerConfig) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":"+conf.Port)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Warn("接受请求失败:" + err.Error())
			continue
		}
		go handleData(conn)
	}
}

func handleData(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		log.Error("创建服务失败")
		os.Exit(1)
	}
}
