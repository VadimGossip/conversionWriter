package app

import (
	"github.com/VadimGossip/conversionWriter/internal/domain"
)

type Factory struct {
	queueAdapter *QueueAdapter
}

var factory *Factory

func newFactory(cfg *domain.Config, queueAdapter *QueueAdapter) *Factory {
	factory = &Factory{queueAdapter: queueAdapter}
	return factory
}
