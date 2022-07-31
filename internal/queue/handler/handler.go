package handler

import "github.com/wagslane/go-rabbitmq"

//Handler represents common interface to any queue message processing struct
//go:generate mockgen -source=handler.go -destination=./mocks/handler.go -package=handler
type Handler interface {
	//Handle message and return action to response to the queue
	Handle(delivery rabbitmq.Delivery) rabbitmq.Action
}
