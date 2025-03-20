package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/login/callback", cont.TelegramLoginCallback)
	r.GET("/qr", cont.GenerateQRCode)
}

// /login/callback
func (cont *Controller) TelegramLoginCallback(c *gin.Context) {
	userData := c.Request.URL.Query()

	isValid := cont.svc.ValidateTelegramData(userData)
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные данные Телеграма"})
		return
	}

	//err := cont.svc.ProcessTelegramLogin(c.Request.Context(), userData)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки данных пользователя"})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{"message": "Успешный вход через Телеграм"})
}

// /qr?patient_id=&token=
func (cont *Controller) GenerateQRCode(c *gin.Context) {
	patientID := c.Query("patient_id")
	token := c.Query("token")

	if patientID == "" || token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметры patient_id и token обязательны"})
		return
	}

	qrCodeData, err := cont.svc.GenerateQRCode(c.Request.Context(), patientID, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации QR-кода"})
		return
	}

	// Устанавливаем заголовки для передачи изображения
	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", string(len(qrCodeData)))
	c.Writer.Write(qrCodeData)
}
