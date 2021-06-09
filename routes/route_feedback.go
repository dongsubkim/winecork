package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/project_winecork/data"
)

func FeedbackRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", submitFeedback)
	return r
}

func submitFeedback(w http.ResponseWriter, r *http.Request) {
	info("Submitting Feedback...")
	info(r.FormValue("feedbackBody"))
	body := r.FormValue("feedbackBody")
	w.WriteHeader(http.StatusNoContent)
	if len(body) == 0 {
		return
	}
	err := data.AddFeedback(body)
	if err != nil {
		warning("Error during submitFeedback:", err)
	}
}
