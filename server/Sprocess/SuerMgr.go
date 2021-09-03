package Sprocess

import (
	"chatroom/common/message"
	"chatroom/common/util"
)

// 用于维护在线用户
type SonlineUerMgr struct {
	OnlineUsers map[int]*SuserProcess
}

var SM *SonlineUerMgr

func NewSonlineUserMgr() {
	SM = &SonlineUerMgr{
		OnlineUsers: make(map[int]*SuserProcess, 1024),
	}
}

// 对在线用户进行增删改查操作
func (this *SonlineUerMgr) AddOnlineUser(sp *SuserProcess) {
	SM.OnlineUsers[sp.UserId] = sp
}
func (this *SonlineUerMgr) DelOnlineUser(sp *SuserProcess) {
	delete(SM.OnlineUsers, sp.UserId)
}
func (this *SonlineUerMgr) ModifyOnlineUser(sp *SuserProcess) {
	this.AddOnlineUser(sp)
}
func (this *SonlineUerMgr) SeletcOnlineUser(sp *SuserProcess) (value *SuserProcess, ok bool) {
	value, ok = this.OnlineUsers[sp.UserId]
	return
}
func (this *SonlineUerMgr) GetAllOnlineUser() (users map[int]*SuserProcess) {
	return this.OnlineUsers
}
func (this *SonlineUerMgr) NotifyOthersUser() {
	// fmt.Println("onlineMes=", this.OnlineUsers)
	// 遍历所有在线用户,并向用户通知状态信息
	for _, sp := range this.OnlineUsers {
		// 获取当前登录的用户信息
		statusMes := message.ResStatusMes{
			User: message.User{
				UserId:     sp.UserId,
				UserName:   sp.UserName,
				UserStatus: sp.UserStatus,
			},
		}
		for _, sp1 := range this.OnlineUsers {
			// 创建Transfer实例
			tf := util.NewTransfer(sp1.Conn)
			// 封装数据到mes
			mes, err := tf.EncapsulationPacket(message.ResStatusMesType, statusMes)
			if err != nil {
				return
			}
			// 通知每一个用户
			err = tf.WritePkg(mes)
			if err != nil {
				return
			}
		}
	}

}
