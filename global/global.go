package global

import (
	"ParkNavigate/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var (
	Config      *config.Config
	Logger      *logrus.Logger
	RedisClient *redis.Client
)
