package setting

import (
	"ParkNavigate/global"
	"context"
	"time"

	"github.com/go-redis/redis"
)

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Address,
		Password: global.Config.Redis.Password,
		DB:       0,
		PoolSize: global.Config.Redis.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		global.Logger.Errorf("Failed to connect redis %s", global.Config.Redis.Address)
	}

	global.RedisClient = rdb
}
