package broker

import "datashuttle/models"

// Consumer represents a message consumer.
type Consumer struct {
	broker *Broker
	offset int64
	topic  string
}

// NewConsumer creates a new Consumer instance.
func NewConsumer(broker *Broker, topic string) *Consumer {
	return &Consumer{broker: broker, topic: topic, offset: 0}
}

// Poll retrieves new messages from the broker.
func (c *Consumer) Poll() []*models.Message {
	messages := c.broker.Subscribe(c.topic, c.offset)
	if len(messages) > 0 {
		c.offset = messages[len(messages)-1].Offset + 1
	}
	return messages
}
