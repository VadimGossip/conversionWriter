package app

import (
	"github.com/VadimGossip/conversionWriter/internal/conversion"
	"github.com/VadimGossip/conversionWriter/internal/domain"
)

type Factory struct {
	queueAdapter *QueueAdapter
	convService  conversion.Service
}

var factory *Factory

func newFactory(cfg *domain.Config, queueAdapter *QueueAdapter) *Factory {
	factory = &Factory{queueAdapter: queueAdapter}
	factory.convService = conversion.NewService(cfg.AMPQServerConfig.ConvQueueName, factory.queueAdapter.convQueueChan)
	return factory
}
