package main

import (
	"context"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-gosh/tomato"
	"github.com/go-gosh/tomato/app/config"
	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/migrate"
	"github.com/go-gosh/tomato/app/minicli"
	"github.com/go-gosh/tomato/app/minicli/adapter"
	"github.com/go-gosh/tomato/app/service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	args := os.Args
	if len(args) > 1 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Println(tomato.VersionTag)
		return
	}

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
	a := minicli.NewModel(adapter.NewDataManager(svc, userConfig, context.Background()))
	err = tea.NewProgram(a).Start()
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
			Config: service.UserConfigCreate{
				RedDuration:   25 * 60,
				GreedDuration: 5 * 60,
			},
		}
		user, err := svc.CreateUser(ctx, create)
		if err != nil {
			return nil, err
		}
		return user.Edges.UserConfigs[0], nil
	}
	if err != nil {
		return nil, err
	}

	return user.Edges.UserConfigs[0], nil
}
