package internal

import (
	"context"
	"fmt"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/logger"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/handler"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/publisher"
)

func Start() int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.ReadConfig() //read configuration from file & env
	log, err := logger.NewLogger(&conf)
	if err != nil {
		fmt.Println("error creating logger: ", err)
		return 1
	}
	//initialize publisher connection to the queue
	//this library assumes using one publisher and one consumer per application
	//https://github.com/wagslane/go-rabbitmq/issues/79
	pub, err := publisher.NewPublisher(conf.QueueConfig, log)
	if err != nil {
		log.Error("error while starting publisher: ", err)
		return 1
	}
	defer publisher.ClosePublisher(pub, log)
	//initialize consumer connection to the queue
	consumer, err := queue.NewConsumer(conf.QueueConfig, log)
	if err != nil {
		log.Error("error while connecting to the queue: ", err)
		return 1
	}
	defer queue.CloseConsumer(consumer, log)

	handl := handler.NewApiSpecDocHandler(pub, conf.QueueConfig, log)
	listener := queue.NewListener()
	err = listener.Start(consumer, &conf.QueueConfig, handl)
	if err != nil {
		log.Error("error while listening queue ", err)
		return 1
	}

	<-ctx.Done()

	log.Info("application stopped gracefully (not)")
	return 0
}
