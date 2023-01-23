package app

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Consumer interface {
	Run() error
}

type consumer struct {
	convQueueName string
	convQueueChan *amqp.Channel
}

var _ Consumer = (*consumer)(nil)

func NewConsumer(convQueueName string, convQueueChan *amqp.Channel) *consumer {
	return &consumer{convQueueName: convQueueName, convQueueChan: convQueueChan}
}

func (c *consumer) Run() error {
	conversions, err := c.convQueueChan.Consume(
		c.convQueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	var forever chan struct{}
	go func() {
		for c := range conversions {
			fmt.Printf("received %s\n", c.Body)
			c.Ack(false)
		}
	}()
	<-forever
	return nil
}
