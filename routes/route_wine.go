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
	r.Get("/suggestion", renderSuggestion)
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
	r.ParseForm()
	log.Println(r.PostForm)
	http.Redirect(w, r, "/suggestion", http.StatusFound)
}

func renderSuggestion(w http.ResponseWriter, r *http.Request) {
	err := goview.Render(w, http.StatusOK, "wine", goview.M{
		"wine": r.PostFormValue("wine"),
		"store": r.PostFormValue("store"),
		"price": r.PostFormValue("price"),
		"food": r.PostFormValue("food"),
	})
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}
}