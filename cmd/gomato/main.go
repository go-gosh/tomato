package main

import (
	"context"
	"os"

	"github.com/go-gosh/tomato/app"
	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/migrate"
	"github.com/go-gosh/tomato/app/interactivecli"
	"github.com/go-gosh/tomato/app/service"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

func main() {
	cb, err := os.Open("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	cf := app.Config{}
	err = yaml.NewDecoder(cb).Decode(&cf)
	if err != nil {
		panic(err)
	}

	db, err := ent.Open(cf.Database.Type, cf.Database.File)
	if err != nil {
		panic(err)
	}

	err = db.Schema.Create(context.Background(), migrate.WithForeignKeys(false))
	if err != nil {
		panic(err)
	}

	svc := service.New(db)
	a := interactivecli.NewMainView(db, svc)
	err = a.Run()
	if err != nil {
		panic(err)
	}
}
