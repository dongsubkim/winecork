package data

import (
	"bufio"
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var db *sqlx.DB

var foodMatchingRed = map[string]string{
	"no-food":       "A01",
	"steak-beef":    "B01",
	"steak-pork":    "B02",
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
	"steak":         "C01",
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

var foodMatchingKorean = map[string]string{
	"no-food":       "그냥 사둘게요",
	"steak-beef":    "고기/스테이크 | 소/양고기",
	"steak-pork":    "고기/스테이크 | 돼지/오리고기",
	"pasta-meat":    "파스타 | 미트/오일/크림",
	"pasta-light":   "파스타 | 담백",
	"salad":         "샐러드",
	"western-oily":  "양식/한식요리 | 기름진 요리",
	"western-spicy": "양식/한식요리 | 매운양념요리",
	"western-sweet": "양식/한식요리 | 달짝지근한 요리",
	"western-light": "양식/한식요리 | 담백한 요리",
	"dessert-salty": "핑거푸드/디저트 | 짭쪼름한 하몽/치즈/크래커",
	"dessert-sweet": "핑거푸드/디저트 | 달달한 디저트",
	"fish":          "해산물",
	"fish-tuna":     "해산물 | 참치/연어",
	"fish-sushi":    "해산물 | 스시",
	"fish-oyster":   "해산물 | 굴",
}

var priceRangeRaw = map[string]string{
	"price0": "￦1만↓",
	"price1": "￦1만",
	"price2": "￦2만",
	"price3": "￦3만~￦4만",
	"price4": "₩4만↑",
}

func init() {
	initLogger()
	info("Initiating DB")
	var err error
	DB_URL := "dbname=winecork sslmode=disable"
	if v, ok := os.LookupEnv("DATABASE_URL"); ok {
		DB_URL = v
	}

	db, err = sqlx.Open("postgres", DB_URL)
	if err != nil {
		danger("ERROR connecting postgres db:", err)
	}
	// csvFile, err := os.Open(path.Join("data", "wine_db.csv"))
	// if err != nil {
	// 	danger("Error opening wine_db.csv: ", err)
	// }
	// defer csvFile.Close()
	// urls := getImageUrl()
	// wines := parseCSV(csvFile, urls)
	// clearDB()
	// err = store(wines)
	// if err != nil {
	// 	danger("Error during storing csv", err)
	// }
	info("DB is up and running.")
}

func clearDB() {
	info("Clearing DB...")
	path := filepath.Join("data", "setup.sql")

	c, err := ioutil.ReadFile(path)
	if err != nil {
		danger("Error during reading setup.sql file:", err)
	}
	schema := string(c)

	db.MustExec(schema)
	info("DB Cleared.")
}

func parseCSV(csvFile io.Reader, urls map[string]string) (wines []Wine) {
	info("Parsing CSV file...")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		var line []string
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			danger("error duing parsing csv:", err)
		}
		wine := Wine{
			Key:       line[1],
			Store:     line[2],
			WineName:  line[3],
			WineType:  line[6],
			Country:   line[7],
			CreatedAt: time.Now(),
			Image:     urls[line[1]],
		}
		wine.Priority, _ = strconv.Atoi(line[0])
		locations := strings.Split(line[4], ", ")
		wine.Locations = pq.Array(locations).(*pq.StringArray)
		wine.Price, _ = strconv.Atoi(line[5])
		wine.PriceType = mappingPrice(wine.Price)
		if line[6] == "레드" {
			wine.WineType = "red"
		} else {
			wine.WineType = "white"
		}
		grapes := []string{line[8]}
		if line[9] != "" && line[9] != "0" && line[8] != "0" {
			grapes = append(grapes, line[9])
		}
		if line[10] != "" && line[10] != "0" && line[9] != "0" {
			grapes = append(grapes, line[10])
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
	info("Parsing CSV finished.")
	return
}

func store(wines []Wine) (err error) {
	info("Storing Wines to DB...")
	_, err = db.NamedExec(`INSERT INTO wines (priority, key_id, store, wine_name, locations, price, price_type, wine_type, country, grapes, acidity, sweetness, sparkling, food_matches, image_url, created_at)
		VALUES (:priority, :key_id, :store, :wine_name, :locations, :price, :price_type, :wine_type, :country, :grapes, :acidity, :sweetness, :sparkling, :food_matches, :image_url, :created_at)`, wines[1:])
	if err != nil {
		danger("ERROR during storing to Postgres DB:", err)
	} else {
		info("Wines stored to DB.")
	}
	return
}
