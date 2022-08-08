package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/mmalessa/mmessenger"
	"github.com/mmalessa/mmessenger/envelope"
	handlerslocator "github.com/mmalessa/mmessenger/handlers_locator"
	transportsynchronous "github.com/mmalessa/mmessenger/transport/transport_synchronous"
)

func main() {

	ctx, endApp := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)

	transport := transportsynchronous.NewSynchronous(ctx)
	handersLocator := handlerslocator.NewHandlersLocatorDefault()
	bus := mmessenger.NewMessageBus(ctx, transport, handersLocator)

	bus.Start()

	msg := "Sophie"
	bus.Dispatch(msg, envelope.StampWithDelay(1))

	endApp()
}
