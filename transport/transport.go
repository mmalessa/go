package transport

import "github.com/mmalessa/mmessenger/envelope"

type Transport interface {
	Publish(message *envelope.Envelope) error
	Subscribe(messageChannel chan (*envelope.Envelope), errorChannel chan (error))
}
