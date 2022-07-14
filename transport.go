package hermessenger

type Transport interface {
	Publish(message *Message, dispatchOptions *DispatchOptions) error
	Subscribe(messageChannel chan (*Message), errorChannel chan (error))
}
