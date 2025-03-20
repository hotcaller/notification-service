package application

import (
	"context"
	"fmt"
	"github.com/Arlandaren/pgxWrappy/pkg/postgres"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"service/internal/infrastructure/config"
	"service/internal/infrastructure/kafka"
	"service/internal/infrastructure/storage/minio"
	"service/internal/infrastructure/storage/redis"
	"sync"
	"syscall"
)

type App struct {
	Controller     *Controller
	allAppConsumer *AppConsumer
	KafkaBrokers   []string
	wg             *sync.WaitGroup
	cfg            *config.Config
}

func NewApp(db *postgres.Wrapper, rdb *redis.RDB, s3 *minio.Minio, r *gin.Engine, cfg *config.Config) *App {
	kafkaBrokers := []string{os.Getenv("KAFKA_ADDRESS")}

	kafkaCfg := kafka.NewConfig(kafkaBrokers)
	fmt.Println(kafkaCfg)
	kafkaProducer, err := kafka.NewProducer(kafkaCfg)
	if err != nil {
		log.Fatalf("failed to initialize Kafka producer: %v", err)
	}

	repo := NewRepository(db, rdb, s3)
	svc := NewService(repo, kafkaProducer)
	controller := NewController(svc, r)

	appConsumer := NewAppConsumer(svc, nil)
	consumer := kafka.NewConsumer(appConsumer)
	appConsumer.SetConsumer(consumer)

	return &App{
		KafkaBrokers:   kafkaBrokers,
		Controller:     controller,
		wg:             &sync.WaitGroup{},
		cfg:            cfg,
		allAppConsumer: appConsumer,
	}
}

func (app *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		group := "server_group"
		topics := []string{
			"notification_created",
		}

		if err := app.allAppConsumer.Start(ctx, app.KafkaBrokers, group, topics); err != nil {
			log.Fatalf("Failed to start Kafka consumer: %v", err)
		}
	}()

	// Start the HTTP server
	wg.Add(1)
	go func() {
		defer wg.Done()
		app.Controller.Run(app.cfg.Address.Http, ctx)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutdown signal received")

	cancel()

	wg.Wait()
	log.Println("Application gracefully stopped.")
}
