package internal

import (
	"context"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/load"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/parse"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/process"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/handler"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/publisher"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/recognize"
	"log"
)

func Start() int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf, err := config.ReadConfig() //read configuration from file & env
	if err != nil {
		log.Println("error while reading configuration")
		return 1
	}

	proc, err := createDefaultProcessor()
	if err != nil {
		log.Println("error while creating processor: ", err)
		return 1
	}
	//initialize publisher connection to the queue
	//this library assumes using one publisher and one consumer per application
	//https://github.com/wagslane/go-rabbitmq/issues/79
	pub, err := publisher.NewPublisher(conf.Queue) //TODO pass logger here and add it to publisher options
	if err != nil {
		log.Println("error while starting publisher: ", err)
		return 1
	}
	defer publisher.ClosePublisher(pub)
	//initialize consumer connection to the queue
	consumer, err := queue.NewConsumer(conf.Queue) //TODO pass logger here and add it to consumer options
	if err != nil {
		log.Println("error while connecting to the queue: ", err)
		return 1
	}
	defer queue.CloseConsumer(consumer)

	handl := handler.NewApiSpecDocHandler(pub, conf.Queue, proc)
	listener := queue.NewListener()
	err = listener.Start(ctx, consumer, &conf.Queue, handl)
	if err != nil {
		log.Println("error while listening queue ", err)
		return 1
	}

	<-ctx.Done()

	log.Println("application stopped gracefully (not)")
	return 0
}

func createDefaultProcessor() (process.UrlProcessor, error) {
	recognizer := recognize.NewRecognizer()
	parsers := []parse.Parser{parse.NewJsonOpenApiParser(), parse.NewYamlOpenApiParser()}
	converter := parse.NewConverter(parsers)
	loader := load.NewContentLoader()

	return process.NewProcessor(recognizer, converter, loader)
}
