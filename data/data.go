package data

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

var foodMatchingRed = map[string]string{
	"no-food":       "A01",
	"steak-beef":    "B01",
	"steak-port":    "B02",
	"pasta-meat":    "C01",
	"pasta-light":   "C02",
	"salad":         "D01",
	"western-oily":  "E01",
	"western-spicy": "E02",
	"western-sweet": "E03",
	"western-light": "E04",
	"dessert-salty": "F01",
	"dessert-sweet": "F02",
	"fish":          "G01",
}

var foodMatchingWhite = map[string]string{
	"no-food":       "A01",
	"fish-tuna":     "B01",
	"fish-sushi":    "B02",
	"fish-oyster":   "B03",
	"steak-beef":    "C01",
	"pasta-meat":    "D01",
	"pasta-light":   "D02",
	"salad":         "E01",
	"western-oily":  "F01",
	"western-spicy": "F02",
	"western-sweet": "F03",
	"western-light": "F04",
	"dessert-salty": "G01",
	"dessert-sweet": "G02",
}

func init() {
	var err error
	DB_URL := "dbname=winecork sslmode=disable"
	// if v, ok := os.LookupEnv("DATABASE_URL"); ok {
	// 	DB_URL = v
	// }

	db, err = sqlx.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal("ERROR connecting postgres db:", err)
	}
}
