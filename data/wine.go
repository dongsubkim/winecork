package data

import (
	"fmt"
	"strconv"

	"github.com/dustin/go-humanize"
)

type Wine struct {
	Id         int
	Priority   int
	Key        string
	Store      string
	WineName   string
	Locations  []string
	Price      int
	WineType   string
	Country    string
	Grape      string
	FoodMatchs []string
	Image      string
}

func (wine *Wine) ConvertedPrice() string {
	return fmt.Sprintf("₩%s", humanize.Comma(int64(wine.Price)))
}

func WineById(id string) (wine Wine, err error) {
	wineId, _ := strconv.Atoi(id)
	wine = Wine{
		Id:       wineId,
		WineName: "풋 프린트 쉬라즈 750mL",
		Price:    22000,
		Image:    "static/icons/EM0001.jpg",
		Country:  "프랑스",
		WineType: "레드와인",
		Grape:    `그르나슈`,
	}
	return
}
