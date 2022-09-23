package kafka

import (
	"net/http"

	"github.com/LambdaTest/photon/pkg/models"
	"github.com/gin-gonic/gin"
)

// ProducerStatsHandler handler for extracting the kafka producer stats.
func ProducerStatsHandler(kafkaProducer models.Producer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, kafkaProducer.Stats())
	}
}
