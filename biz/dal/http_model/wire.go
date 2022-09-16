//go:build wireinject

package http_model

import (
	"github.com/google/wire"
	"ttidl1/conf"
)

func GetGemserverModel() GemserverModel {
	panic(wire.Build(
		conf.NewServerConf,
		NewGemserverModel,
	))
}
