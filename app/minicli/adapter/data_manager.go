package adapter

import (
	"context"
	"time"

	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/usertomato"
	"github.com/go-gosh/tomato/app/minicli"
	"github.com/go-gosh/tomato/app/service"
)

func NewDataManager(svc *service.Service, config *ent.UserConfig, ctx context.Context) minicli.DataManager {
	return &dataManager{svc: svc, config: config, ctx: ctx}
}

type dataManager struct {
	svc    *service.Service
	config *ent.UserConfig
	ctx    context.Context
}

func (d dataManager) Get() *minicli.Data {
	t, err := d.svc.GetWorkingTomatoByUserId(d.ctx, d.config.UserID)
	if ent.IsNotFound(err) {
		return nil
	}
	return d.mustToData(t, err)
}

func (d dataManager) Create() *minicli.Data {
	create := service.TomatoCreate{
		Duration:  int(d.config.Working),
		Color:     usertomato.ColorRed,
		UserId:    d.config.UserID,
		StartTime: time.Now(),
	}
	t, err := d.svc.CreateTomato(d.ctx, create)
	return d.mustToData(t, err)
}

func (d dataManager) GiveUp() *minicli.Data {
	return d.mustToData(nil, d.svc.GiveUpTomatoByUserId(d.ctx, d.config.UserID))
}

func (d dataManager) Submit(color string) *minicli.Data {
	err := d.svc.CloseTomatoByUserId(d.ctx, d.config.UserID)
	if err != nil {
		panic(err)
	}

	var create service.TomatoCreate
	if color == string(usertomato.ColorRed) {
		create = service.TomatoCreate{
			Duration:  int(d.config.Break),
			Color:     usertomato.ColorGreen,
			UserId:    d.config.UserID,
			StartTime: time.Now(),
		}
	} else {
		create = service.TomatoCreate{
			Duration:  int(d.config.Working),
			Color:     usertomato.ColorRed,
			UserId:    d.config.UserID,
			StartTime: time.Now(),
		}
	}
	return d.mustToData(d.svc.CreateTomato(d.ctx, create))
}

func (d dataManager) Close() *minicli.Data {
	return d.mustToData(nil, d.svc.CloseTomatoByUserId(d.ctx, d.config.UserID))
}

func (d dataManager) mustToData(t *ent.UserTomato, err error) *minicli.Data {
	if err != nil {
		panic(err)
	}
	if t == nil {
		return nil
	}

	return &minicli.Data{
		EndTime:  t.RemainTime,
		Duration: int(t.RemainTime.Sub(t.StartTime) / time.Second),
		Color:    string(t.Color),
	}
}
