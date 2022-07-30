package transport

import (
	"context"
	"log"

	"github.com/mmalessa/mmessenger/envelope"
)

const TransportName = "TransportSynchronous"

type TransportSynchronous struct {
	ctx                 context.Context
	localMessageChannel chan (*envelope.Envelope)
}

func NewSynchronous(ctx context.Context) *TransportSynchronous {
	return &TransportSynchronous{
		ctx:                 ctx,
		localMessageChannel: make(chan *envelope.Envelope),
	}
}

func (t *TransportSynchronous) Publish(message *envelope.Envelope) error {
	// log.Printf("[%s] Publish message: %s", TransportName, message.template)
	t.localMessageChannel <- message
	return nil
}

func (t *TransportSynchronous) Subscribe(
	messageChannel chan (*envelope.Envelope),
	errorChannel chan (error),
) {
	log.Printf("[%s] Start", TransportName)
out:
	for {
		select {
		case envelope := <-t.localMessageChannel:
			// log.Printf("[%s] Handle message: %s", TransportName, envelope.stamps.template)
			log.Printf("[%s] Handle message: %s", TransportName, "TODO")
			messageChannel <- envelope
		case <-t.ctx.Done():
			break out
		}
	}
	log.Printf("[%s] Completed", TransportName)
}
