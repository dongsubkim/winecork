package data

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/admin/search"
	"github.com/joho/godotenv"
)

var logger *log.Logger
var cld *cloudinary.Cloudinary

func initLogger() {
	logger = log.Default()
	logger.SetFlags(log.Ldate)
	logger.SetFlags(log.Ltime)

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	cld, err = cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_KEY"), os.Getenv("CLOUDINARY_SECRET"))
	if err != nil {
		log.Println("Error loading cloudinary api")
	}
}

// for logging
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Fatalln(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix(("WARNING "))
	logger.Println(args...)
}

func getImageUrl() map[string]string {
	var ctx = context.Background()
	searchResult, err := cld.Admin.Search(ctx, search.Query{Expression: "folder:WineCork/*"})
	if err != nil {
		danger("Failed to get search result:", err)
	}
	urls := make(map[string]string)
	storeUrls := func(sr *admin.SearchResult) {
		for _, image := range sr.Assets {
			id := image.PublicID
			id = strings.Split(id, "_")[0][len("WineCork/"):]
			url := image.SecureURL
			url = strings.Replace(url, "/upload/", "/upload/e_trim:10/", 1)
			urls[id] = url
		}
	}
	storeUrls(searchResult)
	nextCursor := searchResult.NextCursor

	for len(nextCursor) > 0 {
		sr, err := cld.Admin.Search(ctx, search.Query{Expression: "folder:WineCork/*", NextCursor: nextCursor})
		if err != nil {
			danger("Failed to get search result:", err)
		}
		storeUrls(sr)
		nextCursor = sr.NextCursor
	}

	return urls
}
