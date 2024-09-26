package rabbittemp

import (
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

// DeclareQueue declares a new queue
func (rmq *RabbitMQ) DeclareQueue(queueName string) (amqp.Queue, error) {
	queue, err := rmq.Channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Error().Msgf("Error while declaring queue: %s", err.Error())
		return amqp.Queue{}, err
	}
	rmq.Queues[queueName] = queue
	return queue, nil
}

// SendToQueue sends a message to the specified queue
func (rmq *RabbitMQ) SendToQueue(queueName string, body string) error {
	err := rmq.Channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Error().Msgf("Error while publishing message to queue: %s", err.Error())
		return err
	}
	return nil
}

// ReceiveFromQueue receives messages from the specified queue
func (rmq *RabbitMQ) ReceiveFromQueue(queueName string) (<-chan amqp.Delivery, error) {
	messages, err := rmq.Channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Error().Msgf("Error while consuming messages from queue: %s", err.Error())
		return nil, err
	}
	return messages, nil
}

// Close closes the RabbitMQ channel and connection
func (rmq *RabbitMQ) Close() {
	if rmq.Channel != nil {
		rmq.Channel.Close()
	}
	if rmq.client != nil {
		rmq.client.Close()
	}
}
