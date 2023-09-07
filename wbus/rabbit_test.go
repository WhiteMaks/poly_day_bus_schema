package wbus

import (
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
	queueName := "test-queue"
	exchangeName := "test-exchange"

	err := rabbit.CreateQueue(queueName)
	assert.NoError(t, err, "No error must be")

	err = rabbit.CreateExchange(exchangeName)
	assert.NoError(t, err, "No error must be")

	err = rabbit.LinkQueueWithExchanger(queueName, exchangeName)
	assert.NoError(t, err, "No error must be")
}
