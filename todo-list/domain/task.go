package domain

import (
	"errors"
	"time"
)

const (
	Incomplete TaskStatus = "INCOMPLETE"
	Complete   TaskStatus = "COMPLETE"
)

type TaskID string

func (t TaskID) String() string {
	return string(t)
}

type TaskStatus string

type Task struct {
	id        TaskID
	projectID ProjectID
	title     string
	status    TaskStatus
	events    []Event
}

func NewTask(id TaskID, projectID ProjectID, title string) *Task {
	return &Task{
		id:        id,
		projectID: projectID,
		title:     title,
		status:    Incomplete,
		events:    make([]Event, 0),
	}
}

func NewTaskID(taskID string) TaskID {
	return TaskID(taskID)
}

func (t *Task) AssignToProject(projectID ProjectID) error {
	if t.projectID != "" {
		return errors.New("cannot assign task to a project")
	}
	t.projectID = projectID
	return nil
}

func (t *Task) Complete() {
	if t.status == Complete {
		return
	}
	t.status = Complete
	t.events = append(t.events, TaskCompleted{
		TaskID:     t.id,
		ProjectID:  t.projectID,
		OccurredAt: time.Now(),
	})
}

func (t *Task) PopEvents() []Event {
	events := t.events
	t.events = nil
	return events
}

func (t *Task) ID() TaskID         { return t.id }
func (t *Task) Status() TaskStatus { return t.status }
