package cmd

import (
	"github.com/LambdaTest/photon/config"
	"github.com/LambdaTest/photon/pkg/lumber"
	"github.com/google/wire"
)

// wire set for loading the logger
var loggerSet = wire.NewSet(
	provideLogger,
)

// provideLogger is a Wire provider function that provides a
// logger, configured from the environment.
func provideLogger(cfg *config.Config) (lumber.Logger, error) {
	return lumber.NewLogger(cfg.LogConfig, cfg.Verbose, lumber.InstanceZapLogger)
}
