package mylib

import (
	"syscall"
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
	DEBUG("开始监听[%d]", port)
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
	defer conn.Close()
	Client := conn.RemoteAddr().String()
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		switch err {
		case nil:
			handleMsg(Client, buf)
		case syscall.EAGAIN:
			continue
		default:
			ERROR("数据接收失败,err:%v", err)
			return
		}

		len, err = conn.Write(buf)
		if err != nil {
			ERROR("数据接收失败,err:%v", err)
			return
		}
		DEBUG("send data:%v", string(buf[:len]))
	}
}

func handleMsg(Client string, buf []byte) {
	DEBUG("[%s]:<%s>", Client, buf)
}
