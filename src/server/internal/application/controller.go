package application

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"service/internal/domains/api"
	"service/internal/domains/notifications"
	"time"

	"github.com/gin-contrib/cors"
)

type Controller struct {
	notification *notifications.Controller
	api          *api.Controller
	Router       *gin.Engine
}

func NewController(svc *Service, r *gin.Engine) *Controller {
	return &Controller{
		notification: notifications.NewController(svc.Notification),
		api:          api.NewController(svc.Api),
		Router:       r,
	}
}

func (c *Controller) InitRouter() {
	c.api.Endpoints(c.Router)
	c.notification.Endpoints(c.Router)
}

func (c *Controller) Run(addr string, ctx context.Context) {
	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "X-Telegram-Auth"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	c.Router.Use(cors.New(config))

	c.InitRouter()

	server := &http.Server{
		Addr:    addr,
		Handler: c.Router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("HTTP server exited with error: %v", err)
		}
	}()
	log.Printf("HTTP server listening at %v", addr)

	<-ctx.Done()
	log.Println("Shutting down HTTP server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Errorf("HTTP server Shutdown Failed:%+v", err)
	} else {
		log.Println("HTTP server gracefully stopped")
	}
}
