package cmd

import (
	"github.com/LambdaTest/photon/config"
	apiRouter "github.com/LambdaTest/photon/pkg/api"
	"github.com/LambdaTest/photon/pkg/lumber"
	"github.com/LambdaTest/photon/pkg/models"
	"github.com/LambdaTest/photon/pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// wire set for loading the server.
var serverSet = wire.NewSet(
	provideRouter,
	provideServer,
	newApplication,
)

// application is the main struct for the photon.
type application struct {
	logger lumber.Logger
	server *server.Server
}

// provideRouter is a Wire provider function that returns a
// router that is serves the provided handlers.
func provideRouter(kafkaProducer models.Producer, hookParser models.HookParser, logger lumber.Logger) *gin.Engine {
	router := apiRouter.New(kafkaProducer, hookParser, logger)
	return router.Handler()
}

// provideServer is a Wire provider function that returns an
// http server that is configured from the environment.
func provideServer(handler *gin.Engine, cfg *config.Config, logger lumber.Logger) *server.Server {
	return &server.Server{
		Addr:    ":" + cfg.Port,
		Handler: handler,
		Logger:  logger,
	}
}

// newApplication creates a new application struct.
func newApplication(
	srv *server.Server,
	logger lumber.Logger) application {
	return application{
		server: srv,
		logger: logger,
	}
}
