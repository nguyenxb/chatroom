package Sprocess

import (
	"chatroom/common/message"
	"chatroom/common/util"
	"fmt"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) CreatProcess() {
	// 首先创建一个Transfer实例
	tf := util.NewTransfer(this.Conn)
	// fmt.Println("conn=", this.Conn)
	// 接收来自客户端的数据包
	mes, err := tf.ReadPkg()
	if err != nil {
		return
	}
	// fmt.Println("adsads")
	// 根据不同类型的数据包进行分配给不同的方法
	// fmt.Println("CreatProcess mes.Type=", mes.Type)
	// fmt.Println("mes Data=", mes.Data)
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录信息
		fmt.Println("处理登录信息")
		NewSuerProcess(this.Conn).LoginCheck(mes)

	case message.RegisterMesType:
		// 处理注册信息
		fmt.Println("处理注册信息")
		NewSuerProcess(this.Conn).Register(mes)
	case message.DialogMesType:
		// 处理注册信息
		fmt.Println("处理用户的消息")
		NewSSmsProcess(this.Conn).SendMesToAll(mes)
	default:
		// 处理无效信息
	}

}
