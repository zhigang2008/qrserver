/*
用来创建接收数据的服务
通过配置文件设置相关参数,启动服务,接收数据,然后将数据解析后放入存储中
*/
package quickserver

import (
	//	"fmt"
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

//启动Server
func Server(conf ServerConfig) {

	err := CheckConfig(conf)
	checkError(err)

	tcpAddr, err := net.ResolveTCPAddr(conf.Type, conf.Host+":"+strconv.Itoa(conf.Port))
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

//接收数据
func Receiver(conn net.Conn) (err error) {

	buf := make([]byte, RECV_BUF_LEN)
	remoteHost := conn.RemoteAddr().String()

	defer conn.Close()
	for {
		n, err1 := conn.Read(buf)
		switch err1 {
		case nil:
			log.Info("From " + remoteHost + " read data length:" + strconv.Itoa(n))
			log.Info(buf)

		case io.EOF: //当对方断开连接时触发该方法
			log.Warnf("远程终端[%s]已断开连接: %s \n", remoteHost, err1)
			err = err1
			return
		default: //当对方断开连接时触发该方法
			log.Warnf("1远程终端[%s]已断开连接: %s \n", remoteHost, err1)
			err = err1
			return
		}
	}
	return
}

//检查异常
func checkError(err error) {
	if err != nil {
		//fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		log.Errorf("[创建服务失败]: %s", err.Error())
		os.Exit(1)
	}
}
