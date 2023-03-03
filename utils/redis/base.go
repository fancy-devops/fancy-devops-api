package redis

import (
	"fmt"
	"log"

	"github.com/fancy-devops/fancy-devops-api/utils/setting"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client
var Prefix string

func init() {
	// 参考：https://studygolang.com/articles/27698?fr=sidebar
	sec, err := setting.Cfg.GetSection("redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'redis': %v", err)
	}

	network := sec.Key("NETWORK").String()
	host := sec.Key("HOST").String()
	port := sec.Key("PORT").String()
	pwd := sec.Key("PASSWORD").String()
	db := sec.Key("DB").MustInt()
	Prefix = sec.Key("PREFIX").String()

	options := redis.Options{
		Network:            network,
		Addr:               fmt.Sprintf("%s:%s", host, port),
		Dialer:             nil,
		OnConnect:          nil,
		Password:           pwd,
		DB:                 db,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	}
	// 新建一个client
	RedisClient = redis.NewClient(&options)
}

func CloseRedis() {
	RedisClient.Close()
}
