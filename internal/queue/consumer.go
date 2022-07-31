package queue

import (
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/wagslane/go-rabbitmq"
	"io"
	"log"
)

//Consumer is just an interface for the library consumer which doesn't have one.
//go:generate mockgen -source=consumer.go -destination=./mocks/consumer.go
type Consumer interface {
	io.Closer
	StartConsuming(
		handler rabbitmq.Handler,
		queue string,
		routingKeys []string,
		optionFuncs ...func(*rabbitmq.ConsumeOptions),
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
