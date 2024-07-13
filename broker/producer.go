package broker

// Producer represents a message producer.
type Producer struct {
    broker *Broker
}

// NewProducer creates a new Producer instance.
func NewProducer(broker *Broker) *Producer {
    return &Producer{broker: broker}
}

// Send sends a message to the broker.
func (p *Producer) Send(topic, key, value string) int64 {
    return p.broker.Publish(topic, key, value)
}
