// Package core redis初始化
//@author: zxl

package core

import (
	"github.com/go-redis/redis/v8"
	"go-David/config"
	"go-David/global"
)

var (
	RedisDb *redis.Client
)

func RedisInit() error {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.Db})
	_, err := RedisDb.Ping(global.Ctx).Result()
	if err != nil {
		return err
	}
	global.Log.Infof("[redis]连接成功")
	return nil
}
