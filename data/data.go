package data

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

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
