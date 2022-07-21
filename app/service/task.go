package service

import (
	"context"
	"time"

	"github.com/go-gosh/tomato/app/ent"
)

type TaskCreate struct {
	Title    string
	Category string
	Star     int8
	Content  string
	JoinTime time.Time
	Deadline *time.Time
	StartNow bool
}

// CreateTask create a task
func (s Service) CreateTask(ctx context.Context, task TaskCreate) (*ent.Task, error) {
	taskCreate := s.db.Task.Create().
		SetTitle(task.Title).
		SetCategory(task.Category).
		SetStar(task.Star).
		SetContent(task.Content).
		SetJoinTime(task.JoinTime).
		SetNillableDeadline(task.Deadline)
	if task.StartNow {
		taskCreate.SetStartTime(time.Now())
	}
	t, err := taskCreate.
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}
