package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/foolin/goview"
	"github.com/go-chi/chi/v5"
	"github.com/project_winecork/data"
)

func WineRouter(r chi.Router) {
	r.Get("/", renderIndex)
	r.Post("/", queryWine)
	r.Get("/detail", renderDetail)
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
	wines := []data.Wine{
		data.Wine{
			Id:       1,
			WineName: "풋 프린트 쉬라즈 750mL",
			Price:    22000,
			Image:    "static/icons/EM0001.jpg",
		},
		data.Wine{
			Id:       2,
			WineName: "풋 프린트 쉬라즈 750mL",
			Price:    22000,
			Image:    "static/icons/EM0006.jpg",
		},
	}
	err := goview.Render(w, http.StatusOK, "recommendation", goview.M{
		"wines": wines,
	})
	if err != nil {
		log.Println(err)
	}
}

func renderDetail(w http.ResponseWriter, r *http.Request) {
	log.Println("Rendering details...")
	log.Println("Wine Id: ", r.FormValue("id"))
	wine, err := data.WineById(r.FormValue("id"))
	if err != nil {
		log.Println("Fail to get wine by id: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
	err = goview.Render(w, http.StatusOK, "selection", goview.M{
		"wine": wine,
	})
	if err != nil {
		log.Println("Error during rendering detail: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
}
