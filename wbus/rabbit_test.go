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

func TestRabbit_SendMessageInExchange(t *testing.T) {
	queueName := "test-queue-for-exchange-send-message"
	exchangeName := "test-exchange-for-queue-send-message"

	err := rabbit.CreateQueue(queueName)
	assert.NoError(t, err, "No error must be")

	err = rabbit.CreateExchange(exchangeName)
	assert.NoError(t, err, "No error must be")

	err = rabbit.LinkQueueWithExchanger(queueName, exchangeName)
	assert.NoError(t, err, "No error must be")

	message := schema.Message{
		FromMicroservice: "test-ms-1",
		ToMicroservice:   rabbit.Name,
		Method:           "test-method",
		Version:          "1.0",
		Body:             []byte("{\"test-field\":123}"),
	}

	err = rabbit.SendMessageInExchange(exchangeName, message)
	assert.NoError(t, err, "No error must be")
}

func TestRabbit_SendMessageInExchange_ExchangeNotFound(t *testing.T) {
	exchangeName := "test-exchange-for-queue-send-message-not-exist"

	message := schema.Message{
		FromMicroservice: "test-ms-1",
		ToMicroservice:   rabbit.Name,
		Method:           "test-method",
		Version:          "1.0",
		Body:             []byte("{\"test-field\":123}"),
	}

	err := rabbit.SendMessageInExchange(exchangeName, message)
	assert.NoError(t, err, "No error must be")
}

func TestRabbit_SendMessageInExchange_AndReadMessage(t *testing.T) {
	queueName1 := "test-queue-for-exchange-send-message-and-read-1"
	queueName2 := "test-queue-for-exchange-send-message-and-read-2"
	exchangeName := "test-exchange-for-queue-send-message-and-read"

	err := rabbit.CreateQueue(queueName1)
	assert.NoError(t, err, "No error must be")

	err = rabbit.CreateQueue(queueName2)
	assert.NoError(t, err, "No error must be")

	err = rabbit.CreateExchange(exchangeName)
	assert.NoError(t, err, "No error must be")

	err = rabbit.LinkQueueWithExchanger(queueName1, exchangeName)
	assert.NoError(t, err, "No error must be")

	err = rabbit.LinkQueueWithExchanger(queueName2, exchangeName)
	assert.NoError(t, err, "No error must be")

	expectedMessage := schema.Message{
		FromMicroservice: "test-ms-1",
		ToMicroservice:   rabbit.Name,
		Method:           "test-method",
		Version:          "1.0",
		Body:             []byte("{\"test-field\":123}"),
	}

	messagesChan1 := make(chan schema.Message)
	messagesChan2 := make(chan schema.Message)

	err = rabbit.SubscribeOnQueue(queueName1, func(message schema.Message) {
		messagesChan1 <- message
	})
	assert.NoError(t, err, "No error must be")

	err = rabbit.SubscribeOnQueue(queueName2, func(message schema.Message) {
		messagesChan2 <- message
	})
	assert.NoError(t, err, "No error must be")

	err = rabbit.SendMessageInExchange(exchangeName, expectedMessage)
	assert.NoError(t, err, "No error must be")

	actualMessage1 := <-messagesChan1
	assert.Equal(t, expectedMessage, actualMessage1, "Messages must be equals")

	actualMessage2 := <-messagesChan2
	assert.Equal(t, expectedMessage, actualMessage2, "Messages must be equals")
}
