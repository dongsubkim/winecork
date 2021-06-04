package data

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/lib/pq"
)

type Wine struct {
	Id          int
	Priority    int
	Key         string `db:"key_id"`
	Store       string
	WineName    string `db:"wine_name"`
	Locations   *pq.StringArray
	Price       int
	PriceType   int    `db:"price_type"`
	WineType    string `db:"wine_type"`
	Country     string
	Grapes      *pq.StringArray
	Acidity     int
	Sweetness   int
	Sparkling   int
	FoodMatches *pq.StringArray `db:"food_matches"`
	Image       string          `db:"image_url"`
	CreatedAt   time.Time       `db:"created_at"`
}

func init() {
	csvFile, err := os.Open(path.Join("data", "wine_db.csv"))
	if err != nil {
		log.Fatal("Error opening wine_db.csv: ", err)
	}
	defer csvFile.Close()
	wines := parseCSV(csvFile)
	err = store(wines)
	if err != nil {
		log.Fatal("Error during storing csv")
	}
}

func mappingPrice(price int) int {
	if price < 10000 {
		return 0
	} else if price < 20000 {
		return 1
	} else if price < 30000 {
		return 2
	} else if price < 40000 {
		return 3
	} else {
		return 4
	}
}

func parseCSV(csvFile io.Reader) []Wine {
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var wines []Wine

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// Priority,Key,Store,Wine_Name,Location,Price,Wine_Sort,Wine_Origin,Grape_1,Grape_2,Grape_3,Acid,Sweet,Sparkling,Food_match
		wine := Wine{
			Key:       line[1],
			Store:     line[2],
			WineName:  line[3],
			WineType:  line[6],
			Country:   line[7],
			CreatedAt: time.Now(),
		}
		wine.Priority, _ = strconv.Atoi(line[0])
		locations := strings.Split(line[4], ", ")
		wine.Locations = pq.Array(locations).(*pq.StringArray)
		wine.Price, _ = strconv.Atoi(line[5])
		wine.PriceType = mappingPrice(wine.Price)
		grapes := []string{line[8]}
		if line[9] != "" && line[8] != "0" {
			grapes = append(grapes, line[8])
		}
		if line[10] != "" && line[9] != "0" {
			grapes = append(grapes, line[9])
		}
		wine.Grapes = pq.Array(grapes).(*pq.StringArray)
		wine.Acidity, _ = strconv.Atoi(line[11])
		wine.Sweetness, _ = strconv.Atoi(line[12])
		if line[13] == "1" {
			wine.Sparkling = 1
		} else {
			wine.Sparkling = 0
		}
		foodMatches := strings.Split(line[14], ", ")
		wine.FoodMatches = pq.Array(foodMatches).(*pq.StringArray)
		wines = append(wines, wine)
	}
	return wines
}

func store(wines []Wine) (err error) {
	_, err = db.NamedExec(`INSERT INTO wines (priority, key_id, store, wine_name, locations, price, wine_type, country, grapes, acidity, sweetness, sparkling, food_matches, created_at)
		VALUES (:priority, :key_id, :store, :wine_name, :locations, :price, :price_type, :wine_type, :country, :grapes, :acidity, :sweetness, :sparkling, :food_matches, :created_at)`, wines[1:])
	if err != nil {
		log.Fatalln("ERROR during storing to Postgres DB:", err)
	}
	return
}

func (wine *Wine) ConvertedPrice() string {
	return fmt.Sprintf("₩%s", humanize.Comma(int64(wine.Price)))
}

// func PostByUUID(uuid string) (post Post, err error) {
// 	post = Post{}
// 	err = db.QueryRow("SELECT id, uuid, title, category, content, created_at FROM posts WHERE uuid = $1", uuid).
// 		Scan(&post.Id, &post.Uuid, &post.Title, pq.Array(&post.Category), &post.Content, &post.CreatedAt)
// 	return
// }
func WineById(id string) (wine Wine, err error) {
	wine = Wine{}
	err = db.Get(&wine, "SELECT * FROM person WHERE key_id=$1", id)
	if err != nil {
		log.Println("Error during WineById:", err)
	}
	return
}

func QueryWines(store, location, wineType, foodMatch string, price int) (wines []Wine, err error) {
	var statement string
	if store == "lottemart" {
		statement = "SELECT * FROM wines WHERE store = $1 AND price_type = $2 AND wine_type = $3 AND $4=any(food_matches) ORDER BY priority LIMIT 2"
		err = db.Select(&wines, statement, store, mappingPrice(price), wineType, foodMatch)
	} else {
		statement = "SELECT * FROM wines WHERE store = $1 AND price_type = $2 AND wine_type = $3 AND $4=any(food_matches) AND $5=any(locations) ORDER BY priority LIMIT 2"
		err = db.Select(&wines, statement, store, mappingPrice(price), wineType, foodMatch, location)
	}
	if err != nil {
		log.Println("Error during QueryWines, maybe no matching wines:", err)
	}
	return
}
