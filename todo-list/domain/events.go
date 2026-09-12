package domain

import (
	"context"
	"time"
)

type TaskCompleted struct {
	TaskID     TaskID
	ProjectID  ProjectID
	OccurredAt time.Time
}

func (t TaskCompleted) EventName() string {
	return "TaskCompleted"
}

// Event interface marks any domain event
type Event interface {
	EventName() string
}

// EventDispatcher is the contract for publishing events
type EventDispatcher interface {
	Dispatch(ctx context.Context, event Event) error
}
