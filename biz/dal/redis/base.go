package redis

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
	"time"
	"ttidl1/conf"
)

type RedisClient struct {
	client *redis.Client
	valid  bool
}

var redisClient RedisClient

func (s *RedisClient) SetCache(k, v string, dur time.Duration) {
	err := s.client.Set(k, v, dur)
	if err != nil {
		hlog.Errorf("redis set:%s, %s error: ", k, v, err.Err())
	}
}

func (s *RedisClient) GetCache(k string) string {
	value := s.client.Get(k)
	return value.String()
}

func NewRedisClient(server conf.ConfigServer) RedisClient {
	if redisClient.valid == true {
		return redisClient
	}
	client := redis.NewClient(&redis.Options{
		Addr:     server.RedisUrl,
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		hlog.Errorf("redis: %s ping error. %s", server.RedisUrl, err.Error())
		panic(fmt.Sprintf("redis: %s ping error", server.RedisUrl))
	}
	hlog.Infof("redis: %s ping result", server.RedisUrl, pong)
	redisClient.client = client
	redisClient.valid = true
	return redisClient
}
