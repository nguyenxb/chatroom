package Cprocess

import (
	"chatroom/common/message"
	"chatroom/common/util"
	"fmt"
	"net"
)

type CuserProcess struct {
	Conn net.Conn
}

func NewCuserProcess() *CuserProcess {
	return &CuserProcess{}
}

func (this *CuserProcess) Login(userId int, userPwd string) (err error) {
	// 连接到服务器
	this.Conn, err = net.Dial("tcp", message.ServerHost)
	// fmt.Println("address=", this.Conn.LocalAddr())
	if err != nil {
		return
	}
	// 延时关闭连接
	defer this.Conn.Close()
	// 创建登录消息实例
	loginMes := message.LoginMes{
		User: message.User{
			UserId:  userId,
			UserPwd: userPwd,
		},
	}
	//  创键Transfer实例
	tf := util.NewTransfer(this.Conn)
	// 对数据进行封装
	mes, err := tf.EncapsulationPacket(message.LoginMesType, loginMes)
	if err != nil {
		return
	}
	// 向服务端发送数据包
	tf.WritePkg(mes)
	// 接收服务端的响应
	resMes, err := tf.ReadPkg()
	if err != nil {
		return
	}
	// 解析服务端发回的数据包
	data, err := tf.ParsePacket(resMes)
	if err != nil {
		return
	}
	// 对data 进行类型转换
	resLoginMes, ok := data.(message.ResLoginMes)
	if !ok {
		return
	}
	if resLoginMes.Code == message.CodeLoginSuccessful {
		fmt.Println("登录成功")
		// fmt.Println("conn", this.Conn)
		CurUser = resLoginMes.User

		// 创建消息实例
		Csms = NewCsmsMes(this.Conn)
		// 如果登录成功,则显示登录界面,并启动协程来接收服务端发来的数据
		// 创建一个Cserver 实例
		// cs := NewCserver(this.Conn)
		// var count int
		// 开启协程
		// cs.ServerProcessMes()
		// go cs.ServerProcessMes()
		go ServerProcessMes(this.Conn)
		for {
			ShowLoginInterface()
			// cs.ShowLoginInterface()
			// if count == 0 {
			// 	count++
			// 	// go cs.ServerProcessMes()

			// }
		}

	} else if resLoginMes.Code == message.CodeLoginFailure {
		fmt.Println("登录失败")
	} else if resLoginMes.Code == message.CodeHaveNotRegister {
		fmt.Println("未注册用户")
	}

	return
}

func (this *CuserProcess) Register(userId int, userPwd, userName string) (err error) {
	fmt.Println("用户注册")
	// 连接服务器
	this.Conn, err = net.Dial("tcp", message.ServerHost)
	if err != nil {
		return
	}
	// 延时关闭连接
	defer this.Conn.Close()

	// 创建RegisterMes实例
	registerMes := message.RegisterMes{
		User: message.User{
			UserId:   userId,
			UserPwd:  userPwd,
			UserName: userName,
		},
	}
	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)
	// 将registerMes 封装
	mes, err := tf.EncapsulationPacket(message.RegisterMesType, registerMes)
	if err != nil {
		return
	}
	// fmt.Println("Register    90 mes.type=", mes.Type)
	// 发送数据包
	err = tf.WritePkg(mes)
	if err != nil {
		return
	}

	//  接收服务端返回的数据包
	resMes, err := tf.ReadPkg()
	if err != nil {
		return
	}
	// 解析数据包
	data, err := tf.ParsePacket(resMes)
	if err != nil {
		return
	}
	resRisterMes, ok := data.(message.ResRegisterMes)
	if !ok {
		return
	}
	if resRisterMes.Code == message.CodeRegisterSuccessful {
		fmt.Println("注册成功")
	} else {
		fmt.Println("注册失败")

	}
	return

}
func (this *CuserProcess) ExitLogin(userId int) {
	// 创建Transfer 实例
	tf := util.NewTransfer(this.Conn)
	var ExitLoginMes message.ExitLoginMes
	ExitLoginMes.User = CurUser

	// 将退出登录信息封装到mes
	mes, err := tf.EncapsulationPacket(message.ExitLoginMesType, ExitLoginMes)
	if err != nil {
		return
	}
	// 将数据包发送给服务端
	err = tf.WritePkg(mes)
	if err != nil {
		return
	}
	this.Conn.Close()
	// fmt.Println("发送离线请求", ExitLoginMes.User)
	// fmt.Println("发送离线请求", CurUser)
	// // 接收离线准许
	// resMes, err := tf.ReadPkg()
	// // 解析数据包
	// dataMes, err := tf.ParsePacket(resMes)
	// if err != nil {
	// 	return
	// }
	// _, ok := dataMes.(message.ExitLoginMes)
	// if !ok {
	// 	return
	// }
}
