package hermessenger

import (
	"context"
	"log"
	"time"
)

const TransportName = "TransportSynchronous"

type TransportSynchronous struct {
	ctx                 context.Context
	localMessageChannel chan (*Envelope)
}

func NewTransportSynchronous(ctx context.Context) *TransportSynchronous {
	return &TransportSynchronous{
		ctx:                 ctx,
		localMessageChannel: make(chan *Envelope),
	}
}

func (t *TransportSynchronous) Publish(message *Envelope, dispatchOptions *DispatchOptions) error {
	// log.Printf("[%s] Publish message: %s", TransportName, message.template)
	t.localMessageChannel <- message
	return nil
}

func (t *TransportSynchronous) Subscribe(
	messageChannel chan (*Envelope),
	errorChannel chan (error),
) {
	log.Printf("[%s] Subscribing started", TransportName)
out:
	for {
		select {
		case msg := <-t.localMessageChannel:
			log.Printf("[%s] Handle message: %s", TransportName, "FIXME")
			messageChannel <- msg
		case <-t.ctx.Done():
			break out
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
	log.Printf("[%s] Subscribing complete", TransportName)
}
