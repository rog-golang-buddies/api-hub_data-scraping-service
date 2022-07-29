package internal

import (
	"context"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/handler"
	"log"
)

func Start() int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.ReadConfig() //read configuration from file & env
	//initialize publisher connection to the queue
	//this library assumes using one publisher and one consumer per application
	//https://github.com/wagslane/go-rabbitmq/issues/79
	publisher, err := queue.NewPublisher(conf.QueueConfig) //TODO pass logger here and add it to publisher options
	if err != nil {
		log.Println("error while starting publisher: ", err)
		return 1
	}
	defer queue.ClosePublisher(publisher)
	//initialize consumer connection to the queue
	consumer, err := queue.NewConsumer(conf.QueueConfig) //TODO pass logger here and add it to consumer options
	if err != nil {
		log.Println("error while connecting to the queue: ", err)
		return 1
	}
	defer queue.CloseConsumer(consumer)

	handl := handler.NewApiSpecDocHandler()
	listener := queue.NewListener()
	err = listener.Start(consumer, &conf.QueueConfig, handl)
	if err != nil {
		log.Println("error while listening queue ", err)
		return 1
	}

	<-ctx.Done()

	log.Println("application stopped gracefully (not)")
	return 0
}
