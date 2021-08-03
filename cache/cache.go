package cache

import (
	"legato_server/env"

	"github.com/Kesci/gostore"
)

var Cache gostore.Store

func ConnectToRedis(){
 option := gostore.StoreOptions{
	RedisHost: env.ENV.RedisHost,
 }
 Cache.Init(&option)
}