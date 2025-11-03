package config

import (
	"context"
	"goblog/global"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	// 获取配置信息
	addr := AppConfig.Redis.Addr
	db := AppConfig.Redis.DB
	password := AppConfig.Redis.Password

	// 创建redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       db,
		Password: password,
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect redis client: %v", err)
	}

	// 设置全局变量
	global.RedisDB = redisClient

	log.Println("Configure redis successfully")
}
