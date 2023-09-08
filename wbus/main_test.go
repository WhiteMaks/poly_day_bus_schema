package wbus

import (
	"log"
	"os"
	"testing"
)

const (
	rabbitConStr = "amqp://user:password@localhost:5672/"
)

var rabbit *Rabbit

func TestMain(m *testing.M) {
	rabbit = NewRabbit("test-bus-schema")
	err := rabbit.Start(rabbitConStr)
	if err != nil {
		log.Fatalf("error start bus [ %v ]: %v", rabbit.Name, err)
	}

	os.Exit(m.Run())
}
