package api

import (
	"encoding/json"
	"net/http"
	"newsletter-system/pkg/newsletter"
)

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	var req newsletter.SubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := newsletter.Subscribe(req.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Subscribed successfully"))
}

func UnsubscribeHandler(w http.ResponseWriter, r *http.Request) {
	var req newsletter.SubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := newsletter.Unsubscribe(req.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Unsubscribed successfully"))
}
