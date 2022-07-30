package envelope

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
