package data

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type StoreInfo struct {
	StoreType string
	Location  string
	Latitude  string
	Longitude string
}

var db *sqlx.DB

var foodMatchingCode = map[string]string{
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
	"fish-tuna":     "F01",
	"fish-sushi":    "F02",
	"fish-oyster":   "F03",
	"dessert-salty": "G01",
	"dessert-sweet": "G02",
}

var foodMatchingRed = map[string]string{
	"F01": "G01",
	"F02": "G02",
	"G01": "",
}

var foodMatchingWhite = map[string]string{
	"B01": "F01",
	"B02": "F02",
	"B03": "F03",
	"C01": "",
	"D01": "C01",
	"D02": "C02",
	"E01": "D01",
	"F01": "E01",
	"F02": "E02",
	"F03": "E03",
	"F04": "E04",
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

var stores = []StoreInfo{}

func init() {
	initLogger()
	info("Initiating DB")
	var err error
	// DB_URL := "dbname=winecork2 sslmode=disable"
	hostname := os.Getenv("RDS_HOSTNAME")
	dbname := os.Getenv("RDS_DB_NAME")
	username := os.Getenv("RDS_USERNAME")
	password := os.Getenv("RDS_PASSWORD")
	port := os.Getenv("RDS_PORT")
	DB_URL := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s", hostname, dbname, username, password, port)
	log.Println("DB_URL:", DB_URL)
	if v, ok := os.LookupEnv("DATABASE_URL"); ok {
		DB_URL = v
	}

	db, err = sqlx.Open("postgres", DB_URL)
	if err != nil {
		danger("ERROR connecting postgres db:", err)
	}

	parseStoreInfo()
	info("DB is up and running.")
}

func clearWineDB() {
	info("Clearing Wine DB...")
	schema := `
	drop table if exists wines;
	
	create table wines (
	  id           serial primary key,
	  priority     int,
	  key_id       varchar(255) not null unique,
	  store        varchar(255),
	  wine_name    varchar(255),
	  locations    varchar(255)[],
	  price        int,
	  price_type   varchar(64),
	  wine_type    varchar(64),
	  country      varchar(255),
	  grapes       varchar(255)[],
	  acidity      int,
	  sweetness    int,
	  sparkling    int,
	  food_matches varchar(255)[],
	  image_url    text,
	  created_at   timestamp not null  
	);
	`

	db.MustExec(schema)
	info("DB Cleared.")
}

func GetStoreLocations() ([]byte, error) {
	return json.Marshal(stores)
}

func SaveCSV(csvFile io.Reader) (err error) {
	urls := getImageUrl()
	wines := parseCSV(csvFile, urls)
	clearWineDB()
	err = insertWines(wines)
	if err != nil {
		danger("ERROR during insertWines:", err)
	}
	return
}

func parseStoreInfo() (err error) {
	info("Parsing store_coord.csv...")
	csvFile, err := os.Open(path.Join("data", "store_coord.csv"))
	if err != nil {
		danger("Error opening store_coord.csv: ", err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		var line []string
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			warning("error during parsing csv: ", err)
			continue
		}
		location := strings.ReplaceAll(line[1], " ", "")
		storeInfo := StoreInfo{
			StoreType: line[0],
			Location:  location,
			Latitude:  line[2],
			Longitude: line[3],
		}
		stores = append(stores, storeInfo)
	}
	info("Store info parsed.")
	return
}

func parseCSV(csvFile io.Reader, urls map[string]string) (wines []Wine) {
	info("Parsing CSV file...")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	isHeader := true
	for {
		var line []string
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			danger("error duing parsing csv:", err)
		} else if isHeader {
			isHeader = false
			continue
		}
		store := line[2]
		if store == "이마트몰" {
			store = "이마트"
		}
		wine := Wine{
			Key:       line[1],
			Store:     store,
			WineName:  line[3],
			WineType:  line[6],
			Country:   line[7],
			CreatedAt: time.Now(),
			Image:     urls[line[1]],
		}
		wine.Priority, _ = strconv.Atoi(line[0])
		locations := strings.Split(line[4], ", ")
		if len(locations) == 0 && store != "이마트" {
			continue
		}
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
		foodMatches = convertFoodMatch(wine.WineType, foodMatches)
		wine.FoodMatches = pq.Array(foodMatches).(*pq.StringArray)
		wines = append(wines, wine)
	}
	info("Parsing CSV finished.")
	return
}

func convertFoodMatch(wineType string, code []string) []string {
	converted := []string{}
	convert := func(mapping map[string]string) {
		for _, v := range code {
			if x, ok := mapping[v]; ok {
				if x != "" {
					converted = append(converted, x)
				}
			} else {
				converted = append(converted, v)
			}
		}
	}
	if wineType == "red" {
		convert(foodMatchingRed)
	} else if wineType == "white" {
		convert(foodMatchingWhite)
	} else {
		danger("Not valid wine type: ", wineType)
	}

	return converted
}

func insertWines(wines []Wine) (err error) {
	info("Storing Wines to DB...")
	_, err = db.NamedExec(`INSERT INTO wines (priority, key_id, store, wine_name, locations, price, price_type, wine_type, country, grapes, acidity, sweetness, sparkling, food_matches, image_url, created_at)
		VALUES (:priority, :key_id, :store, :wine_name, :locations, :price, :price_type, :wine_type, :country, :grapes, :acidity, :sweetness, :sparkling, :food_matches, :image_url, :created_at)`, wines)
	if err != nil {
		danger("ERROR during storing to Postgres DB:", err)
	} else {
		info("Wines stored to DB.")
	}
	return
}
