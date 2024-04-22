package newsletter

import "newsletter-system/internal/storage"

type SubscriptionRequest struct {
    Email string `json:"email"`
}

func Subscribe(email string) error {
    return storage.AddEmail(email)
}

func Unsubscribe(email string) error {
    return storage.RemoveEmail(email)
}
