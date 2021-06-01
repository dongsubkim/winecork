package routes

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AdminRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", renderAdmin)
	return r
}

func renderAdmin(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading Admin page...")
}
