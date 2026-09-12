package domain

import (
	"context"
)

type ProjectRepository interface {
	FindById(context.Context, ProjectID) (*Project, error)
	Save(context.Context, *Project) error
}

type TaskRepository interface {
	FindById(context.Context, TaskID) (*Task, error)
	Save(context.Context, *Task) error
}
