//go:build wireinject

package db

import (
	"github.com/google/wire"
	"ttidl1/conf"
)

func GetMongoClient() MongoClient {
	panic(wire.Build(
		conf.NewServerConf,
		NewMongoClient,
	))
}
