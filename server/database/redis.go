package database

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

func InitPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	UD = &UserDao{
		Pool: &redis.Pool{
			
			MaxIdle:     maxIdle,
			MaxActive:   maxActive,
			IdleTimeout: idleTimeout,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", address)
			},
		},
	}
}
