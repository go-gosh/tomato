package main

import (
	"log"

	"cauliflower/app"
)

func main() {
	_, err := app.New("./config/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
}
