package Sprocess

import (
	"chatroom/common/message"
	"chatroom/common/util"
	"chatroom/server/database"
	"fmt"
	"net"
)

type SuserProcess struct {
	Conn net.Conn
	message.User
}

func NewSuerProcess(conn net.Conn) *SuserProcess {
	return &SuserProcess{
		Conn: conn,
	}
}
func (this *SuserProcess) LoginCheck(mes message.Message) {
	// defer this.Conn.Close()
	// 解析数据包
	tf := util.NewTransfer(this.Conn)
	data, err := tf.ParsePacket(mes)
	if err != nil {
		return
	}
	// 转化数据
	loginMes, ok := data.(message.LoginMes)
	if !ok {
		return
	}

	// 定义变量, 用于返回结果给客户端
	var resLoginMes message.ResLoginMes

	// 从数据中获取数据
	rdConn := database.UD.GetConn()
	defer rdConn.Close()
	user, ok := database.UD.SelectById(rdConn, message.DatabaseKey, loginMes.UserId)
	// fmt.Println("ok====", ok)
	// fmt.Println("loginMes=", loginMes)
	// fmt.Println("user=", user)
	if ok {
		// 验证数据是否合法
		if loginMes.UserId == user.UserId && loginMes.UserPwd == user.UserPwd {
			resLoginMes.Code = message.CodeLoginSuccessful
			// 用户登录成功，则将数据存到在线用户管理中
			this.UserId = user.UserId
			this.UserName = user.UserName
			this.UserStatus = true
			user.UserStatus = true
			SM.AddOnlineUser(this)
			resLoginMes.User = message.User{
				UserId:     user.UserId,
				UserName:   user.UserName,
				UserStatus: user.UserStatus,
			}
		} else {
			resLoginMes.Code = message.CodeLoginFailure
		}
	} else {
		resLoginMes.Code = message.CodeHaveNotRegister
	}

	// 封装数据包成mes
	ResMes, err := tf.EncapsulationPacket(message.ResLoginMesType, resLoginMes)
	if err != nil {
		return
	}
	// 将数据包发送回客户端
	err = tf.WritePkg(ResMes)
	if err != nil {
		return
	}
	// 将客户的用户信息返回客户端

	// // 给其他在线用户发送好友状态
	// SM.NotifyOthersUser(user)
	if this.UserStatus {
		SM.NotifyOthersUser()
	}
	go Reading(this.Conn)

}
func (this *SuserProcess) Register(mes message.Message) {
	defer this.Conn.Close()
	tf := util.NewTransfer(this.Conn)
	// 解析数据包
	data, err := tf.ParsePacket(mes)
	if err != nil {
		return
	}
	registerMes, ok := data.(message.RegisterMes)
	if !ok {
		return
	}
	// 定义返回消息类型
	var resRegisterMes message.ResRegisterMes

	// 获取数据库的连接,并读取
	rdConn := database.UD.GetConn()
	defer rdConn.Close()
	_, ok = database.UD.SelectById(rdConn, message.DatabaseKey, registerMes.UserId)
	if ok {
		resRegisterMes.Code = message.CodeRegisterFailure
	} else {
		ok := database.UD.Add(rdConn, message.DatabaseKey, registerMes.UserId, mes.Data)
		if ok {
			resRegisterMes.Code = message.CodeRegisterSuccessful
		} else {
			resRegisterMes.Code = message.CodeRegisterFailure
		}
	}
	// 封装resRegisterMes
	resMes, err := tf.EncapsulationPacket(message.ResRegisterMesType, resRegisterMes)
	if err != nil {
		return
	}
	//发送数据包
	err = tf.WritePkg(resMes)
	if err != nil {
		fmt.Println("server err", err)
		return
	}
	// fmt.Println("返回状态码成功")
}
func (this *SuserProcess) ExitLogin(mes message.Message) {
	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)
	// 解析退出登录的数据包
	dataMes, err := tf.ParsePacket(mes)
	if err != nil {
		return
	}
	exitLoginMes, ok := dataMes.(message.ExitLoginMes)
	if !ok {
		return
	}
	fmt.Println("接收离线请求", exitLoginMes.User)
	// 将要退出登录的用户从在线列表中删除
	delete(SM.OnlineUsers, exitLoginMes.UserId)
	// 向通知其他用户， 此用户已经下线
	// 将数据封装
	resMes, err := tf.EncapsulationPacket(message.ExitLoginMesType, exitLoginMes)
	this.Conn.Close()
	fmt.Println(SM.OnlineUsers)
	// // 向用户发送离线准许
	// err = tf.WritePkg(resMes)
	// if err != nil {
	// 	return
	// }

	// 所有在线客户发送数据包
	for _, sp := range SM.OnlineUsers {
		tfSp := util.NewTransfer(sp.Conn)
		err := tfSp.WritePkg(resMes)
		if err != nil {
			return
		}
		// fmt.Printf("向用户:%s[%d], 发送离线通知成功", sp.UserName, sp.UserId)
	}
}
