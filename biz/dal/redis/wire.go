//go:build wireinject

package redis

import (
	"github.com/google/wire"
	"ttidl1/conf"
)

//
//func GetMongoClient() MongoClient {
//	panic(wire.Build(
//		conf.NewServerConf,
//		NewMongoClient,
//	))
//}

func GetRedisClient() RedisClient {
	panic(wire.Build(
		conf.NewServerConf,
		NewRedisClient,
	))
}
