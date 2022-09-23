//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package cmd

import (
	"github.com/LambdaTest/photon/config"
	"github.com/google/wire"
)

func InitializeApp(cfg *config.Config) (application, error) {
	panic(wire.Build(loggerSet, scmClientSet, storeSet, kafkaSet, serverSet))
}
