package transportsynchronous

import (
	"context"
	"log"

	"github.com/mmalessa/mmessenger/envelope"
)

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

func (t *TransportSynchronous) Publish(envel *envelope.Envelope) error {
	log.Printf("[transport synchronous] Publish message: %v", envel)
	t.localMessageChannel <- envel
	return <-t.localErrorChannel
}

func (t *TransportSynchronous) Subscribe(
	busMessageChannel chan (*envelope.Envelope),
	busErrorChannel chan (error),
) {
	log.Print("[transport synchronous] Start")
endfor:
	for {
		select {
		case envel := <-t.localMessageChannel:
			log.Printf("[transport synchronous] Received message: %v", envel)
			busMessageChannel <- envel
			t.localErrorChannel <- <-busErrorChannel
		case <-t.ctx.Done():
			break endfor
		}
	}
	log.Print("[transport synchronous] Completed")
}
