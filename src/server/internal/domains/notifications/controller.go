// internal/domains/notifications/controller.go

package notifications

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"service/internal/domains/notifications/models"
)

type Controller struct {
	svc *Service
}

func NewController(svc *Service) *Controller {
	return &Controller{svc: svc}
}

func (c *Controller) Endpoints(r *gin.Engine) {
	r.GET("/notifications", c.ListNotifications)
	r.GET("/notifications/:id", c.GetNotificationByID)
	r.POST("/notifications", c.SendNotification)
}

func (c *Controller) ListNotifications(ctx *gin.Context) {
	notifications, err := c.svc.ListNotifications(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"notifications": notifications})
}

func (c *Controller) GetNotificationByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID уведомления"})
		return
	}

	notification, err := c.svc.GetNotificationByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if notification == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Уведомление не найдено"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"notification": notification})
}

func (c *Controller) SendNotification(ctx *gin.Context) {
	var notification models.Notification
	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.svc.SendNotification(ctx.Request.Context(), notification); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Уведомление отправлено"})
}
