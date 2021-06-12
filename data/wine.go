package data

import (
	"fmt"
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
	PriceType   string `db:"price_type"`
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

type QueryLog struct {
	Id        int
	Store     string
	Price     string
	FoodMatch string    `db:"food_match"`
	CreatedAt time.Time `db:"created_at"`
}

func mappingPrice(price int) string {
	if price < 10000 {
		return "0"
	} else if price < 20000 {
		return "1"
	} else if price < 30000 {
		return "2"
	} else if price < 40000 {
		return "3"
	} else {
		return "4"
	}
}

func (wine *Wine) ConvertedPrice() string {
	return fmt.Sprintf("₩%s", humanize.Comma(int64(wine.Price)))
}

func (q *QueryLog) Date() string {
	return q.CreatedAt.Format("06/01/02 3:04pm")
}

func (wine *Wine) GetWineInfo() string {
	return fmt.Sprintf("%s | %s | %s", wine.Country, wine.StripGrapes(), wine.GetWineType())
}

func (wine *Wine) StripGrapes() string {
	s := fmt.Sprint(wine.Grapes)
	return s[2 : len(s)-1]
}

func (wine *Wine) GetImage(h int) string {
	return strings.Replace(wine.Image, "/upload/", fmt.Sprintf("/upload/c_fill,w_130,h_%d/", h), 1)
}

func (wine *Wine) ListGrapes() string {
	return listPqArray(wine.Grapes)
}
func (wine *Wine) ListLocations() string {
	return listPqArray(wine.Locations)
}
func (wine *Wine) ListFoodMatches() string {
	return listPqArray(wine.FoodMatches)
}
func listPqArray(sa *pq.StringArray) string {
	s := fmt.Sprint(sa)
	return strings.Join(strings.Split(s[2:len(s)-1], " "), ", ")
}

func (wine *Wine) GetWineType() string {
	if wine.WineType == "red" {
		return "레드 와인"
	} else {
		return "화이트 와인"
	}
}

func WineById(id string) (wine Wine, err error) {
	wine = Wine{}
	err = db.Get(&wine, "SELECT * FROM wines WHERE key_id=$1", id)
	if err != nil {
		warning("Error during WineById:", err)
	}
	return
}

func QueryWines(store, foodMatch, price string) (wines []Wine, err error) {
	logQuery(store, foodMatch, price)
	var statement string
	var location string
	var wineType string

	if foodMatch == "steak" {
		wineType = "red"
	} else if foodMatch == "fish" {
		wineType = "white"
	}
	foodMatch = foodMatchingCode[foodMatch]

	storeLocation := strings.Split(store, " ")
	if len(storeLocation) > 1 {
		store, location = storeLocation[0], storeLocation[1]
	}

	price = string(price[len(price)-1])

	if store == "롯데마트" {
		statement = "SELECT * FROM wines WHERE store = $1 AND price_type = $2 AND $3=any(food_matches) ORDER BY priority LIMIT 2"
		if len(wineType) > 0 {
			statement = strings.Replace(statement, "ORDER BY", "AND wine_type = $4 ORDER BY", 1)
			err = db.Select(&wines, statement, store, price, foodMatch, wineType)
		} else {
			err = db.Select(&wines, statement, store, price, foodMatch)
		}
	} else {
		statement = "SELECT * FROM wines WHERE store = $1 AND price_type = $2 AND $3=any(food_matches) AND $4=any(locations) ORDER BY priority LIMIT 2"
		if len(wineType) > 0 {
			statement = strings.Replace(statement, "ORDER BY", "AND wine_type = $5 ORDER BY", 1)
			err = db.Select(&wines, statement, store, price, foodMatch, location, wineType)
		} else {
			err = db.Select(&wines, statement, store, price, foodMatch, location)
		}
	}
	if err != nil || len(wines) == 0 {
		warning("Error during QueryWines, maybe no matching wines:", err)
	}
	return
}

func GetAllWines() (wines []Wine) {
	db.Select(&wines, "SELECT * FROM wines ORDER BY priority")
	return
}

func logQuery(store, foodMatch, price string) {
	query := QueryLog{
		Store:     store,
		Price:     priceRangeRaw[price],
		FoodMatch: foodMatchingKorean[foodMatch],
		CreatedAt: time.Now(),
	}
	_, err := db.NamedExec(`INSERT INTO querylogs (store, price, food_match, created_at) VALUES (:store, :price, :food_match, :created_at)`, query)
	if err != nil {
		warning("Error during logQuery:", err)
	}
}

func GetAllQueries() (quries []QueryLog) {
	db.Select(&quries, "SELECT * FROM querylogs ORDER BY created_at DESC")
	return
}
