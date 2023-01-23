package config

import (
	"github.com/VadimGossip/conversionWriter/internal/domain"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func setFromEnv(cfg *domain.Config) error {
	if err := envconfig.Process("writer_http", &cfg.WriterHttpServer); err != nil {
		return err
	}
	if err := envconfig.Process("writer_metrics_http", &cfg.WriterMetricsHttpServer); err != nil {
		return err
	}
	if err := envconfig.Process("ampq_server", &cfg.AMPQServerConfig); err != nil {
		return err
	}

	return nil
}

func Init() (*domain.Config, error) {

	var cfg domain.Config
	if err := setFromEnv(&cfg); err != nil {
		return nil, err
	}
	//temp
	cfg.AMPQServerConfig.ConvQueueName = "ConversionQueue"

	logrus.Infof("Config %v", cfg)
	return &cfg, nil
}
