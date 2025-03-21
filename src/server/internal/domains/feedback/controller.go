package feedback

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "service/internal/domains/feedback/models"
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
    // Public endpoints
    r.POST("/feedback-send", c.SendFeedback)
    
    // Admin endpoints with authorization
    authorized := r.Group("/", middleware.UnifiedAuthenticationMiddleware())
    authorized.GET("/feedback", c.ListFeedback)
    authorized.POST("/feedback-answer/:id", c.AnswerFeedback)
}

func (c *Controller) ListFeedback(ctx *gin.Context) {
    // Optional: Check if user has admin rights here
    
    feedbacks, err := c.svc.ListFeedback(ctx.Request.Context())
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"feedback": feedbacks})
}

func (c *Controller) SendFeedback(ctx *gin.Context) {
    var feedback models.Feedback
    if err := ctx.BindJSON(&feedback); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // Validate required fields
    if feedback.Header == "" || feedback.Content == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Header and content are required"})
        return
    }
    
    // If user is authenticated, get their ID
    userID := ctx.GetString("telegram_user_id")
    if userID != "" {
        userIDint, err := strconv.ParseInt(userID, 10, 64)
        if err == nil {
            feedback.UserID = userIDint
        }
    }
    
    if err := c.svc.SendFeedback(ctx.Request.Context(), feedback); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    ctx.JSON(http.StatusOK, gin.H{"message": "Feedback sent successfully"})
}

func (c *Controller) AnswerFeedback(ctx *gin.Context) {
    // Check if the user has admin rights
    // This is a simplified authorization check - implement proper role checking
    
    idParam := ctx.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feedback ID"})
        return
    }
    
    var request struct {
        Answer string `json:"answer" binding:"required"`
    }
    
    if err := ctx.BindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if err := c.svc.AnswerFeedback(ctx.Request.Context(), id, request.Answer); err != nil {
        if err == ErrFeedbackNotFound {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "Feedback not found"})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    ctx.JSON(http.StatusOK, gin.H{"message": "Feedback answered successfully"})
}