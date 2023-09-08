package wbus

import (
	"github.com/WhiteMaks/poly_day_bus_schema/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRabbit_CreateQueue(t *testing.T) {
	err := rabbit.CreateQueue("test-queue")
	assert.NoError(t, err, "No error must be")
}

func TestRabbit_CreateExchange(t *testing.T) {
	err := rabbit.CreateExchange("test-exchange")
	assert.NoError(t, err, "No error must be")
}

func TestRabbit_LinkQueueWithExchanger(t *testing.T) {
	queueName := "test-queue-for-exchange"
	exchangeName := "test-exchange-for-queue"

	err := rabbit.CreateQueue(queueName)
	assert.NoError(t, err, "No error must be")

	err = rabbit.CreateExchange(exchangeName)
	assert.NoError(t, err, "No error must be")

	err = rabbit.LinkQueueWithExchanger(queueName, exchangeName)
	assert.NoError(t, err, "No error must be")
}

func TestRabbit_LinkQueueWithExchanger_QueueNotFound(t *testing.T) {
	queueName := "test-queue-for-exchange-not-exist"
	exchangeName := "test-exchange-for-queue-found"

	err := rabbit.CreateExchange(exchangeName)
	assert.NoError(t, err, "No error must be")

	err = rabbit.LinkQueueWithExchanger(queueName, exchangeName)
	assert.Error(t, err, "Error must be")
}

func TestRabbit_LinkQueueWithExchanger_ExchangeNotFound(t *testing.T) {
	queueName := "test-queue-for-exchange-found"
	exchangeName := "test-exchange-for-queue-not-exist"

	err := rabbit.CreateQueue(queueName)
	assert.NoError(t, err, "No error must be")

	err = rabbit.LinkQueueWithExchanger(queueName, exchangeName)
	assert.Error(t, err, "Error must be")
}

func TestRabbit_SubscribeOnQueue(t *testing.T) {
	queueName := "test-queue-for-subscribe"

	err := rabbit.CreateQueue(queueName)
	assert.NoError(t, err, "No error must be")

	messagesChan := make(chan schema.Message)

	err = rabbit.SubscribeOnQueue(queueName, func(message schema.Message) {
		messagesChan <- message
	})
	assert.NoError(t, err, "No error must be")
}

func TestRabbit_SubscribeOnQueue_QueueNotFound(t *testing.T) {
	queueName := "test-queue-for-subscribe-queue-not-exist"

	messagesChan := make(chan schema.Message)

	err := rabbit.SubscribeOnQueue(queueName, func(message schema.Message) {
		messagesChan <- message
	})
	assert.Error(t, err, "Error must be")
}

func TestRabbit_SendMessageInQueue(t *testing.T) {
	queueName := "test-queue-for-send-message"

	err := rabbit.CreateQueue(queueName)
	assert.NoError(t, err, "No error must be")

	message := schema.Message{
		FromMicroservice: "test-ms-1",
		ToMicroservice:   rabbit.Name,
		Method:           "test-method",
		Version:          "1.0",
		Body:             []byte("{\"test-field\":123}"),
	}

	err = rabbit.SendMessageInQueue(queueName, message)
	assert.NoError(t, err, "No error must be")
}

func TestRabbit_SendMessageInQueue_QueueNotFound(t *testing.T) {
	queueName := "test-queue-for-send-message-queue-not-exist"

	message := schema.Message{
		FromMicroservice: "test-ms-1",
		ToMicroservice:   rabbit.Name,
		Method:           "test-method",
		Version:          "1.0",
		Body:             []byte("{\"test-field\":123}"),
	}

	err := rabbit.SendMessageInQueue(queueName, message)
	assert.NoError(t, err, "No error must be")
}

func TestRabbit_SendMessageInQueue_AndReadMessage(t *testing.T) {
	queueName := "test-queue-for-subscribe-and-read"

	expectedMessage := schema.Message{
		FromMicroservice: "test-ms-1",
		ToMicroservice:   rabbit.Name,
		Method:           "test-method",
		Version:          "1.0",
		Body:             []byte("{\"test-field\":123}"),
	}

	err := rabbit.CreateQueue(queueName)
	assert.NoError(t, err, "No error must be")

	messagesChan := make(chan schema.Message)

	err = rabbit.SubscribeOnQueue(queueName, func(message schema.Message) {
		messagesChan <- message
	})
	assert.NoError(t, err, "No error must be")

	err = rabbit.SendMessageInQueue(queueName, expectedMessage)
	assert.NoError(t, err, "No error must be")

	actualMessage := <-messagesChan
	assert.Equal(t, expectedMessage, actualMessage, "Messages must be equals")
}
