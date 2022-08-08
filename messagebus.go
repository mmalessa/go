package mmessenger

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mmalessa/mmessenger/envelope"
	handlerslocator "github.com/mmalessa/mmessenger/handlers_locator"
	"github.com/mmalessa/mmessenger/transport"
)

type MessageBus struct {
	ctx             context.Context
	transport       transport.Transport
	handlersLocator handlerslocator.HandlersLocator
}

func NewMessageBus(
	ctx context.Context,
	optArgs ...interface{},
) *MessageBus {
	b := &MessageBus{
		ctx: ctx,
	}
	b.setOptArgs(optArgs)
	return b
}

func (b *MessageBus) setOptArgs(optArgs []interface{}) error {
	for _, arg := range optArgs {
		switch argTyped := arg.(type) {
		case transport.Transport:
			b.transport = argTyped
		case handlerslocator.HandlersLocator:
			b.handlersLocator = argTyped
		default:
			log.Printf("[messagebus] Unknown argument type: %T", argTyped)
		}
	}
	return nil
}

func (b *MessageBus) Dispatch(message interface{}, stamps ...func(*envelope.EnvelopeStamps)) error {
	envel := envelope.Wrap(message, stamps...)
	log.Printf("[messagebus] Dispatch message: %#v", envel)
	log.Printf("[messagebus] Dispatch to transport: %s", fmt.Sprintf("%T", b.transport))
	return b.transport.Publish(envel)
}

func (b *MessageBus) Start() {
	log.Println("[messagebus] Start")
	go func() {
		busMessageChannel := make(chan *envelope.Envelope)
		busErrorChannel := make(chan error)
		defer close(busMessageChannel)
		defer close(busErrorChannel)
		go b.transport.Subscribe(busMessageChannel, busErrorChannel)
	out:
		for {
			select {
			case envel := <-busMessageChannel:
				log.Printf("[messagebus] Process the message: %#v", envel)
				busErrorChannel <- b.processTheMessage(envel)
				log.Printf("[messagebus] Message has been processed: %#v", envel)
			case <-b.ctx.Done():
				break out
			}
		}
		log.Println("[messagebus] Completed")
	}()
	time.Sleep(100 * time.Millisecond)
}

// TODO
func (b *MessageBus) processTheMessage(envel *envelope.Envelope) error {
	handler, err := b.handlersLocator.GetHandler(envel)
	if err != nil {
		// FIXME
		return err
	}
	result := handler(envel)
	if result != nil {
		log.Printf("[messagebus] handler ERROR: %#v", result)
	} else {
		log.Print("[messagebus] handler success")
	}
	// retry policy
	// failed message processing

	time.Sleep(333 * time.Millisecond)
	return result
}
