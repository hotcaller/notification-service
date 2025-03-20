package notifications

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/internal/domains/notifications/models"
	middleware "service/internal/infrastructure/middlewares"
	"strconv"
)

type Controller struct {
	svc *Service
}

func NewController(svc *Service) *Controller {
	return &Controller{svc: svc}
}

func (c *Controller) Endpoints(r *gin.Engine) {

	authorized := r.Group("/", middleware.UnifiedAuthenticationMiddleware())
	authorized.GET("/notifications", c.ListNotifications)
	authorized.GET("/notifications/:id", c.GetNotificationByID)
	authorized.POST("/notifications", c.SendNotification)
}

func (c *Controller) ListNotifications(ctx *gin.Context) {
	// Получаем идентификатор аутентифицированного пользователя из контекста
	userID := ctx.GetString("telegram_user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неавторизованный пользователь"})
		return
	}

	// Вызываем сервисный метод с учетом userID
	notifications, err := c.svc.ListNotifications(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"notifications": notifications})
}

func (c *Controller) GetNotificationByID(ctx *gin.Context) {
	// Получаем идентификатор аутентифицированного пользователя из контекста
	userID := ctx.GetString("telegram_user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неавторизованный пользователь"})
		return
	}

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID уведомления"})
		return
	}

	notification, err := c.svc.GetNotificationByID(ctx.Request.Context(), id, userID)
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
	if err := ctx.BindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.svc.SendNotification(ctx.Request.Context(), notification); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Уведомление успешно отправлено"})
}
