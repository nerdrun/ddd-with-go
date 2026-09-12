package app

import (
	"context"
	"fmt"
	"todo/domain"
)

// ProjectTaskCompletedHandler reacts to TaskCompleted events by updating Project capacity.
type ProjectTaskCompletedHandler struct {
	projectRepo domain.ProjectRepository
}

func NewProjectTaskCompletedHandler(projectRepo domain.ProjectRepository) *ProjectTaskCompletedHandler {
	return &ProjectTaskCompletedHandler{
		projectRepo: projectRepo,
	}
}

// Handle executes when a TaskCompleted event is dispatched.
func (h *ProjectTaskCompletedHandler) Handle(ctx context.Context, event domain.Event) error {
	e, ok := event.(domain.TaskCompleted)
	if !ok {
		return nil
	}

	if e.ProjectID == "" {
		return nil
	}

	project, err := h.projectRepo.FindById(ctx, e.ProjectID)
	if err != nil {
		return fmt.Errorf("handler failed to find project: %w", err)
	}

	project.RemoveTask(e.TaskID)

	if err := h.projectRepo.Save(ctx, project); err != nil {
		return fmt.Errorf("handler failed to save project: %w", err)
	}

	return nil
}
