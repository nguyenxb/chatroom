package Cprocess

import (
	"chatroom/common/message"
	"chatroom/common/util"
	"fmt"
	"net"
	"os"
)

type Cserver struct {
	// Conn net.Conn
}

// func NewCserver(conn net.Conn) *Cserver {
// 	return &Cserver{
// 		Conn: conn,
// 	}
// }

func ShowLoginInterface() {
	// func (this *Cserver) ShowLoginInterface() {
	var key int
	fmt.Println("----------------登录成功-------------")
	fmt.Println("\t\t\t 1 显示在线用户列表")
	fmt.Println("\t\t\t 2 发送消息")
	fmt.Println("\t\t\t 3 消息列表")
	fmt.Println("\t\t\t 4 退出登录")
	fmt.Println("\t\t\t 请选择(1-4):")
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("1 显示在线用户列表")
		Cusers.OutputAllOnlineUser()
	case 2:
		fmt.Println("2 发送消息")
		fmt.Println("你要给谁发送消息,请选择(1-2)：")
		fmt.Println("1 群发")
		fmt.Println("2 私聊")
		fmt.Println("请选择(1-2):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			Csms.SendDialogToAll()
		case 2:
			Csms.SendDialogToAnother()
		default:
			fmt.Println("没有该选项")
		}
	case 3:
		fmt.Println("3 消息列表")
		fmt.Println("该功能还没有实现,请联系后台维护人员")
	case 4:
		fmt.Println(" 4 退出登录")
		cp := NewCuserProcess()
		cp.Conn = Csms.Conn
		cp.ExitLogin(CurUser.UserId)
		os.Exit(0)
	default:
		fmt.Println("无此功能")
		break
	}
}

func ServerProcessMes(conn net.Conn) {
	// func (this *Cserver) ServerProcessMes() {
	// 创建transfer实例，不断的读取服务端发来的数据包
	tf := util.NewTransfer(conn)
	// tf := util.NewTransfer(this.Conn)
	for {
		resMes, err := tf.ReadPkg()
		// fmt.Println("resMes=", resMes.Type)
		// fmt.Println("resMes=", resMes.Data)
		if err != nil {
			// fmt.Println("服务器出错啦111")
			return
		}
		//解析数据包
		DataMes, err := tf.ParsePacket(resMes)
		if err != nil {
			// fmt.Println("服务器出错啦")
			return
		}
		// 读取消息类型
		switch resMes.Type {
		case message.ResStatusMesType:

			// 当前在线用户
			resStatusMes, ok := DataMes.(message.ResStatusMes)
			if !ok {
				break
			}
			fmt.Printf("----------用户:%s[%d]上线了--------\n", resStatusMes.UserName, resStatusMes.UserId)
			// 将客户端返回的用户存到客户端维护的map中
			Cusers.OnlineUsers[resStatusMes.UserId] = resStatusMes.User
			// info := fmt.Sprintf("用户ID%d\t用户名%s\t用户状态%s",
			// 	resStatusMes.UserId, resStatusMes.UserName, "在线")
			// // fmt.Println("user status = ", resStatusMes)
			// fmt.Println(info)
		case message.DialogMesType:
			// 将数据转换成CurUserMes
			dialogMes, ok := DataMes.(message.DialogMes)
			if !ok {
				return
			}
			// 将数据输出在控制台
			info := fmt.Sprintf("用户:%s[%d]对大家说:%s",
				dialogMes.UserName, dialogMes.UserId, dialogMes.Dialog)
			fmt.Println(info)
		case message.DialogOtherUserMesType:
			// 将数据转换成CurUserMes
			dialogOtherMes, ok := DataMes.(message.DialogOtherUserMes)
			if !ok {
				return
			}
			// 获取接收放的数据
			otherUser, ok := Cusers.OnlineUsers[dialogOtherMes.OtherUserId]
			if !ok {
				return
			}

			// 将数据输出在控制台
			info := fmt.Sprintf("用户%s[%d]对你%s[%d]说:%s",
				dialogOtherMes.UserName, dialogOtherMes.UserId, otherUser.UserName, dialogOtherMes.OtherUserId, dialogOtherMes.Dialog)
			fmt.Println(info)
		case message.ExitLoginMesType:
			//  类型断言
			exitLoginMes, ok := DataMes.(message.ExitLoginMes)
			if !ok {
				return
			}
			// 将退出登录的客户从在线列表中删除
			Cusers.DelOnlineUser(exitLoginMes.UserId)
			fmt.Printf("----------用户:%s[%d]离线了--------\n", exitLoginMes.UserName, exitLoginMes.UserId)
		default:
			break
		}
	}
}
