package service

import (
	"context"
	"time"

	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/task"
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

func (s Service) AddCheckpoints(ctx context.Context, taskId int, checkpoints ...*ent.CheckpointCreate) error {
	t, err := s.db.Task.Get(ctx, taskId)
	if err != nil {
		return err
	}
	max, _ := checkpoints[0].Mutation().CheckTime()
	for i := 0; i < len(checkpoints); i++ {
		checkpoints[i].SetTaskID(taskId)
		ct, ok := checkpoints[i].Mutation().CheckTime()
		if ok && ct.After(max) {
			max = ct
		}
	}

	if max.After(*t.Deadline) {
		_, err := s.db.Task.UpdateOneID(taskId).SetDeadline(max).Save(ctx)
		if err != nil {
			return err
		}
	}

	_, err = s.db.Checkpoint.CreateBulk(checkpoints...).Save(ctx)
	return err
}

// ListTask list task
func (s Service) ListTask(ctx context.Context) ([]*ent.Task, error) {
	return s.db.Task.Query().
		WithCheckpoints().
		All(ctx)
}

// GetTaskByDay get task with checkpoints by day.
func (s Service) GetTaskByDay(ctx context.Context, date time.Time) ([]*ent.Task, error) {
	y, m, d := date.Date()
	day := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	end := day.AddDate(0, 0, 1)
	return s.db.Task.Query().
		Where(task.Or(task.StartTimeLT(end), task.DeadlineGT(day))).
		All(ctx)
}
