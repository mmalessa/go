package mmessenger

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mmalessa/mmessenger/envelope"
	"github.com/mmalessa/mmessenger/transport"
)

type MessageBus struct {
	ctx       context.Context
	transport transport.Transport
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
		default:
			log.Printf("[messagebus] Unknown argument type: %T", argTyped)
		}
	}
	return nil
}

func (b *MessageBus) Dispatch(message interface{}, stamps ...func(*envelope.EnvelopeStamps)) error {
	envelope := envelope.Wrap(message, stamps...)
	log.Printf("[messagebus] Dispatch message: %#v", envelope)
	log.Printf("[messagebus] Dispatch to transport: %s", fmt.Sprintf("%T", b.transport))
	return b.transport.Publish(envelope)
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
			case msg := <-busMessageChannel:
				busErrorChannel <- b.handleMessage(msg)
			case <-b.ctx.Done():
				break out
			}
		}
		log.Println("[messagebus] Completed")
	}()
	time.Sleep(100 * time.Millisecond)
}

// TODO
func (b *MessageBus) handleMessage(envelope *envelope.Envelope) error {
	log.Printf("[messagebus] Handle message: %#v", envelope)
	// TODO
	time.Sleep(333 * time.Millisecond)
	log.Printf("[messagebus] ACK message: %#v", envelope)
	return nil
}
