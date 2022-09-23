package hook

import (
	"errors"
	"net/http"

	"github.com/LambdaTest/photon/internal/json"

	"github.com/LambdaTest/photon/pkg/errs"
	"github.com/LambdaTest/photon/pkg/lumber"
	"github.com/LambdaTest/photon/pkg/models"
	"github.com/gin-gonic/gin"
)

// TODO: can add checks to see if org is active and credits left.

// Handler is the git scm webhook handler.
func Handler(
	hookParser models.HookParser,
	kafkaProducer models.Producer,
	logger lumber.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		pathParam := c.Param("scmdriver")
		driver := models.SCMDriver(pathParam)
		if err := driver.VerifyDriver(); err != nil {
			c.JSON(http.StatusNotFound, err)
			return
		}
		payload, headers, err := hookParser.Parse(c.Request, driver)
		if err != nil {
			if errors.Is(err, errs.ErrPingEvent) || errors.Is(err, errs.ErrNotSupported) {
				c.JSON(http.StatusOK, err)
				return
			}
			if errors.Is(err, errs.ErrRepoNotFound) || errors.Is(err, errs.ErrInvalidSCMProvider) {
				c.JSON(http.StatusNotFound, err)
				return
			}
			if errors.Is(err, errs.ErrUnknownEvent) || errors.Is(err, errs.ErrSignatureInvalid) {
				c.JSON(http.StatusBadRequest, err)
				return
			}
			logger.Errorf("failed to parse webhook for repository, git_provider %s, error %v", driver, err)
			c.JSON(http.StatusInternalServerError, errs.GenericErrorMessage)
			return
		}

		message, err := json.Marshal(payload)
		if err != nil {
			logger.Errorf("failed to marshal webhook, repository %s/%s, git_provider %s, error: %v",
				payload.Repository().Name, payload.Repository().Namespace, driver, err)
			c.JSON(http.StatusInternalServerError, errs.GenericErrorMessage)
		}

		if err := kafkaProducer.WriteMessage(message, headers...); err != nil {
			logger.Errorf("failed to write message to kafka topic, repository %s/%s, git_provider %s, error: %v",
				payload.Repository().Name, payload.Repository().Namespace, driver, err)
			c.JSON(http.StatusInternalServerError, errs.GenericErrorMessage)
		}
		c.Status(http.StatusNoContent)
	}
}
