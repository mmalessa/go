package hermessenger

import "log"

type Bus struct {
	transport Transport
}

func NewBus() *Bus {
	b := &Bus{}
	return b
}

func (b *Bus) Dispatch(message *Message, options ...func(*DispatchOptions)) error {
	dispatchOptions := getDefaultDispatchOptions()
	for _, option := range options {
		option(dispatchOptions)
	}
	// TODO
	log.Printf("Dispatch message: %#v\n", message)
	log.Printf(" ...with options: %#v\n", dispatchOptions)

	return nil
}

// --------------
type Transport interface {
}
