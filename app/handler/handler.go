package handler

import (
	"net/http"
	"time"

	"github.com/go-gosh/tomato/app/context"
	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/usertomato"
	"github.com/go-gosh/tomato/app/service"

	"github.com/gin-gonic/gin"
)

type Service struct {
	svc *service.Service
}

func New(svc *service.Service) *Service {
	return &Service{svc: svc}
}

func (s *Service) RegisterRoute(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	v1.GET("/working-tomato", ginAdapter(s.GetWorkingOnTomato))
	v1.POST("/tomato", ginAdapter(s.StartTomato))
	v1.POST("/closing-tomato", ginAdapter(s.CloseTomato))
}

// GetWorkingOnTomato get a working on tomato clock for login user
func (s Service) GetWorkingOnTomato(ctx *context.Context) error {
	us, err := ctx.LoginUser()
	if err != nil {
		return err
	}
	tomato, err := s.svc.GetWorkingTomatoByUserId(ctx, us.ID)
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
		// Color signed tomato clock type
		Color string `json:"color" binding:"required,oneof=red green"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return err
	}

	us, err := ctx.LoginUser()
	if err != nil {
		return err
	}
	param := service.TomatoCreate{
		Duration:  req.Duration,
		Color:     usertomato.Color(req.Color),
		UserId:    us.ID,
		StartTime: time.Now(),
	}
	t, err := s.svc.CreateTomato(ctx, param)
	if err != nil {
		return err
	}
	ctx.Response(http.StatusOK, t, "")

	return nil
}

func (s Service) CloseTomato(ctx *context.Context) (err error) {
	us, err := ctx.LoginUser()
	if err != nil {
		return err
	}
	err = s.svc.CloseTomatoByUserId(ctx, us.ID)
	if err != nil {
		return
	}
	ctx.Response(http.StatusOK, gin.H{}, "success")

	return
}
