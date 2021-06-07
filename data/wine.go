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

func (wine *Wine) GetWineInfo() string {
	return fmt.Sprintf("%s | %s | %s", wine.Country, wine.stripGrapes(), wine.getWineType())
}

func (wine *Wine) stripGrapes() string {
	s := fmt.Sprint(wine.Grapes)
	return s[2 : len(s)-1]
}

func (wine *Wine) getWineType() string {
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

func QueryWines(store, wineType, foodMatch, price string) (wines []Wine, err error) {
	var statement string
	if wineType == "red" {
		foodMatch = foodMatchingRed[foodMatch]
	} else {
		foodMatch = foodMatchingWhite[foodMatch]
	}
	storeLocation := strings.Split(store, " ")
	var location string
	if len(storeLocation) > 1 {
		store, location = storeLocation[0], storeLocation[1]
	}
	price = string(price[len(price)-1])
	fmt.Println(store, location, wineType, foodMatch, price)
	if store == "롯데마트" {
		statement = "SELECT * FROM wines WHERE store = $1 AND price_type = $2 AND wine_type = $3 AND $4=any(food_matches) ORDER BY priority LIMIT 2"
		err = db.Select(&wines, statement, store, price, wineType, foodMatch)
	} else {
		statement = "SELECT * FROM wines WHERE store = $1 AND price_type = $2 AND wine_type = $3 AND $4=any(food_matches) AND $5=any(locations) ORDER BY priority LIMIT 2"
		err = db.Select(&wines, statement, store, price, wineType, foodMatch, location)
	}
	if err != nil || len(wines) == 0 {
		warning("Error during QueryWines, maybe no matching wines:", err)
	}
	return
}
