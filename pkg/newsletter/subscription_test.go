package newsletter

import (
    "testing"
)

func TestSubscribe(t *testing.T) {
    email := "test@example.com"
    if err := Subscribe(email); err != nil {
        t.Errorf("Subscribe failed: %v", err)
    }
}

func TestUnsubscribe(t *testing.T) {
    email := "test@example.com"
    if err := Unsubscribe(email); err != nil {
        t.Errorf("Unsubscribe failed: %v", err)
    }
}
