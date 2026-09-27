package infrastructure

import "context"

type SlackNotification struct {
	webhookURL string
}

func NewSlackNotification(webhookURL string) *SlackNotification {
	return &SlackNotification{
		webhookURL: webhookURL,
	}
}

func (n *SlackNotification) SendTaskCompleteAlert(ctx context.Context, taskID string, projectID string) error {
	// DONE
	return nil
}
