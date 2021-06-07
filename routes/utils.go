package routes

import (
	"log"
)

var logger *log.Logger

func init() {
	logger = log.Default()
	logger.SetFlags(log.Ldate)
	logger.SetFlags(log.Ltime)
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
