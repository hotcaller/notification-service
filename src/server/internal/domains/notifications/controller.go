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
    r.POST("/notifications", c.SendNotification)
}

func (c *Controller) ListNotifications(ctx *gin.Context) {
    // Получаем идентификатор аутентифицированного пользователя из контекста
    userID := ctx.GetString("telegram_user_id")
    if userID == "" {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неавторизованный пользователь"})
        return
    }

    userIDint, err := strconv.ParseInt(userID, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный идентификатор пользователя"})
        return
    }
    // Вызываем сервисный метод с учетом userID
    notifications, err := c.svc.ListNotifications(ctx.Request.Context(), userIDint)
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

    userIDint, err := strconv.ParseInt(userID, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный идентификатор пользователя"})
        return
    }
    idInt, err := strconv.ParseInt(idParam, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный идентификатор уведомления"})
        return
    }

    notification, err := c.svc.GetNotificationByID(ctx.Request.Context(), idInt, userIDint)
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
    
    // Validate required fields
    if notification.Header == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Header is required"})
        return
    }
    
    if notification.Message == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Message is required"})
        return
    }
    
    switch notification.Type {
    case models.NotificationTypeNews, models.NotificationTypeReminder, 
         models.NotificationTypeWarning, models.NotificationTypeImportant:
    case "":
        notification.Type = models.NotificationTypeNews
    default:
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid notification type. Must be one of: news, reminder, warning, important",
        })
        return
    }
    
    // Validate broadcast case - must have org_token when target_id is 0
    if notification.TargetID == 0 && notification.OrgToken == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Organization token is required for broadcast notifications",
        })
        return
    }
    
    if err := c.svc.SendNotification(ctx.Request.Context(), notification); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    // Customize the response message based on broadcast or targeted notification
    message := "Уведомление успешно отправлено"
    if notification.TargetID == 0 {
        message = "Уведомление успешно отправлено всем пользователям организации"
    }
    
    ctx.JSON(http.StatusOK, gin.H{"message": message})
}