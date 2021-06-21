package app

import (
	"context"
	"fmt"
	"os"

	"cauliflower/app/ent"
	"cauliflower/app/ent/migrate"
	"cauliflower/app/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

type App struct {
}

func New(path string) (*App, error) {
	cb, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	cf := Config{}
	err = yaml.NewDecoder(cb).Decode(&cf)
	if err != nil {
		return nil, err
	}

	db, err := ent.Open(cf.Database.Type, cf.Database.File)
	if err != nil {
		return nil, err
	}

	err = db.Schema.Create(context.Background(), migrate.WithForeignKeys(false))
	if err != nil {
		return nil, err
	}

	engine := gin.Default()
	svc := handler.New(db)
	svc.RegisterRoute(engine)

	err = engine.Run(fmt.Sprintf(":%d", cf.Application.Port))
	if err != nil {
		return nil, err
	}

	return &App{}, nil
}
