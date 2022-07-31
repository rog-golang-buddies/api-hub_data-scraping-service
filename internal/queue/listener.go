package queue

import (
	"errors"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/handler"
	"github.com/wagslane/go-rabbitmq"
)

//Listener represents consumer wrapper with the method to start listening for all events for this service
//go:generate mockgen -source=listener.go -destination=./mocks/listener.go
type Listener interface {
	//Start listening queues
	Start(
		consumer Consumer,
		config *config.QueueConfig,
		handler handler.Handler,
	) error
}

type ListenerImpl struct {
}

func (listener *ListenerImpl) Start(
	consumer Consumer,
	config *config.QueueConfig,
	handler handler.Handler,
) error {
	if consumer == nil {
		return errors.New("queue consumer must not be nil")
	}
	if config == nil {
		return errors.New("configuration must not be nil")
	}

	err := consumer.StartConsuming(
		handler.Handle,
		config.UrlRequestQueue,
		[]string{}, //No binding, consuming with the default exchange directly by queue name
		rabbitmq.WithConsumeOptionsConcurrency(config.Concurrency),
		rabbitmq.WithConsumeOptionsQueueDurable,
	)
	if err != nil {
		return err
	}

	return nil
}

func NewListener() Listener {
	return &ListenerImpl{}
}
