package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func FeedbackRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", submitFeedback)
	return r
}

func submitFeedback(w http.ResponseWriter, r *http.Request) {
	info("Submitting Feedback...")
	info(r.FormValue("feedbackBody"))
	w.WriteHeader(http.StatusNoContent)
}
