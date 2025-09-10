package handler

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	subscriptionHandler *SubscriptionHandler
}

func New(subscriptionService Subscription) *Handler {
	return &Handler{
		subscriptionHandler: NewSubscriptionHandler(subscriptionService),
	}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.subscriptionHandler.Create(w, r)
		case http.MethodGet:
			h.subscriptionHandler.List(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/subscriptions/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.subscriptionHandler.Read(w, r)
		case http.MethodPut:
			h.subscriptionHandler.Update(w, r)
		case http.MethodDelete:
			h.subscriptionHandler.Delete(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/subscriptions/total", h.subscriptionHandler.Total)
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	return mux
}
