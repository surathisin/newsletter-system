package api

import (
    "net/http"
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/subscribe", SubscribeHandler).Methods("POST")
    router.HandleFunc("/unsubscribe", UnsubscribeHandler).Methods("POST")
    return router
}
