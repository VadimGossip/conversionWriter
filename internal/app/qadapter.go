package app

import (
	"github.com/VadimGossip/conversionWriter/internal/domain"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type QueueAdapter struct {
	cfg            *domain.Config
	convRabbitConn *amqp.Connection
	convQueueChan  *amqp.Channel
}

func NewQueueAdapter(cfg *domain.Config) *QueueAdapter {
	dba := &QueueAdapter{}
	dba.cfg = cfg
	return dba
}

func (q *QueueAdapter) Connect() error {
	connectRabbitMQ, err := amqp.Dial(q.cfg.AMPQServerConfig.Url)
	if err != nil {
		return err
	}
	q.convRabbitConn = connectRabbitMQ

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		return err
	}
	q.convQueueChan = channelRabbitMQ
	return nil
}

func (q *QueueAdapter) Close() {
	if err := q.convQueueChan.Close(); err != nil {
		logrus.Errorf("Error occured on convQueueChan close: %s", err.Error())
	}

	if err := q.convRabbitConn.Close(); err != nil {
		logrus.Errorf("Error occured on convRabbitConn close: %s", err.Error())
	}
}
