package routes

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

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
	r.Post("/upload", uploadCsv)
	r.Get("/download", downloadDB)
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
		"wines":      wines,
	})
	if err != nil {
		warning("Error during rendering wine db page: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
}

func getUserQueries(w http.ResponseWriter, r *http.Request) {
	quereis := data.GetAllQueries()
	err := goview.Render(w, http.StatusOK, "/admin/userQuery", goview.M{
		"adminRoute": os.Getenv("ADMIN_ROUTE"),
		"queries":    quereis,
	})
	if err != nil {
		warning("Error during rendering user query page: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
}

func getFeedbacks(w http.ResponseWriter, r *http.Request) {
	feedbacks := data.GetAllFeedbacks()
	err := goview.Render(w, http.StatusOK, "/admin/feedback", goview.M{
		"adminRoute": os.Getenv("ADMIN_ROUTE"),
		"feedbacks":  feedbacks,
	})
	if err != nil {
		warning("Error during rendering user feedback page: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
}

func uploadCsv(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 32)
	fileHeader := r.MultipartForm.File["csv"][0]
	file, err := fileHeader.Open()
	if err != nil {
		log.Println("Fail to open file")
		return
	}
	defer file.Close()
	err = data.SaveCSV(file)
	if err != nil {
		danger("Error during SaveCSV:", err)
	}
	http.Redirect(w, r, "/", http.StatusNoContent)
}

func downloadDB(w http.ResponseWriter, r *http.Request) {
	info("Downloading Wine DB...")
	wines := data.GetAllWines()
	csvWriter := csv.NewWriter(w)
	header := []string{"Priority", "Key", "Store", "Wine_Name", "Location", "Price", "Wine_Sort", "Wine_Origin", "Grape_1", "Grape_2", "Grape_3", "Acid", "Sweet", "Sparkling", "Food_match"}
	err := csvWriter.Write(header)
	if err != nil {
		warning("Fail to create csv file:", err)
		return
	}
	for _, wine := range wines {
		grapes := []string{"", "", ""}
		rawGrapes := strings.Split(data.StringfyPqArray(wine.Grapes), ", ")
		copy(grapes, rawGrapes)

		line := []string{
			strconv.Itoa(wine.Priority),
			wine.Key,
			wine.Store,
			wine.WineName,
			data.StringfyPqArray(wine.Locations),
			strconv.Itoa(wine.Price),
			wine.WineType,
			wine.Country,
			grapes[0],
			grapes[1],
			grapes[2],
			strconv.Itoa(wine.Acidity),
			strconv.Itoa(wine.Sweetness),
			strconv.Itoa(wine.Sparkling),
			data.StringfyPqArray(wine.FoodMatches),
		}
		err := csvWriter.Write(line)
		if err != nil {
			warning("Fail to write a line to csv file: ", err)
		}

	}
	csvWriter.Flush()
}
