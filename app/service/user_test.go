package service

import (
	"context"
	"testing"

	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type _userServiceTestSuite struct {
	suite.Suite
	svc *Service
}

func (s *_userServiceTestSuite) SetupSuite() {
	db, err := ent.Open("sqlite3", ":memory:?_fk=1")
	s.Require().NoError(err)
	db = db.Debug()
	s.Require().NoError(db.Schema.Create(context.TODO(), migrate.WithForeignKeys(false)))
	s.svc = New(db)
}

func (s _userServiceTestSuite) TestCreateUser() {
	type args struct {
		req UserCreate
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"insert one user success", args{UserCreate{"user1", "", UserConfigCreate{1500, 300}}}, false},
		{"insert duplicate user failed", args{UserCreate{"user1", "", UserConfigCreate{1500, 300}}}, true},
		{"insert another user success", args{UserCreate{"user2", "", UserConfigCreate{1000, 3000}}}, false},
	}
	ctx := context.TODO()
	for _, tt := range tests {
		s.Run(tt.name, func() {
			actual, err := s.svc.CreateUser(ctx, tt.args.req)
			if tt.wantErr {
				s.Error(err)
				s.Nil(actual)
				return
			}
			s.NoError(err)
			s.NotNil(actual)
			s.NotEmpty(actual.ID)
			s.True(actual.Enabled)
			s.EqualValues(tt.args.req.Username, actual.Username)
			s.EqualValues(tt.args.req.Password, actual.Password)
			// assert user config
			s.Len(actual.Edges.UserConfigs, 1)
			c := actual.Edges.UserConfigs[0]
			s.EqualValues(tt.args.req.Config.RedDuration, c.Working)
			s.EqualValues(tt.args.req.Config.GreedDuration, c.Break)
		})
	}
}

func TestUserService(t *testing.T) {
	suite.Run(t, &_userServiceTestSuite{})
}
