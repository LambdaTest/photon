// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/LambdaTest/photon/config"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeApp(cfg *config.Config) (application, error) {
	logger, err := provideLogger(cfg)
	if err != nil {
		return application{}, err
	}
	producer := provideKafkaProducer(cfg, logger)
	db, err := provideDatabase(cfg, logger)
	if err != nil {
		return application{}, err
	}
	repoStore := provideRepoStore(db, logger)
	scmProvider := provideSCM(logger)
	hookParser := provideParser(repoStore, scmProvider, logger)
	engine := provideRouter(producer, hookParser, logger)
	server := provideServer(engine, cfg, logger)
	cmdApplication := newApplication(server, logger)
	return cmdApplication, nil
}
