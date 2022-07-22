package hermessenger

type Transport interface {
	Publish(message *Envelope, dispatchOptions *DispatchOptions) error
	Subscribe(messageChannel chan (*Envelope), errorChannel chan (error))
}
