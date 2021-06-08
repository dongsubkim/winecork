package routes

import (
	"net/http"
	"os"

	"github.com/foolin/goview"
	"github.com/go-chi/chi/v5"
	"github.com/project_winecork/data"
)

func AdminRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", renderAdmin)
	r.Get("/wines", getWineDb)
	r.Get("/queries", getUserQueries)
	r.Get("/feedbacks", getFeedbacks)

	return r
}

func renderAdmin(w http.ResponseWriter, r *http.Request) {
	err := goview.Render(w, http.StatusOK, "/admin/admin", goview.M{
		"adminRoute": os.Getenv("ADMIN_ROUTE"),
	})
	if err != nil {
		warning("Error during rendering admin page: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
}

func getWineDb(w http.ResponseWriter, r *http.Request) {
	wines := data.GetAllWines()
	err := goview.Render(w, http.StatusOK, "/admin/winedb", goview.M{
		"adminRoute": os.Getenv("ADMIN_ROUTE"),
		"wines": wines,
	})
	if err != nil {
		warning("Error during rendering wine db page: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
}

func getUserQueries(w http.ResponseWriter, r *http.Request) {
	err := goview.Render(w, http.StatusOK, "/admin/userQuery", goview.M{
		"adminRoute": os.Getenv("ADMIN_ROUTE"),
	})
	if err != nil {
		warning("Error during rendering user query page: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
}

func getFeedbacks(w http.ResponseWriter, r *http.Request) {
	err := goview.Render(w, http.StatusOK, "/admin/feedback", goview.M{
		"adminRoute": os.Getenv("ADMIN_ROUTE"),
	})
	if err != nil {
		warning("Error during rendering user feedback page: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
}
