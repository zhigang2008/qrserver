/*
用来创建接收数据的服务
通过配置文件设置相关参数,启动服务,接收数据,然后将数据解析后放入存储中
*/
package quickserver

import (
	"fmt"
	log "github.com/cihub/seelog"
	"io"
	"net"
	//"runtime"
	"strconv"
	"time"
)

const (
	RECV_BUF_LEN = 1024
)

var (
	server         *Server
	ClientNum      int = 0
	ConnecitonPool map[string]*net.Conn
)

//服务器对象结构
type Server struct {
	tcpType     string
	serverHost  string
	serverPost  string
	dataManager *DataManager
}

//服务器初始化并启动监听
func InitAndStart(conf ServerConfig) (err error) {
	server = &Server{
		serverHost: conf.Host,
		serverPost: strconv.Itoa(conf.Port),
		tcpType:    conf.Type,
	}
	server.dataManager, err = InitDatabase(conf.Database)
	if err != nil {
		return
	}

	//会话连接池
	ConnecitonPool = make(map[string]*net.Conn)

	return server.start()

}

//启动Server
func (server *Server) start() (err error) {

	tcpAddr, err := net.ResolveTCPAddr(server.tcpType, server.serverHost+":"+server.serverPost)
	if err != nil {
		log.Errorf("无法解析监听地址:%s", err.Error())
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Errorf("监听失败:%s", err.Error())
		return
	}

	defer listener.Close()

	var tempDelay time.Duration

	//不断监听
	for {
		conn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Errorf("Server :Accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			log.Warn("接受请求失败:" + err.Error())
			continue
		}

		tempDelay = 0

		//
		go receiver(server, conn)
	}

}

//接收数据处理数据
func receiver(server *Server, conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, RECV_BUF_LEN)
	var deviceId string = ""
	var sessionSaved = false

	remoteHost := conn.RemoteAddr().String()
	log.Infof("终端建立连接:[%s]", remoteHost)

	ClientNum++
	log.Infof("当前建立连接的设备:%d", ClientNum)

	//获取一个数据处理器
	dataProcessor := NewDataProcessor(server.dataManager)

	for {
		n, err1 := conn.Read(buf)
		switch err1 {
		case nil:
			log.Infof("From [%s] read data length:%d ", remoteHost, n)
			if n < 11 {
				log.Infof("无效数据:数据长度过短")
				continue
			}
			deviceId = string(buf[0:10])

			//存储会话connection
			if sessionSaved == false {
				saveConnection(deviceId, &conn)
				sessionSaved = true
			}

			//数据处理
			log.Info(string(buf[0:n]))
			dataProcessor.DataProcess(buf[0:n], remoteHost)

		case io.EOF: //当对方断开连接时触发该方法
			log.Warnf("远程终端[%s]已断开连接: %s \n", remoteHost, err1)
			ClientNum--
			//设备下线
			if deviceId != "" {
				removeConnection(deviceId)
				dataProcessor.DeviceOffline(deviceId)
			}

			log.Infof("当前建立连接的设备:%d", ClientNum)
			return
		default: //断开连接
			log.Warnf("远程终端[%s]读取失败: %s \n", remoteHost, err1)
			ClientNum--
			//设备下线
			if deviceId != "" {
				removeConnection(deviceId)
				dataProcessor.DeviceOffline(deviceId)
			}
			log.Infof("当前建立连接的设备:%d", ClientNum)
			return
		}

	}

}

//会话连接池
func saveConnection(id string, conn *net.Conn) {
	ConnecitonPool[id] = conn
	for k, v := range ConnecitonPool {
		fmt.Printf("[%s]的会话:%x \n", k, v)
	}
}

//会话连接池移除会话
func removeConnection(id string) {
	fmt.Printf("remove id=%s\n", id)
	delete(ConnecitonPool, id)
	for k, v := range ConnecitonPool {
		fmt.Printf("[%s]的会话:%x \n", k, v)
	}
}

//停止服务
func Stop() {
	server.dataManager.DataClose()
	DllUtil.FreeDLL()
	log.Warn("Server Stop")
}
