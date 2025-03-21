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
    r.POST("/feedback-send", c.SendFeedback)
    r.GET("/admin-feedback", c.ListAllFeedback) 
    
    authorized := r.Group("/", middleware.UnifiedAuthenticationMiddleware())
    authorized.GET("/feedback", c.ListUserFeedback)
    authorized.POST("/feedback-answer/:id", c.AnswerFeedback)
}

// ListUserFeedback returns only feedbac	k submitted by the authenticated user
func (c *Controller) ListUserFeedback(ctx *gin.Context) {
    userID := ctx.GetString("telegram_user_id")
    if userID == "" {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    
    userIDint, err := strconv.ParseInt(userID, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    
    feedbacks, err := c.svc.ListUserFeedback(ctx.Request.Context(), userIDint)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"feedback": feedbacks})
}

// ListAllFeedback returns all feedback (for admin use)
func (c *Controller) ListAllFeedback(ctx *gin.Context) {
    // No authorization check - this is an admin endpoint
    feedbacks, err := c.svc.ListAllFeedback(ctx.Request.Context())
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