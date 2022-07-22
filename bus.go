package hermessenger

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Bus struct {
	ctx             context.Context
	transport       Transport
	envelopeFactory EnvelopeFactory
}

func NewBus(
	ctx context.Context,
	optArgs ...interface{},
) *Bus {
	b := &Bus{
		ctx: ctx,
	}
	b.setOptArgs(optArgs)
	return b
}

func (b *Bus) setOptArgs(optArgs []interface{}) error {
	for _, arg := range optArgs {
		switch argType := arg.(type) {
		case Transport:
			b.transport = arg.(Transport)
		case EnvelopeFactory:
			b.envelopeFactory = arg.(EnvelopeFactory)
		default:
			log.Printf("[BUS] Unknown argument type: %T", argType)
		}
	}
	return nil
}

func (b *Bus) StartConsume() error {
	log.Println("[BUS] Consuming started")
	if b.transport == nil {
		return fmt.Errorf("[BUS] Transport not specified")
	}
	go func() {
		messageChannel := make(chan *Envelope)
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
		log.Println("[BUS] Consuming complete")
	}()
	time.Sleep(100 * time.Millisecond)
	return nil
}

func (b *Bus) Dispatch(message interface{}, options ...func(*DispatchOptions)) error {
	dispatchOptions := getDefaultDispatchOptions()
	for _, option := range options {
		option(dispatchOptions)
	}
	envelope, err := b.getEnvelopeFromMessage(message)
	if err != nil {
		return err
	}
	log.Printf("[BUS] Dispatch message: %s", envelope.stamps.template)
	b.transport.Publish(envelope, dispatchOptions)

	return nil
}

func (b *Bus) getEnvelopeFromMessage(message interface{}) (*Envelope, error) {
	if fmt.Sprintf("%T", message) == "*hermessenger.Envelope" {
		return message.(*Envelope), nil
	}
	if b.envelopeFactory == nil {
		return nil, fmt.Errorf("[BUS] EnvelopeFactory not specified")
	}
	envelope, err := b.envelopeFactory.GetEnvelope(message)
	if err != nil {
		return nil, err
	}
	return envelope, nil
}

// TODO
func (b *Bus) handleError(err error) {
	log.Println("[BUS] Error from transport:", err)
}

// TODO
func (b *Bus) handleMessage(envelope *Envelope) {
	log.Printf("[BUS] Handle message: %s", envelope.stamps.template)
	log.Printf("[BUS] Envelope: %#v", envelope)
}
