package app

import "context"

type NotificationAlert interface {
	SendTaskCompleteAlert(ctx context.Context, taskID, projectID string) error
}
