package main

import (
	"log"

	"github.com/go-gosh/tomato/app"
)

func main() {
	_, err := app.New("./config/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
}
