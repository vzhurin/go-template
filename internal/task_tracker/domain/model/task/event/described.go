package event

import (
	"github.com/google/uuid"
	"github.com/vzhurin/template/internal/task_tracker/domain/model/task"
	"time"
)

type Described struct {
	event
	taskID      task.ID
	title       task.Title
	description task.Description
}

func NewDescribed(id uuid.UUID, occurredAt time.Time, taskID task.ID, title task.Title, description task.Description) *Described {
	return &Described{
		event: event{
			id:         id,
			occurredAt: occurredAt,
		},
		taskID:      taskID,
		title:       title,
		description: description,
	}
}
