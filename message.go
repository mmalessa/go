package hermessenger

type Message struct {
	body string
}

func NewMessage(body string) *Message {
	m := &Message{
		body: body,
	}
	return m
}
