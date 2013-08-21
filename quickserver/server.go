// server.go
package quickserver

import (
	"fmt"
	log "github.com/cihub/seelog"
	"net"
	"os"
	"time"
)

const{
	RECV_BUF_LEN = 1024
}

func Start() {
	conf := ReadConfigFromFile()
	Server(conf)
}

func Server(conf ServerConfig) {

	err := CheckConfig(conf)
	checkError(err)

	tcpAddr, err := net.ResolveTCPAddr(conf.Type, ":"+string(conf.Port))
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
			//n, _ := conn.Write(buf[0:n])
			var out LspMsg
			//Decode(b, &out)
			var outout LspMsgBig
			if err := Decode(buf, &outout); err != nil {
				fmt.Println("decode fail: " + err.Error())
			}
			fmt.Println("outout is ", outout)
			fmt.Println("Byte2Int32 is ", BytesToInt32(buf[0:4]))
			fmt.Println("length is ", buf[0:n])
			fmt.Println("length is ", buf[0:4])
			fmt.Println("length is ", BytesToInt8(buf[1:4]))
			out.seq = BytesToInt32(buf[0:4])
			out.protocol = BytesToInt32(buf[4:8])
			out.length = BytesToInt32(buf[8:12])
			out.times = BytesToInt64(buf[12:20])
			out.lens = BytesToInt32(buf[20:24])
			out.lsp = BytesToInt32(buf[24:28])
			bytes := out.bytes[0:20]
			copy(bytes, buf[28:n])
			//out.bytes = &(buf[28:n])
			fmt.Println(out.bytes)
			/*
			   for j := 0; j < 20; j++ {
			       out.bytes[j] = buf[j+28]
			   }
			*/
			fmt.Println("length is ", out)
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
