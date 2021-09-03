package message

const (
	LoginMesType           = "LoginMes"
	ResLoginMesType        = "LoginResMes"
	RegisterMesType        = "RegisterMes"
	ResRegisterMesType     = "ResRegisterMes"
	DialogMesType          = "DialogMes"
	ResStatusMesType       = "ResStatusMes"
	DialogOtherUserMesType = "DialogOtherUserMes"
	ExitLoginMesType       = "ExitLoginMes"
)
const (
	ServerHost = "192.168.232.245:9999"
	// ServerHost = "169.254.237.236:9999"
	// ServerHost   = "localhost:9999"
	DatabaseHost = "localhost:6379"
)

// 数据库key值
const (
	DatabaseKey = "users"
)

// 状态码协议
const (
	CodeLoginSuccessful    = 102
	CodeLoginFailure       = 105
	CodeRegisterSuccessful = 202
	CodeRegisterFailure    = 205
	CodeHaveNotRegister    = 400
)

type User struct {
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus bool   `json:"userStatus"` // 在线就是true
}
type MesType string

// 定义消息类型
type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` //数据的内容
}

// 协议：100 登录成功 ,105 密码错误,200 注册成功,300 用户存在,400 用户不存在
type ResMes struct {
	Code int `json:"code"`
}

// 登录
type LoginMes struct {
	User
}

type ResLoginMes struct {
	ResMes
	User
}

type ExitLoginMes struct {
	User
}

// 注册
type RegisterMes struct {
	User
}
type ResRegisterMes struct {
	ResMes
}

// 状态
type ResStatusMes struct {
	User
}

// 消息
type DialogMes struct {
	User
	Dialog string `json:"dialog"`
}
type DialogOtherUserMes struct {
	User
	OtherUserId int
	Dialog      string `json:"dialog"`
}
