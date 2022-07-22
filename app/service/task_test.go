package service

import "time"

func (s _tomatoServiceTestSuite) Test_Tasks() {
	param := TaskCreate{
		Title:    "test title",
		Category: "category",
		Star:     0,
		Content:  "test content",
		JoinTime: time.Now(),
		StartNow: true,
	}
	task, err := s.svc.CreateTask(s.getContext(), param)
	s.NoError(err)
	s.NotNil(task)
	s.EqualValues(param.Title, task.Title)
	s.EqualValues(param.Category, task.Category)
	s.EqualValues(param.Star, task.Star)
	s.EqualValues(param.Content, task.Content)
	s.EqualValues(param.JoinTime.Unix(), task.JoinTime.Unix())
	s.NotNil(task.JoinTime)
	s.Nil(task.Deadline)
}

func (s _tomatoServiceTestSuite) Test_ListTask() {
	s.Run("setup_test", s.Test_Tasks)
	task, err := s.svc.ListTask(s.getContext())
	s.NoError(err)
	s.Len(task, 1)
	s.T().Logf("%+v", task)
}
