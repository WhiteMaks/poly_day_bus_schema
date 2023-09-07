package wbus

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Name       string
	connection *amqp.Connection
	chanel     *amqp.Channel
	queues     map[string]amqp.Queue
}

func NewRabbit(name string) *Rabbit {
	return &Rabbit{
		Name:   name,
		queues: make(map[string]amqp.Queue),
	}
}

func (rabbit *Rabbit) Start(address string) error {
	connection, err := amqp.Dial(address)
	if err != nil {
		return err
	}

	chanel, err := connection.Channel()
	if err != nil {
		return err
	}

	rabbit.connection = connection
	rabbit.chanel = chanel

	return nil
}

func (rabbit *Rabbit) CreateQueue(name string) error {
	queue, err := rabbit.chanel.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	rabbit.queues[name] = queue

	return nil
}

func (rabbit *Rabbit) CreateExchange(name string) error {
	err := rabbit.chanel.ExchangeDeclare(
		name,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (rabbit *Rabbit) LinkQueueWithExchanger(queueName string, exchangeName string) error {
	err := rabbit.chanel.QueueBind(
		queueName,
		"",
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
