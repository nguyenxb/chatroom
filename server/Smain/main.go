// server
package main

import (
	"chatroom/common/message"
	"chatroom/server/Sprocess"
	"chatroom/server/database"
	"fmt"
	"net"
	"time"
)

func creatProcessor(conn net.Conn) {
	// 延时关闭连接
	// defer conn.Close()
	// 创建一个任务管家，专门用于分配任务
	processor := Sprocess.Processor{
		Conn: conn,
	}
	fmt.Println("客户端连接conn = ", conn)
	processor.CreatProcess()
}

func init() {
	// 初始化线程池
	database.InitPool(message.DatabaseHost, 16, 0, 300*time.Second)
	// 维护一个用户管理集合
	Sprocess.NewSonlineUserMgr()
}

func main() {

	// 1 监听端口 9999
	fmt.Println("服务器在端口 9999 开始监听")
	listen, err := net.Listen("tcp", message.ServerHost)
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}

	// 2 等待获取来自于客户端的连接
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		fmt.Println("address=", conn.LocalAddr())
		if err != nil {
			fmt.Println("listen.Accept err", err)
			return
		}
		// 每连接到一个客户端就启动一个协程为其服务
		go creatProcessor(conn)

	}

}
