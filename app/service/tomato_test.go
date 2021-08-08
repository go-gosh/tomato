package service

import (
	"context"
	"testing"
	"time"

	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/migrate"
	"github.com/go-gosh/tomato/app/ent/usertomato"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type _tomatoServiceTestSuite struct {
	suite.Suite
	svc    *Service
	userId int
}

func (s *_tomatoServiceTestSuite) SetupSuite() {
	db, err := ent.Open("sqlite3", ":memory:?_fk=1")
	s.Require().NoError(err)
	db = db.Debug()
	s.Require().NoError(db.Schema.Create(context.TODO(), migrate.WithForeignKeys(false)))
	s.svc = New(db)
}

func (s _tomatoServiceTestSuite) Test_NormalCase() {
	s.testUserNoWorkingOnTomato()
	s.testStartTomato()
	s.testUserHasWorkingOnTomato()
	s.testCloseTomato()
	s.testUserNoWorkingOnTomato()
}

func (s _tomatoServiceTestSuite) testUserNoWorkingOnTomato() {
	tomato, err := s.svc.GetWorkingTomatoByUserId(s.getContext(), s.userId)
	s.True(ent.IsNotFound(err))
	s.Nil(tomato)
}

func (s _tomatoServiceTestSuite) testUserHasWorkingOnTomato() {
	tomato, err := s.svc.GetWorkingTomatoByUserId(s.getContext(), s.userId)
	s.NoError(err)
	s.NotNil(tomato)
	s.NotEmpty(tomato.ID)
}

func (s _tomatoServiceTestSuite) testStartTomato() {
	param := TomatoCreate{
		Duration:  60,
		Color:     usertomato.ColorRed,
		UserId:    s.userId,
		StartTime: time.Now(),
	}
	tomato, err := s.svc.CreateTomato(s.getContext(), param)
	s.NoError(err)
	s.NotNil(tomato)
	s.NotEmpty(tomato.ID)
	s.EqualValues(param.Color, tomato.Color)
	s.EqualValues(param.UserId, tomato.UserID)
	s.EqualValues(param.StartTime.Unix(), tomato.StartTime.Unix())
	s.EqualValues(param.Duration, tomato.RemainTime.Sub(tomato.StartTime)/time.Second)
	s.Nil(tomato.EndTime)
}

func (s _tomatoServiceTestSuite) testCloseTomato() {
	err := s.svc.CloseTomatoByUserId(s.getContext(), s.userId)
	s.NoError(err)
}

func (s *_tomatoServiceTestSuite) getContext() context.Context {
	return context.TODO()
}

func TestTomatoService(t *testing.T) {
	suite.Run(t, &_tomatoServiceTestSuite{})
}
