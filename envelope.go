package mmessenger

type Envelope struct {
	message interface{}
	stamps  *EnvelopeStamps
}

func NewEnvelope(message interface{}, stamps ...func(*EnvelopeStamps)) *Envelope {
	e := &Envelope{
		message: message,
	}
	e.stamps = getDefaultEnvelopeStamps()
	for _, stamp := range stamps {
		stamp(e.stamps)
	}
	return e
}

func (e *Envelope) Stamp(stamp func(*EnvelopeStamps)) {
	stamp(e.stamps)
}
