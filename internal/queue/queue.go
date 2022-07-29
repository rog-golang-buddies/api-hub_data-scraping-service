package queue

import (
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/wagslane/go-rabbitmq"
	"io"
	"log"
)

type Consumer interface {
	io.Closer
	StartConsuming(
		handler rabbitmq.Handler,
		queue string,
		routingKeys []string,
		optionFuncs ...func(*rabbitmq.ConsumeOptions),
	) error
}

type Publisher interface {
	io.Closer
	Publish(
		data []byte,
		routingKeys []string,
		optionFuncs ...func(*rabbitmq.PublishOptions),
	) error
}

func NewConsumer(conf config.QueueConfig) (Consumer, error) {
	consumer, err := rabbitmq.NewConsumer(
		conf.Url,
		rabbitmq.Config{},
		rabbitmq.WithConsumerOptionsLogging,
	)
	if err != nil {
		return nil, err
	}
	return &consumer, nil
}

func CloseConsumer(consumer Consumer) {
	log.Println("closing consumer")
	err := consumer.Close()
	if err != nil {
		log.Println("error while closing consumer: ", err)
	}
}

func NewPublisher(conf config.QueueConfig) (Publisher, error) {
	return rabbitmq.NewPublisher(
		conf.Url,
		rabbitmq.Config{},
	)
}

func ClosePublisher(publisher Publisher) {
	log.Println("closing publisher")
	err := publisher.Close()
	if err != nil {
		log.Println("error while closing publisher: ", err)
	}
}
