package envelope

import (
	"errors"
	"fmt"
)

type Envelope struct {
	message interface{}
	stamps  EnvelopeStamps
}

func New(message interface{}, stamps ...func(*EnvelopeStamps)) *Envelope {
	e := &Envelope{
		message: message,
		stamps:  make(EnvelopeStamps),
	}
	for _, stamp := range stamps {
		e.Stamp(stamp)
	}
	return e
}

func Wrap(message interface{}, stamps ...func(*EnvelopeStamps)) *Envelope {
	switch msg := message.(type) {
	case *Envelope:
		return msg
	default:
		return New(msg, stamps...)
	}
}

func (e *Envelope) Stamp(stamp func(*EnvelopeStamps)) {
	stamp(&e.stamps)
}

func (e *Envelope) GetStamp(key string) (interface{}, error) {
	if val, ok := e.stamps[key]; ok {
		return val, nil
	}
	return nil, errors.New(fmt.Sprintf("Stamp [%s] not found", key))
}

func (e *Envelope) GetMessage() interface{} {
	return e.message
}

func (e *Envelope) GetMessageType() string {
	return fmt.Sprintf("%T", e.GetMessage())
}
