package app

import (
	"context"
	"fmt"
	"todo/domain"
)

type ProjectApp struct {
	projectRepo     domain.ProjectRepository
	taskRepo        domain.TaskRepository
	eventDispatcher domain.EventDispatcher
}

func NewProjectApp(projectRepo domain.ProjectRepository, taskRepo domain.TaskRepository, dispatcher domain.EventDispatcher) *ProjectApp {
	return &ProjectApp{
		projectRepo:     projectRepo,
		taskRepo:        taskRepo,
		eventDispatcher: dispatcher,
	}
}

func (p *ProjectApp) CompleteTask(ctx context.Context, taskID string) error {
	domainTaskID := domain.NewTaskID(taskID)
	task, err := p.taskRepo.FindById(ctx, domainTaskID)
	if err != nil {
		return fmt.Errorf("failed to find task: %w", err)
	}
	task.Complete()

	if err := p.taskRepo.Save(ctx, task); err != nil {
		return fmt.Errorf("failed to save task: %w", err)
	}

	for _, event := range task.PopEvents() {
		if err := p.eventDispatcher.Dispatch(ctx, event); err != nil {
			return err
		}
	}

	return nil
}

func (p *ProjectApp) AddTaskToProject(ctx context.Context, taskID string, projectID string) error {
	task, err := p.taskRepo.FindById(ctx, domain.NewTaskID(taskID))
	if err != nil {
		return fmt.Errorf("failed to find task: %w", err)
	}
	project, err := p.projectRepo.FindById(ctx, domain.NewProjectID(projectID))
	if err != nil {
		return fmt.Errorf("failed to find project: %w", err)
	}

	if err := project.AddTask(task); err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	if err := p.taskRepo.Save(ctx, task); err != nil {
		return fmt.Errorf("failed to save task: %w", err)
	}
	if err := p.projectRepo.Save(ctx, project); err != nil {
		return fmt.Errorf("failed to save project: %w", err)
	}

	return nil
}
