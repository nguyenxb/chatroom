package Cprocess

import (
	"chatroom/common/message"
	"chatroom/common/util"
	"fmt"
	"net"
)

type CsmsMes struct {
	Conn net.Conn
}

var Csms *CsmsMes

func NewCsmsMes(conn net.Conn) *CsmsMes {
	return &CsmsMes{
		Conn: conn,
	}
}
func (this *CsmsMes) SendDialogToAll() {
	fmt.Println("请你输入想对大家说的话：")
	var dialog string
	fmt.Scanf("%s\n", &dialog)

	// 定义变量，将数据发送给服务器
	var dialogMes message.DialogMes
	// 将发送者的用户信息和发送信息封装到dialogMes中
	dialogMes.User = CurUser
	dialogMes.Dialog = dialog
	// fmt.Println("SendDialog == dialogMes", dialogMes.User)
	// fmt.Println("SendDialog == CurUser", CurUser)
	// fmt.Println("SendDialog == Conn", this.Conn)

	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)

	// 将dialogMes封装
	mes, err := tf.EncapsulationPacket(message.DialogMesType, dialogMes)
	if err != nil {
		return
	}

	// 发送数据包
	err = tf.WritePkg(mes)
	if err != nil {
		return
	}

}
func (this *CsmsMes) SendDialogToAnother() {
	var otherUserId int
	var dialog string
	var dialogOtherUserMes message.DialogOtherUserMes
	fmt.Println("请输入你想私聊的用户Id:")
	fmt.Scanf("%d\n", &otherUserId)
	// 获取用户id 的名字,并判断对方是否在线
	otherUser, ok := Cusers.SearchOnlineUser(otherUserId)
	if !ok {
		fmt.Println("该用户不在线")
	}

	fmt.Printf("请输入你想对用户:%s[%d]说的话：\n", otherUser.UserName, otherUserId)
	fmt.Scanf("%s\n", &dialog)

	// 将数据添加到dialogOtherUserMes 中
	dialogOtherUserMes.Dialog = dialog
	dialogOtherUserMes.OtherUserId = otherUserId
	dialogOtherUserMes.User = CurUser

	// 创建一个Transfer实例
	tf := util.NewTransfer(this.Conn)

	// 将dialogOtherMes 封装成 Message
	mes, err := tf.EncapsulationPacket(message.DialogOtherUserMesType, dialogOtherUserMes)
	if err != nil {
		return
	}
	// 发送数据包给服务端
	err = tf.WritePkg(mes)
	if err != nil {
		return
	}
}
