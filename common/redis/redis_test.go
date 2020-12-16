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
	"log"
	"testing"
)

func TestInit(t *testing.T) {

	cfg := Config{
		Addr:     "10.7.1.184:6379",
		Db:       0,
		PoolSize: 100,
	}
	Init(cfg)

	ctx := context.Background()
	client := GetConn()

	log.Println(client.Set(ctx, "demo", 1, 0).Err())
	log.Println(client.Get(ctx, "demo").Result())

	log.Println(client.MSet(ctx, "num", 2, "num2", 4).Result())
	log.Println(client.MGet(ctx, "num"))
	log.Println(client.MGet(ctx, "num2"))

	log.Println(client.HMSet(ctx, "hm", "1", 2, "3", 4))
	log.Println(client.HMGet(ctx, "hm", "3"))
	log.Println(client.HGetAll(ctx, "hm"))

	log.Println(client.HSet(ctx, "h", "1", 1))
	log.Println(client.HGetAll(ctx, "h"))
}
