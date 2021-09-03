package database

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type UserDao struct {
	Pool *redis.Pool
}

var UD *UserDao

// 获取连接
func (this *UserDao) GetConn() (conn redis.Conn) {
	return this.Pool.Get()
}

// 对数据库进行增加数据
func (this *UserDao) Add(conn redis.Conn, key string, field int, value string) bool {
	_, err := conn.Do("HSet", key, field, value)
	if err != nil {
		return false
	}
	return true
}

// 删除数据
func (this *UserDao) Del(conn redis.Conn, key int, field string) bool {
	_, err := conn.Do("HDel", key, field)
	if err != nil {
		return false
	}
	return true
}

// 修改数据
func (this *UserDao) Modify(conn redis.Conn, key string, field int, value string) bool {

	return this.Add(conn, key, field, value)
}

// 查询数据
func (this *UserDao) SelectById(conn redis.Conn, key string, filed int) (data message.User, flag bool) {
	res, _ := conn.Do("HGet", key, filed)
	if res != nil {
		flag = true
		resByte := res.([]byte)
		resStr := string(resByte)
		fmt.Println("resByte=", resByte)
		fmt.Println("resStr=", resStr)
		// 将 res 反序列化
		err := json.Unmarshal([]byte(resStr), &data)
		if err != nil {
			return
		}
		fmt.Println("data==", data)
		return data, flag
	}
	fmt.Println("res==", res)
	flag = false
	return
}
