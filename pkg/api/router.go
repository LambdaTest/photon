package api

import (
	"github.com/LambdaTest/photon/pkg/api/health"
	"github.com/LambdaTest/photon/pkg/api/hook"
	"github.com/LambdaTest/photon/pkg/api/kafka"
	"github.com/LambdaTest/photon/pkg/global"
	"github.com/LambdaTest/photon/pkg/lumber"
	"github.com/LambdaTest/photon/pkg/models"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// Router represents the routes for the http server.
type Router struct {
	logger        lumber.Logger
	hookParser    models.HookParser
	kafkaProducer models.Producer
}

// New returns a new Router instance
func New(kafkaProducer models.Producer, hookParser models.HookParser, logger lumber.Logger) Router {
	return Router{
		logger:        logger,
		hookParser:    hookParser,
		kafkaProducer: kafkaProducer,
	}
}

// Handler function will perform all the route operations
func (r Router) Handler() *gin.Engine {
	r.logger.Infof("Setting up routes")
	// set gin to release mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	// skip /health API from logs as will be required in probes
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/health"))
	// open-telemetry middleware
	router.Use(otelgin.Middleware(global.ServiceName))
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	r.logger.Infof("Setting up http handler")

	router.GET("/health", health.Handler)
	router.GET("/debug/kafkaproducer", kafka.ProducerStatsHandler(r.kafkaProducer))
	router.POST("/hook/:scmdriver", hook.Handler(r.hookParser, r.kafkaProducer, r.logger))
	pprof.Register(router)
	return router
}
