package hermessenger

import (
	"context"
	"log"
	"time"
)

const TransportName = "TransportSynchronous"

type TransportSynchronous struct {
	ctx                 context.Context
	localMessageChannel chan (*Message)
}

func NewTransportDev(ctx context.Context) *TransportSynchronous {
	return &TransportSynchronous{
		ctx:                 ctx,
		localMessageChannel: make(chan *Message),
	}
}

func (t *TransportSynchronous) Publish(message *Message, dispatchOptions *DispatchOptions) error {
	log.Printf("[%s] Publish message %T", TransportName, *message)
	t.localMessageChannel <- message
	return nil
}

func (t *TransportSynchronous) Subscribe(
	messageChannel chan (*Message),
	errorChannel chan (error),
) {
	log.Printf("[%s] Subscribing started", TransportName)
out:
	for {
		select {
		case msg := <-t.localMessageChannel:
			log.Printf("[%s] Handle message: %T", TransportName, *msg)
			messageChannel <- msg
		case <-t.ctx.Done():
			break out
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
	log.Printf("[%s] Subscribing complete", TransportName)
}
