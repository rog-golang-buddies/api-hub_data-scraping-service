package handler

import "github.com/wagslane/go-rabbitmq"

//Handler represents common interface to any queue message processing struct
type Handler interface {
	//Handle message and return action to response to the queue
	Handle(delivery rabbitmq.Delivery) rabbitmq.Action
}
