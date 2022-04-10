package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Connection *amqp.Connection
}

type RabbitMQRepo interface {
	Publish(queueName, routingKey, exchange string, data []byte) error
	StartConsumer(queueName, routingKey, exchange string) (<-chan amqp.Delivery, error)
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

func (r *RabbitMQ) Publish(queueName, routingKey, exchange string, data []byte) error {
	ch, err := r.Connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	if err != nil {
		return err
	}
	_, err = ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	err = ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *RabbitMQ) StartConsumer(queueName, routingKey, exchange string) (<-chan amqp.Delivery, error) {
	ch, err := r.Connection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange,
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	queue, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(
		queue.Name,
		routingKey,
		exchange,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	msgs, err := ch.Consume(
		queue.Name,
		"", // consumer name
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	return msgs, nil
}
