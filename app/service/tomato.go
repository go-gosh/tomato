package service

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/usertomato"
)

var (
	_once sync.Once
	_inst *Service
)

type Service struct {
	db *ent.Client
}

func New(db *ent.Client) *Service {
	_once.Do(func() {
		_inst = &Service{db: db}
	})

	return _inst
}

// GetWorkingTomatoByUserId get a working on tomato clock by user id
func (s Service) GetWorkingTomatoByUserId(ctx context.Context, userId int) (*ent.UserTomato, error) {
	return s.db.UserTomato.Query().
		Where(usertomato.And(
			usertomato.UserIDEQ(userId),
			usertomato.EndTimeIsNil(),
		)).
		Order(ent.Desc(usertomato.FieldID)).
		First(ctx)
}

// TomatoCreate create new tomato clock param
type TomatoCreate struct {
	// Duration unit is second
	Duration int
	// Color signed tomato clock type
	Color  usertomato.Color
	UserId int
	// StartTime clock start time
	StartTime time.Time
}

// CreateTomato create a tomato clock if there is no working on tomato clock
func (s Service) CreateTomato(ctx context.Context, param TomatoCreate) (*ent.UserTomato, error) {
	c, err := s.db.UserTomato.Query().
		Where(
			usertomato.UserID(param.UserId),
			usertomato.EndTimeIsNil(),
		).Count(ctx)
	if err != nil {
		return nil, err
	}
	if c > 0 {
		return nil, errors.New("one tomato clock is working")
	}
	t, err := s.db.UserTomato.Create().
		SetUsersID(param.UserId).
		SetStartTime(param.StartTime).
		SetColor(param.Color).
		SetRemainTime(param.StartTime.Add(time.Duration(param.Duration) * time.Second)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// CloseTomatoByUserId close all tomato clock by user id.
func (s Service) CloseTomatoByUserId(ctx context.Context, userId int) error {
	return s.db.UserTomato.Update().
		Where(usertomato.And(
			usertomato.UserIDEQ(userId),
			usertomato.EndTimeIsNil(),
		)).SetEndTime(time.Now()).
		Exec(ctx)
}

// GiveUpTomatoByUserId give up all tomato clock by user id.
func (s Service) GiveUpTomatoByUserId(ctx context.Context, userId int) error {
	_, err := s.db.UserTomato.Delete().
		Where(
			usertomato.UserIDEQ(userId),
			usertomato.EndTimeIsNil(),
		).Exec(ctx)
	return err
}
