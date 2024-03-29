package app

import (
	"context"
	"fmt"

	"github.com/go-gosh/tomato/app/config"
	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/migrate"
	"github.com/go-gosh/tomato/app/handler"
	"github.com/go-gosh/tomato/app/service"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
}

func New() (*App, error) {
	cf := config.LoadDefaultConfig()

	db, err := ent.Open(cf.Database.Type, cf.Database.File)
	if err != nil {
		return nil, err
	}

	err = db.Schema.Create(context.Background(), migrate.WithForeignKeys(false))
	if err != nil {
		return nil, err
	}

	engine := gin.Default()
	svc := service.New(db)
	hd := handler.New(svc)
	hd.RegisterRoute(engine)

	err = engine.Run(fmt.Sprintf(":%d", cf.Application.Port))
	if err != nil {
		return nil, err
	}

	return &App{}, nil
}
