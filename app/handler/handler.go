package handler

import (
	"net/http"

	"github.com/go-gosh/tomato/app/context"
	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/usertomato"

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
	v1.GET("/working-tomato", ginAdapter(s.GetWorkingOnTomato))
	v1.POST("/tomato", ginAdapter(s.StartTomato))
}

// GetWorkingOnTomato get a working on tomato clock for login user
func (s Service) GetWorkingOnTomato(ctx *context.Context) error {
	us, err := ctx.LoginUser()
	if err != nil {
		return err
	}
	tomato, err := s.db.UserTomato.Query().
		Where(usertomato.And(
			usertomato.UserIDEQ(us.ID),
			usertomato.EndTimeIsNil(),
		)).
		Order(ent.Desc(usertomato.FieldID)).
		First(ctx)
	if ent.IsNotFound(err) {
		ctx.Response(http.StatusOK, nil, "not found tomato")
		return nil
	}
	if err != nil {
		return err
	}
	ctx.Response(http.StatusOK, tomato, "")

	return nil
}

// StartTomato TODO implements me
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
