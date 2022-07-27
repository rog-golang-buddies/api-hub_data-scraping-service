package queue

import (
	"errors"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/handler"
	"github.com/wagslane/go-rabbitmq"
	"log"
)

//Listener represents consumer wrapper with the method to start listening for all events for this service
type Listener interface {
	Start() error
}

type ApiSpecListenerImpl struct {
	config  *config.QueueConfig //configuration struct
	handler handler.Handler
}

func (asl *ApiSpecListenerImpl) Start() error {
	if asl.config == nil {
		return errors.New("queue configuration must not be nil")
	}
	consumer, err := rabbitmq.NewConsumer(
		asl.config.Url,
		rabbitmq.Config{},
		rabbitmq.WithConsumerOptionsLogging,
	)
	if err != nil {
		return err
	}
	defer func() {
		log.Printf("closing consumer")
		err := consumer.Close()
		if err != nil {
			log.Println("error while closing consumer: ", err)
		}
	}()
	err = consumer.StartConsuming(
		asl.handler.Handle,
		asl.config.ConsumerQueue,
		[]string{}, //No binding, consuming with the default exchange directly by queue name
		rabbitmq.WithConsumeOptionsConcurrency(asl.config.Concurrency),
		rabbitmq.WithConsumeOptionsQueueDurable,
		rabbitmq.WithConsumeOptionsBindingExchangeDurable,
	)

	if err != nil {
		return err
	}
	log.Println("listening stopped")

	return nil
}

func NewListener(config *config.QueueConfig, handler handler.Handler) Listener {
	return &ApiSpecListenerImpl{
		config:  config,
		handler: handler,
	}
}
