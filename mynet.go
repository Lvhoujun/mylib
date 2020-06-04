package mylib

import (
	//"fmt"
	"net"
	"strconv"
)

// ====================================================================
// Api functions
// ====================================================================

func NetTest() {
	StartTcpServer(5000)

}

// ====================================================================
// Internal functions
// ====================================================================

func StartTcpServer(port uint16) {
	Listener, err := net.Listen("tcp", ":"+strconv.Itoa(int(port)))
	DEBUG("开始监听")
	if err != nil {
		ERROR("网络服务启动失败,err:%v", err)
		panic(err)
	}
	DEBUG("服务器启动成功！！！")
	for {
		conn, err := Listener.Accept()
		if err != nil {
			if err != nil {
				ERROR("连接建立失败,err:%v", err)
			}
		}
		go HandleConn(conn)
	}
}

// @doc 连接成功建立处理
func HandleConn(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			ERROR("数据接收失败,err:%v", err)
			return
		}
		DEBUG("received data:%v", string(buf[:len]))
	}
}
