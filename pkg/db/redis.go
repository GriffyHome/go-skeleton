package db

import (
	"fmt"

	"github.com/GriffyHome/go-skeleton/pkg/config"
	"github.com/GriffyHome/go-skeleton/pkg/constants/log_constants/errorLogs"

	"github.com/go-redis/redis"
)

func NewRedisConnection() (*redis.Client, error) {
	host, password := config.GetRedisConnectionDetails()
	redisConn := redis.NewClient(&redis.Options{
		Addr:     host,     // Redis server address
		Password: password, // Redis password (leave empty if no password is set)
		DB:       0,        // Redis database index
	})

	// Ping the Redis server to check the connection
	_, err := redisConn.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf(errorlogs.RedisConnectionError, err.Error())
	}

	return redisConn, nil
}
