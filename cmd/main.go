package main

import (
	"github.com/WhiteMaks/poly_day_bus_schema/schema"
	"github.com/WhiteMaks/poly_day_bus_schema/wbus"
	"log"
	"os"
)

var (
	rabbitConStr = os.Getenv("RABBIT_CONNECTING_STRING")
)

func main() {
	rabbit := wbus.NewRabbit("bus-schema")
	err := rabbit.Start(rabbitConStr)
	if err != nil {
		log.Fatalf("error start [ %v ]: %v", rabbit.Name, err)
	}

	//create queues for microservices. One microservice one consumer
	for _, ms := range schema.Microservices {
		err = rabbit.CreateQueue(ms)
		if err != nil {
			log.Fatalf("error create queue for ms [ %v ]: %v", ms, err)
		}
	}

	//create exchanges for events. One event many subscribers
	for _, event := range schema.Events {
		err = rabbit.CreateExchange(event)
		if err != nil {
			log.Fatalf("error create exchange for event [ %v ]: %v", event, err)
		}
	}

	//create queues for subscribers. One subscribe one event
	for exchange, queues := range schema.Subscribers {
		for _, queue := range queues {
			err = rabbit.CreateQueue(queue)
			if err != nil {
				log.Fatalf("error create queue [ %v ] for for exchange [ %v ]: %v", queue, exchange, err)
			}

			err = rabbit.LinkQueueWithExchanger(queue, exchange)
			if err != nil {
				log.Fatalf("error link queue [ %v ] with exchange [ %v ]: %v", queue, exchange, err)
			}
		}
	}
}
