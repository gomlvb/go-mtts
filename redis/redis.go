/*****************************************************************************************************
copyright (C),2020-2060,wondershare .Co.,Ltd.

FileName     : redis.go
Author       : Shijh      Version : 1.0    Date: 2020年12月16日
Description  : redis
Version      : 1.0
Function List:
			// Example Config:
			// redis:
			//   addr: 127.0.0.1:6379
			//   password:
			//   db: 0
			//   pool_size: 100
History      :
<author>       <time>             <version>            <desc>
Shijh       2020年12月16日          1.0          build this moudle
******************************************************************************************************/
package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	client *redis.Client
)

type Config struct {
	Addr     string `json:"addr,omitempty"`
	Password string `json:"password,omitempty"`
	Db       int    `json:"db,omitempty"`
	PoolSize int    `json:"pool_size,omitempty"`
}

func Init(cfg Config) {
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("redis connect successfully")
}

func GetConn() *redis.Client {
	return client
}

func Close() {
	if client != nil {
		client.Close()
	}
	logrus.Info("redis connect closed")
}
