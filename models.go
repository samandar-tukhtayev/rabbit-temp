package rabbittemp

import "github.com/streadway/amqp"

// Config holds RabbitMQ configuration
type Config struct {
	RabbitMQUser     string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string
}

// RabbitMQ holds the RabbitMQ connection and channel
type RabbitMQ struct {
	Queues  map[string]amqp.Queue
	Channel *amqp.Channel
	Cfg     Config
	client  *amqp.Connection
}
