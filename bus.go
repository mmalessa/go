package hermessenger

import (
	"context"
	"log"
	"time"
)

type Bus struct {
	ctx       context.Context
	transport Transport
}

func NewBus(
	ctx context.Context,
	transport Transport,
) *Bus {
	b := &Bus{
		ctx:       ctx,
		transport: transport,
	}
	return b
}

func (b *Bus) Dispatch(message *Message, options ...func(*DispatchOptions)) error {
	dispatchOptions := getDefaultDispatchOptions()
	for _, option := range options {
		option(dispatchOptions)
	}

	b.transport.Publish(message, dispatchOptions)

	return nil
}

func (b *Bus) StartConsume() {
	log.Println("[BUS] Consuming started")
	go func() {
		messageChannel := make(chan *Message)
		errorChannel := make(chan error)
		defer close(messageChannel)
		defer close(errorChannel)
		go b.transport.Subscribe(messageChannel, errorChannel)
	out:
		for {
			select {
			case msg := <-messageChannel:
				log.Printf("[BUS] Handle message: %T", *msg)
				//TODO
			case err := <-errorChannel:
				log.Println("[BUS] Error from transport:", err)
			case <-b.ctx.Done():
				break out
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}

		log.Println("[BUS] Consuming complete")
	}()
	time.Sleep(50 * time.Millisecond)
}
