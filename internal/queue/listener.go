package queue

import (
	"context"
	"errors"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/handler"
	"github.com/wagslane/go-rabbitmq"
	"log"
)

//Listener represents consumer wrapper with the method to start listening for all events for this service
type Listener interface {
	//Start listening queues
	Start(ctx context.Context) error
}

type ListenerImpl struct {
	config  *config.QueueConfig //configuration struct
	handler handler.Handler
}

func (listener *ListenerImpl) Start(ctx context.Context) error {
	if listener.config == nil {
		return errors.New("queue configuration must not be nil")
	}
	consumer, err := rabbitmq.NewConsumer(
		listener.config.Url,
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
		listener.handler.Handle,
		listener.config.ConsumerQueue,
		[]string{}, //No binding, consuming with the default exchange directly by queue name
		rabbitmq.WithConsumeOptionsConcurrency(listener.config.Concurrency),
		rabbitmq.WithConsumeOptionsQueueDurable,
	)

	if err != nil {
		return err
	}

	<-ctx.Done()
	log.Println("listening stopped")

	return nil
}

func NewListener(config *config.QueueConfig, handler handler.Handler) Listener {
	return &ListenerImpl{
		config:  config,
		handler: handler,
	}
}
