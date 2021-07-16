package main

import (
	"github.com/go-gosh/tomato/app/interactivecli"
)

func main() {
	a := &interactivecli.MainView{}
	err := a.Run()
	if err != nil {
		panic(err)
	}
}
