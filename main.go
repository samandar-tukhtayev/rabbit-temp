package rabbittemp

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

// New initializes and returns a new RabbitMQ connection and channel
func New(cfg Config) (*RabbitMQ, error) {
	conn, err := amqp.Dial(
		fmt.Sprintf(
			"amqp://%s:%s@%s:%s/",
			cfg.RabbitMQUser,
			cfg.RabbitMQPassword,
			cfg.RabbitMQHost,
			cfg.RabbitMQPort,
		),
	)
	if err != nil {
		log.Error().Msgf("Error while connecting to RabbitMQ: %s", err.Error())
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Error().Msgf("Error while opening channel: %s", err.Error())
		conn.Close() // Ensure the connection is closed if the channel cannot be opened
		return nil, err
	}

	err = channel.Qos(1, 0, false)
	if err != nil {
		log.Error().Msgf("Error while setting QoS: %s", err.Error())
		channel.Close()
		conn.Close() // Ensure both channel and connection are closed in case of QoS error
		return nil, err
	}

	rmq := &RabbitMQ{
		Queues:  make(map[string]amqp.Queue),
		Channel: channel,
		Cfg:     cfg,
		client:  conn,
	}

	return rmq, nil
}
