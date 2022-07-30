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
	// b.setDefaultArgs()
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

// func (b *MessageBus) setDefaultArgs() error {
// 	if b.transport == nil {
// 		b.transport = transport.NewSynchronous(b.ctx)
// 	}
// 	// TODO
// 	return nil
// }

func (b *MessageBus) Start() error {
	log.Println("[messagebus] Start")
	go func() {
		messageChannel := make(chan *envelope.Envelope)
		errorChannel := make(chan error)
		defer close(messageChannel)
		defer close(errorChannel)
		go b.transport.Subscribe(messageChannel, errorChannel)
	out:
		for {
			select {
			case msg := <-messageChannel:
				b.handleMessage(msg)
			case err := <-errorChannel:
				b.handleError(err)
			case <-b.ctx.Done():
				break out
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
		log.Println("[messagebus] Complete")
	}()
	time.Sleep(100 * time.Millisecond)
	return nil
}

// TODO
func (b *MessageBus) handleError(err error) {
	log.Println("[BUS] Error from transport:", err)
}

// TODO
func (b *MessageBus) handleMessage(envelope *envelope.Envelope) {
	// log.Printf("[BUS] Handle message: %s", envelope.stamps.template)
	log.Printf("[BUS] Envelope: %#v", envelope)
}

func (b *MessageBus) Dispatch(message interface{}, stamps ...func(*envelope.EnvelopeStamps)) error {
	envelope := envelope.Wrap(message, stamps...)
	fmt.Println(envelope)
	log.Printf("[BUS] Dispatch message: %s", "TODO")
	return b.transport.Publish(envelope)
}
