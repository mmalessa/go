package mmessenger

import "github.com/mmalessa/mmessenger/envelope"

type Transport interface {
	Publish(message *envelope.Envelope, dispatchOptions *DispatchOptions) error
	Subscribe(messageChannel chan (*envelope.Envelope), errorChannel chan (error))
}
