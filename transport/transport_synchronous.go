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
	localErrorChannel   chan (error)
}

func NewSynchronous(ctx context.Context) *TransportSynchronous {
	return &TransportSynchronous{
		ctx:                 ctx,
		localMessageChannel: make(chan *envelope.Envelope),
		localErrorChannel:   make(chan error),
	}
}

func (t *TransportSynchronous) Publish(e *envelope.Envelope) error {
	log.Printf("[%s] Publish message: %v", TransportName, e)
	t.localMessageChannel <- e
	return <-t.localErrorChannel
}

func (t *TransportSynchronous) Subscribe(
	busMessageChannel chan (*envelope.Envelope),
	busErrorChannel chan (error),
) {
	log.Printf("[%s] Start", TransportName)
endfor:
	for {
		select {
		case e := <-t.localMessageChannel:
			log.Printf("[%s] Handle message: %v", TransportName, e)
			busMessageChannel <- e
			t.localErrorChannel <- <-busErrorChannel
		case <-t.ctx.Done():
			break endfor
		}
	}
	log.Printf("[%s] Completed", TransportName)
}
