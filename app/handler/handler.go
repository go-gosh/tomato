package handler

import (
	"net/http"

	"github.com/go-gosh/tomato/app/context"
	"github.com/go-gosh/tomato/app/ent"

	"github.com/gin-gonic/gin"
)

type Service struct {
	db *ent.Client
}

func New(db *ent.Client) *Service {
	return &Service{db: db}
}

func (s *Service) RegisterRoute(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	v1.POST("/tomato", ginAdapter(s.StartTomato))
}

func (s Service) StartTomato(ctx *context.Context) error {
	user, err := ctx.LoginUser()
	if err != nil {
		return err
	}
	t, err := s.db.UserTomato.Create().
		SetUsersID(user.ID).
		Save(ctx)
	if err != nil {
		return err
	}
	ctx.Response(http.StatusOK, t, "")

	return nil
}
