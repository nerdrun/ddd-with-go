package todo_list

import (
	"context"
	"fmt"
	"todo/app"
	"todo/infrastructure"
)

func main() {
	ctx := context.Background()
	projectRepo := infrastructure.NewProjectDB()
	taskRepo := infrastructure.NewTaskDB()
	eventDispatcher := infrastructure.NewInMemoryEventDispatcher()
	slackNotification := infrastructure.NewSlackNotification("URL")

	projectHandler := app.NewProjectTaskCompletedHandler(projectRepo)
	slackHandler := app.NewSlackTaskCompletedHandler(slackNotification)

	eventDispatcher.Subscribe("TaskCompleted", projectHandler.Handle)
	eventDispatcher.Subscribe("TaskCompleted", slackHandler.Handle)

	appService := app.NewProjectApp(projectRepo, taskRepo, eventDispatcher)

	// 5. Run a Use Case!
	err := appService.CompleteTask(ctx, "task-123")
	if err != nil {
		fmt.Println("Error completing task:", err)
	}
}
