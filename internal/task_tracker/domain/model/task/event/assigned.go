package event

import (
	"github.com/google/uuid"
	"github.com/vzhurin/template/internal/task_tracker/domain/model/task"
	"time"
)

type Assigned struct {
	event
	taskID   task.ID
	assignee task.Assignee
}

func NewAssigned(id uuid.UUID, occurredAt time.Time, taskID task.ID, assignee task.Assignee) *Assigned {
	return &Assigned{
		event: event{
			id:         id,
			occurredAt: occurredAt,
		},
		taskID:   taskID,
		assignee: assignee,
	}
}
