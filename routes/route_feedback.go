package routes

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func FeedbackRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", submitFeedback)
	return r
}

func submitFeedback(w http.ResponseWriter, r *http.Request) {
	log.Println("Submitting Feedback...")
}
