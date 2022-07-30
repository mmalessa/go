package envelope

import "testing"

func TestNewMessage(t *testing.T) {
	message := "My Content"
	msg := New(message)

	if msg.message != message {
		t.Fatalf("got %s, wanted %s", msg.message, message)
	}
}
