package main

import (
	"context"

	"github.com/go-gosh/tomato/app/config"
	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/migrate"
	"github.com/go-gosh/tomato/app/interactivecli"
	"github.com/go-gosh/tomato/app/service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cf := config.LoadDefaultConfig()

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
