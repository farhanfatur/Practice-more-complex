package connectionredis

import (
	"TheLast/model/baseconnection"
	"TheLast/service"
	"log"

	"github.com/gomodule/redigo/redis"
)

type redisStruct struct {
}

var takeConnectionRedis *redis.Pool

// Connection Redis
func (r redisStruct) Connection() {
	takeConnectionRedis = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 1200,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				log.Fatal(err.Error())
			}
			return c, err
		},
	}
}

func GetConnectionRedis() *redis.Pool {
	return takeConnectionRedis
}

func NewRedis() service.ConnectInterface {
	return &redisStruct{}
}

func init() {
	baseconnection.RegisterConnection(NewRedis())
}
