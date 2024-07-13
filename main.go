package main

import (
	"fmt"
	"time"

	"datashuttle/broker"
	"datashuttle/models"
)

// displayMessages prints out the messages in a readable format.
func displayMessages(messages []*models.Message) {
	if len(messages) == 0 {
		fmt.Println("No new messages.")
		return
	}

	for _, msg := range messages {
		fmt.Printf("Received message -> Topic: %s, Key: %s, Value: %s, Offset: %d\n", msg.Topic, msg.Key, msg.Value, msg.Offset)
	}
}

func main() {
	// Initialize broker
	b := broker.NewBroker()

	// Initialize producer
	producer := broker.NewProducer(b)

	// Sending messages to the broker
	fmt.Println("Producing messages...")
	producer.Send("test-topic", "key1", "Hello, World!")
	producer.Send("test-topic", "key2", "Kafka in Go!")
	producer.Send("test-topic", "key3", "Another message!")

	// Initialize consumer
	consumer := broker.NewConsumer(b, "test-topic")

	// Demonstrating message consumption
	fmt.Println("\nConsuming messages...")
	messages := consumer.Poll()
	displayMessages(messages)

	// Demonstrating continuous polling
	fmt.Println("\nStarting continuous polling...")
	stopPolling := make(chan bool)
	go func() {
		for {
			select {
			case <-stopPolling:
				return
			default:
				messages := consumer.Poll()
				displayMessages(messages)
				time.Sleep(3 * time.Second)
			}
		}
	}()

	// Producing more messages after some delay to simulate real-time message production
	time.Sleep(5 * time.Second)
	fmt.Println("\nProducing more messages...")
	producer.Send("test-topic", "key4", "Real-time message 1")
	producer.Send("test-topic", "key5", "Real-time message 2")

	// Allow time for messages to be polled
	time.Sleep(10 * time.Second)
	stopPolling <- true

	fmt.Println("\nDemonstration complete.")
}
