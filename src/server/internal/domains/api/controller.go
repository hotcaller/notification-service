package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"service/internal/infrastructure/metrics"
)

type Controller struct {
	svc *Service
}

func NewController(svc *Service) *Controller {
	return &Controller{
		svc: svc,
	}
}

func (cont *Controller) Endpoints(r *gin.Engine) {
	r.Use(metrics.MetricsMiddleware())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })

}
