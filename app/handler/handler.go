package handler

import (
	"errors"
	"net/http"
	"time"

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

// StartTomato start a tomato clock if there is no working tomato clock.
func (s Service) StartTomato(ctx *context.Context) error {
	req := struct {
		// Duration unit is second
		Duration int `json:"duration" binding:"required,min=1"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return err
	}

	us, err := ctx.LoginUser()
	if err != nil {
		return err
	}

	tx, err := s.db.Tx(ctx)
	if err != nil {
		return err
	}

	c, err := tx.UserTomato.Query().Where(
		usertomato.UserID(us.ID),
		usertomato.EndTimeIsNil(),
	).Count(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	if c > 0 {
		tx.Rollback()
		return errors.New("one tomato clock is working")
	}

	timestamp := time.Now()
	t, err := tx.UserTomato.Create().
		SetUsersID(us.ID).
		SetStartTime(timestamp).
		SetRemainTime(timestamp.Add(time.Duration(req.Duration) * time.Second)).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	ctx.Response(http.StatusOK, t, "")

	return nil
}
