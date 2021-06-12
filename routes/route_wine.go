package routes

import (
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
	info("Rendering index page...")
	err := goview.Render(w, http.StatusOK, "index", goview.M{
		"key":            os.Getenv("MAP_API_KEY"),
		"storeLocations": data.GetStoreLocations(),
	})
	if err != nil {
		return
	}
}

func queryWine(w http.ResponseWriter, r *http.Request) {
	info("Querying Wine...")
	wines, err := data.QueryWines(r.FormValue("store"), r.FormValue("wine_type"), r.FormValue("food_match"), r.FormValue("price_range"))
	if err != nil {
		warning("ERROR during queryWine:", err)
	}
	if len(wines) == 0 {
		warning("No matching wines with:", r.Form)
	}
	err = goview.Render(w, http.StatusOK, "recommendation", goview.M{
		"wines":        wines,
		"wineNotFound": len(wines) == 0,
	})
	if err != nil {
		warning(err)
	}
}

func renderDetail(w http.ResponseWriter, r *http.Request) {
	info("Rendering details...")
	wine, err := data.WineById(r.FormValue("id"))
	if err != nil {
		warning("Fail to get wine by id: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
	err = goview.Render(w, http.StatusOK, "selection", goview.M{
		"wine":           wine,
		"convertedPrice": wine.ConvertedPrice(),
		"wineDetail":     wine.GetWineInfo(),
		"wineImage":      wine.GetImage(300),
	})
	if err != nil {
		warning("Error during rendering detail: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
}
