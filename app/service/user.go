package service

import (
	"context"

	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/user"
)

type UserCreate struct {
	Username string
	Password string
	Config   struct {
		RedDuration   uint
		GreedDuration uint
	}
}

// CreateUser init user config and return it.
func (s Service) CreateUser(ctx context.Context, create UserCreate) (*ent.User, error) {
	tx, err := s.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	us, err := s.db.User.Create().
		SetUsername(create.Username).
		SetPassword(create.Password).
		SetEnabled(true).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	_, err = tx.UserConfig.Create().
		SetWorking(create.Config.RedDuration).
		SetBreak(create.Config.GreedDuration).
		SetRank(0).
		SetUsers(us).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return us, nil
}

// GetUserByUsername returns user entity by username if exists.
// If user is not exists, it will return errors of not found
func (s Service) GetUserByUsername(ctx context.Context, username string) (*ent.User, error) {
	return s.db.User.Query().Where(
		user.UsernameEQ(username),
	).Only(ctx)
}
