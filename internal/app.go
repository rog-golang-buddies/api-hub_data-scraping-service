package internal

import (
	"context"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/handler"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/publisher"
	"log"
)

func Start() int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.ReadConfig() //read configuration from file & env
	//initialize publisher connection to the queue
	//this library assumes using one publisher and one consumer per application
	//https://github.com/wagslane/go-rabbitmq/issues/79
	pub, err := publisher.NewPublisher(conf.QueueConfig) //TODO pass logger here and add it to publisher options
	if err != nil {
		log.Println("error while starting publisher: ", err)
		return 1
	}
	defer publisher.ClosePublisher(pub)
	//initialize consumer connection to the queue
	consumer, err := queue.NewConsumer(conf.QueueConfig) //TODO pass logger here and add it to consumer options
	if err != nil {
		log.Println("error while connecting to the queue: ", err)
		return 1
	}
	defer queue.CloseConsumer(consumer)

	handl := handler.NewApiSpecDocHandler(pub, conf.QueueConfig)
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
