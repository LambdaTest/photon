package cmd

import (
	"github.com/LambdaTest/photon/config"
	"github.com/LambdaTest/photon/pkg/kafka"
	"github.com/LambdaTest/photon/pkg/lumber"
	"github.com/LambdaTest/photon/pkg/models"
	"github.com/google/wire"
)

var kafkaSet = wire.NewSet(
	provideKafkaProducer,
)

// provideKafka is a Wire provider function that returns a
// kafka producer client that is serves the provided handlers.
func provideKafkaProducer(cfg *config.Config, logger lumber.Logger) models.Producer {
	return kafka.New(cfg, logger)
}
