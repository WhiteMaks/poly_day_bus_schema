package wbus

import (
	"context"
	"encoding/json"
	"github.com/WhiteMaks/poly_day_bus_schema/schema"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Rabbit struct {
	Name       string
	connection *amqp.Connection
	//chanel     *amqp.Channel
	queues map[string]amqp.Queue
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

	rabbit.connection = connection
	//rabbit.chanel = chanel

	return nil
}

func (rabbit *Rabbit) CreateQueue(name string) error {
	chanel, err := rabbit.connection.Channel()
	if err != nil {
		return err
	}
	defer chanel.Close()

	queue, err := chanel.QueueDeclare(
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
	chanel, err := rabbit.connection.Channel()
	if err != nil {
		return err
	}
	defer chanel.Close()

	err = chanel.ExchangeDeclare(
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
	chanel, err := rabbit.connection.Channel()
	if err != nil {
		return err
	}
	defer chanel.Close()

	err = chanel.QueueBind(
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

func (rabbit *Rabbit) SubscribeOnQueue(queueName string, messageHandler schema.MessageHandler) error {
	chanel, err := rabbit.connection.Channel()
	if err != nil {
		return err
	}

	deliveryChan, err := chanel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		defer chanel.Close()

		for message := range deliveryChan {
			log.Printf("Received a message: %v", string(message.Body))

			var messageStruct schema.Message
			err = json.Unmarshal(message.Body, &messageStruct)
			if err != nil {
				log.Printf("Error unmarshal json: %v", err)
				continue
			}

			if messageStruct.ToMicroservice == rabbit.Name {
				messageHandler(messageStruct)
			}
		}
	}()

	return nil
}

func (rabbit *Rabbit) SendMessageInQueue(queueName string, message schema.Message) error {
	chanel, err := rabbit.connection.Channel()
	if err != nil {
		return err
	}
	defer chanel.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonBody, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = chanel.PublishWithContext(
		ctx,
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (rabbit *Rabbit) SendMessageInExchange(exchangeName string, message schema.Message) error {
	chanel, err := rabbit.connection.Channel()
	if err != nil {
		return err
	}
	defer chanel.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonBody, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = chanel.PublishWithContext(
		ctx,
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
