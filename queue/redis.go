package queue

import (
	"github.com/go-redis/redis/v8"
	"github.com/zzjbattlefield/delay_queue/config"
)

var RedisDB *redis.Client

func initRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     config.Setting.Redis.Host,
		Password: config.Setting.Redis.Password,
		DB:       config.Setting.Redis.Db,
	})
}
