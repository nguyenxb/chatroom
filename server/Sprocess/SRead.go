package Sprocess

import (
	"chatroom/common/message"
	"chatroom/common/util"
	"net"
)

func Reading(conn net.Conn) {
	for {
		// 首先创建一个Transfer实例
		tf := util.NewTransfer(conn)
		// 接收来自客户端的数据包
		mes, err := tf.ReadPkg()
		if err != nil {
			return
		}
		// 根据不同类型的数据包进行分配给不同的方法
		switch mes.Type {
		case message.DialogMesType:
			// 处理群发信息
			NewSSmsProcess(conn).SendMesToAll(mes)
		case message.DialogOtherUserMesType:
			// 处理私聊信息
			NewSSmsProcess(conn).SendMesToAnother(mes)
		case message.ExitLoginMesType:
			// 处理私聊信息
			NewSuerProcess(conn).ExitLogin(mes)
		default:
			// 处理无效信息
		}
	}
}
