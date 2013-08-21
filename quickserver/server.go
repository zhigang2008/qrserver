// server.go
package quickserver

import (
	"fmt"
	log "github.com/cihub/seelog"
	"io"
	"net"
	"os"
	"strconv"
)

const (
	RECV_BUF_LEN = 1024
)

func Start() {
	conf := ReadConfigFromFile()
	Server(conf)
}

func Server(conf ServerConfig) {

	err := CheckConfig(conf)
	checkError(err)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":"+strconv.Itoa(conf.Port))
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Warn("接受请求失败:" + err.Error())
			continue
		}
		go Receiver(conn)
	}
}

func Receiver(conn net.Conn) (err error) {

	buf := make([]byte, RECV_BUF_LEN)
	defer conn.Close()
	for {
		n, err1 := conn.Read(buf)
		switch err1 {
		case nil:
			fmt.Println("read length:" + strconv.Itoa(n))
			fmt.Println(buf)

		case io.EOF: //当对方断开连接时触发该方法
			fmt.Printf("Warning: End of data: %s \n", err1)
			err = err1
			return
		default: //当对方断开连接时触发该方法
			fmt.Printf("Error: Reading data: %s \n", err1)
			err = err1
			return
		}
	}
	return
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		log.Error("创建服务失败")
		os.Exit(1)
	}
}
