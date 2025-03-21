package subscriptions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	svc *Service
}

func NewController(svc *Service) *Controller {
	return &Controller{svc: svc}
}

func (c *Controller) Endpoints(r *gin.Engine) {
	r.GET("/subscriptions", c.GetSubscriptions)
}

func (c *Controller) GetSubscriptions(ctx *gin.Context) {
	// Get token from query parameter
	token := ctx.Query("token")
	if token == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Token parameter is required"})
			return
	}

	// No need to convert to int64 anymore
	subscriptions, err := c.svc.GetSubscriptionsByToken(ctx.Request.Context(), token)
	if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}

	ctx.JSON(http.StatusOK, gin.H{"subscriptions": subscriptions})
}