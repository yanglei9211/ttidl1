//go:build wireinject

package model

import (
	"github.com/google/wire"
	"ttidl1/biz/dal/db"
	"ttidl1/conf"
)

func GetItemDbModel() DbItemModel {
	panic(wire.Build(
		conf.NewServerConf,
		db.NewMongoClient,
		NewDbItemModel,
	))
}
