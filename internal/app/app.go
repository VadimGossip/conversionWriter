package app

import (
	"github.com/VadimGossip/conversionWriter/internal/config"
	"github.com/VadimGossip/conversionWriter/internal/domain"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

type App struct {
	*Factory
	name         string
	appStartedAt time.Time
	cfg          *domain.Config
	convConsumer *consumer
}

func NewApp(name string, appStartedAt time.Time) *App {
	return &App{
		name:         name,
		appStartedAt: appStartedAt,
	}
}

func (app *App) Run() {
	cfg, err := config.Init()
	if err != nil {
		logrus.Fatalf("Config initialization error %s", err)
	}
	app.cfg = cfg

	qAdapter := NewQueueAdapter(app.cfg)
	if err := qAdapter.Connect(); err != nil {
		logrus.Fatalf("Fail to connect ampq %s", err)
	}

	app.Factory = newFactory(app.cfg, qAdapter)
	app.convConsumer = NewConsumer(app.cfg.AMPQServerConfig.ConvQueueName, app.queueAdapter.convQueueChan)

	if err := app.convConsumer.Run(); err != nil {
		logrus.Fatalf("Fail to run consumer ampq %s", err)
	}

	logrus.Print("Conversion writer service started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	//disconnect from queue

	logrus.Print("Conversion writer service stopped")
}
