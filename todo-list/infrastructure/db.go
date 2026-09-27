package infrastructure

import (
	"context"
	"todo/domain"
)

type ProjectDB struct{}

func NewProjectDB() *ProjectDB {
	return &ProjectDB{}
}

func (db *ProjectDB) FindById(ctx context.Context, id domain.ProjectID) (*domain.Project, error) {
	return &domain.Project{}, nil
}
func (db *ProjectDB) Save(ctx context.Context, project *domain.Project) error {
	return nil
}

type TaskDB struct{}

func NewTaskDB() *TaskDB {
	return &TaskDB{}
}

func (db *TaskDB) FindById(ctx context.Context, id domain.TaskID) (*domain.Task, error) {
	return &domain.Task{}, nil
}
func (db *TaskDB) Save(ctx context.Context, task *domain.Task) error {
	return nil
}
