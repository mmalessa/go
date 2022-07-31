package transport

import "github.com/mmalessa/mmessenger/envelope"

type Transport interface {
	Publish(message *envelope.Envelope) error
	Subscribe(busMessageChannel chan (*envelope.Envelope), busErrorChannel chan (error))
}
