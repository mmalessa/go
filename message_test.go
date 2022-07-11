package hermessenger

import "testing"

func TestNewMessage(t *testing.T) {
	body := "My Content"
	msg := NewMessage(body)

	if msg.body != body {
		t.Fatalf("got %s, wanted %s", msg.body, body)
	}
}
