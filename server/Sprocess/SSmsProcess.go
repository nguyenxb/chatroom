package Sprocess

import (
	"chatroom/common/message"
	"chatroom/common/util"
	"net"
)

type SSmsProcess struct {
	Conn net.Conn
	User message.User
}

func NewSSmsProcess(conn net.Conn) *SSmsProcess {
	return &SSmsProcess{
		Conn: conn,
	}
}
func (this *SSmsProcess) SendMesToAll(smsMes message.Message) {
	tf := util.NewTransfer(this.Conn)
	// 将信息数据解析
	dataMes, err := tf.ParsePacket(smsMes)
	if err != nil {
		return
	}
	// 将信息转成DialogMes
	dialogMes, ok := dataMes.(message.DialogMes)
	if !ok {
		return
	}
	// // 将发送方的用户信息保存到dialogMes 中
	// dialogMes.User = message.User{
	// 	UserId:     this.User.UserId,
	// 	UserName:   this.User.UserName,
	// 	UserStatus: this.User.UserStatus,
	// }
	// 获取所有在线用户
	users := SM.GetAllOnlineUser()
	// 将数据发送给所有在线客户
	for _, sp := range users {
		sTf := util.NewTransfer(sp.Conn)
		// 封装消息数据
		resMes, err := sTf.EncapsulationPacket(message.DialogMesType, dialogMes)
		if err != nil {
			return
		}
		// 发送消息数据包给所有在线用户
		err = sTf.WritePkg(resMes)
		if err != nil {
			return
		}
	}

}

func (this *SSmsProcess) SendMesToAnother(smsMes message.Message) {
	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)

	// 解析私聊数据包
	dataMes, err := tf.ParsePacket(smsMes)
	if err != nil {
		return
	}
	// fmt.Println("this.COnn", this.Conn)
	// fmt.Println("this.User=", this.User)
	// 进行类型断言
	dialogOtherUserMes, ok := dataMes.(message.DialogOtherUserMes)
	if !ok {
		return
	}
	// 获取接收方的连接数据
	sp, ok := SM.OnlineUsers[dialogOtherUserMes.OtherUserId]
	if !ok {
		return
	}
	// fmt.Printf("用户%s[%d]对你%s[%d]说:%s",)
	// 创建接收方的Transfer实例
	tfSp := util.NewTransfer(sp.Conn)
	// fmt.Println("other Conn", sp.Conn)
	// fmt.Println("Other COnn", sp.User)
	// 将数据封装起来
	resMes, err := tfSp.EncapsulationPacket(message.DialogOtherUserMesType, dialogOtherUserMes)
	if err != nil {
		return
	}
	// 向客户端发送数据
	err = tfSp.WritePkg(resMes)
	if err != nil {
		return
	}
}
