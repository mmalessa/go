package handlerslocator

import (
	"fmt"

	"github.com/mmalessa/mmessenger/envelope"
)

type HandlersLocator interface {
	GetHandler(envel *envelope.Envelope) (func(envel *envelope.Envelope) error, error)
}

type HandlersLocatorDefault struct {
}

func NewHandlersLocatorDefault() *HandlersLocatorDefault {
	return &HandlersLocatorDefault{}
}

func (hl *HandlersLocatorDefault) GetHandler(envel *envelope.Envelope) (func(envel *envelope.Envelope) error, error) {
	return func(envel *envelope.Envelope) error {
		fmt.Println("Hi, I'm default handler")
		return nil
	}, nil
}
