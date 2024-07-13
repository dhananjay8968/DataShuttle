package broker

import (
	"sync"

	"datashuttle/models"
)

// Broker represents the message broker.
type Broker struct {
	mu      sync.Mutex
	topics  map[string][]*models.Message
	offsets map[string]int64
}

// NewBroker creates a new Broker instance.
func NewBroker() *Broker {
	return &Broker{
		topics:  make(map[string][]*models.Message),
		offsets: make(map[string]int64),
	}
}

// Publish adds a message to the broker.
func (b *Broker) Publish(topic, key, value string) int64 {
	b.mu.Lock()
	defer b.mu.Unlock()

	offset := b.offsets[topic]
	message := &models.Message{
		Topic:  topic,
		Key:    key,
		Value:  value,
		Offset: offset,
	}

	b.topics[topic] = append(b.topics[topic], message)
	b.offsets[topic]++

	return offset
}

// Subscribe retrieves messages from a topic starting from the given offset.
func (b *Broker) Subscribe(topic string, offset int64) []*models.Message {
	b.mu.Lock()
	defer b.mu.Unlock()

	messages, ok := b.topics[topic]
	if !ok {
		return nil
	}

	var result []*models.Message
	for _, msg := range messages {
		if msg.Offset >= offset {
			result = append(result, msg)
		}
	}

	return result
}
