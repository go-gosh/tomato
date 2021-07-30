package main

import (
	"context"
	"time"

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
	userConfig, err := initApp(svc)
	if err != nil {
		panic(err)
	}

	a := interactivecli.NewMainView(db, svc, userConfig)
	err = a.Run()
	if err != nil {
		panic(err)
	}
}

func initApp(svc *service.Service) (*ent.UserConfig, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()
	cf := config.LoadDefaultConfig()
	user, err := svc.GetUserByUsername(ctx, cf.Runtime.DefaultUser)
	if ent.IsNotFound(err) {
		// user not found, create a default user
		create := service.UserCreate{
			Username: cf.Runtime.DefaultUser,
			Password: "",
			Config: struct {
				RedDuration   uint
				GreedDuration uint
			}{
				RedDuration:   25,
				GreedDuration: 5,
			},
		}
		user, err := svc.CreateUser(ctx, create)
		if err != nil {
			return nil, err
		}
		return user.QueryUserConfigs().First(ctx)
	}
	if err != nil {
		return nil, err
	}

	return user.QueryUserConfigs().First(ctx)
}
