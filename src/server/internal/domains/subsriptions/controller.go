package subscriptions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	tokenStr := ctx.Query("token")
	if tokenStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Token parameter is required"})
		return
	}

	// Convert token to int64
	token, err := strconv.ParseInt(tokenStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token format"})
		return
	}

	// Call service method to get subscriptions
	subscriptions, err := c.svc.GetSubscriptionsByToken(ctx.Request.Context(), token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"subscriptions": subscriptions})
}