package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/foolin/goview"
	"github.com/go-chi/chi/v5"
)

func WineRouter(r chi.Router) {
	r.Get("/", renderIndex)
	r.Post("/", queryWine)
}

func renderIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("Rendering index page...")
	err := goview.Render(w, http.StatusOK, "index", goview.M{
		"key": os.Getenv("MAP_API_KEY"),
	})
	if err != nil {
		return
	}
}

func queryWine(w http.ResponseWriter, r *http.Request) {
	log.Println("Querying Wine...")
}
