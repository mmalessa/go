package hermessenger

import "fmt"

type EnvelopeFactory interface {
	GetEnvelope(message interface{}) (*Envelope, error)
}

// default
type EnvelopeFactoryDefault struct {
}

func NewEnvelopeFactoryDefault() *EnvelopeFactoryDefault {
	ef := &EnvelopeFactoryDefault{}
	return ef
}

func (ef *EnvelopeFactoryDefault) GetEnvelope(message interface{}) (*Envelope, error) {
	template := fmt.Sprintf("%T", message)
	envelope := NewEnvelope(message)
	envelope.Stamp(EnvelopeStampWithTemplate(template))
	return envelope, nil
}
