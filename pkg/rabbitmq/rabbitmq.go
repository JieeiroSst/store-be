package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Connection *amqp.Connection
}

type RabbitMQRepo interface {
	Publish(queueName string, data []byte) error
	StartConsumer(queueName string) (<-chan amqp.Delivery, error)
}

func GetConnRabbitmq(rabbitURL string) (RabbitMQRepo, error) {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return nil, err
	}
	return &RabbitMQ{
		Connection: conn,
	}, err
}

func (r *RabbitMQ) Publish(queueName string, data []byte) error {
	channelRabbitMQ, err := r.Connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	_, err = channelRabbitMQ.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		panic(err)
	}
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(data),
	}

	if err := channelRabbitMQ.Publish(
		"",        // exchange
		queueName, // queue name
		false,     // mandatory
		false,     // immediate
		message,   // message to publish
	); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) StartConsumer(queueName string) (<-chan amqp.Delivery, error) {
	channelRabbitMQ, err := r.Connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	messages, err := channelRabbitMQ.Consume(
		queueName, // queue name
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // arguments
	)
	
	if err != nil {
		return nil, err
	}

	return messages, nil
}
