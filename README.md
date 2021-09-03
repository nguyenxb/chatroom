```text
# chatroom
## 基于go语言tcp的聊天室
master 
├─client 客户端
│  ├─Cmain
│  │      main.go 程序入口
│  │
│  └─CProcess 处理进程入口
│          CsmsMes.go 处理消息收发
│          CuserMgr.go 处理用户管理
│          CuserProcess.go 处理用户登录,注销请求
│          server.go 用户服务管理
│
├─common 公共包
│  ├─message
│  │      message.go  服务端与客户端的通讯协议
│  │
│  └─util
│          util.go 服务端与客户端的公共工具
│
└─server 服务端
    ├─database  
    │      redis.go 连接redis
    │      userDao.go 用户数据的读取
    │
    ├─Smain
    │      main.go 程序入口
    │
    └─Sprocess
            Sprocessor.go 根据客户端发送的不同请求,使用不同的响应方式给客户端
            SRead.go   接收客户端数据的协程
            SSmsProcess.go  响应客户端,给客户端发送数据
            SuerMgr.go 用户管理器
            SuserProcess.go 处理客户端的用户业务

```
