package data

import (
	"database/sql"
	"log"
)

var db *sql.DB

func init() {
	var err error
	DB_URL := "dbname=winecork sslmode=disable"
	// if v, ok := os.LookupEnv("DATABASE_URL"); ok {
	// 	DB_URL = v
	// }
	db, err = sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal("ERROR connecting postgres db:", err)
	}
}
