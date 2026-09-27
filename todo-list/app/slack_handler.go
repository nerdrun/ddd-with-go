package app

import (
	"context"
	"todo/domain"
)

type SlackTaskCompletedHandler struct {
	notifier NotificationAlert
}

func NewSlackTaskCompletedHandler(notifier NotificationAlert) *SlackTaskCompletedHandler {
	return &SlackTaskCompletedHandler{
		notifier: notifier,
	}
}

func (h *SlackTaskCompletedHandler) Handle(ctx context.Context, event domain.Event) error {
	e, ok := event.(domain.TaskCompleted)
	if !ok {
		return nil
	}

	err := h.notifier.SendTaskCompleteAlert(ctx, e.TaskID.String(), e.ProjectID.String())

	if err != nil {
		return nil
	}

	return nil
}
