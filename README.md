# Message Broker in Go

## Overview
This project implements a basic message broker in Go, demonstrating core functionalities such as message production, consumption, and broker operations. It illustrates message brokering concepts and basic concurrency handling in Go.

## Core Features
- **Broker**: Manages topics and handles message publishing and subscribing.
- **Producer**: Sends messages to the broker.
- **Consumer**: Polls messages from the broker.
- **Thread Safety**: Ensures safe concurrent operations using mutexes.
- **Continuous Polling**: Supports real-time message consumption.
