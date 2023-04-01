package event

import (
	"github.com/google/uuid"
	"github.com/vzhurin/template/internal/task_tracker/domain/model/task"
	"time"
)

type StatusSwitched struct {
	event
	taskID task.ID
	status task.Status
}

func NewStatusSwitched(id uuid.UUID, occurredAt time.Time, taskID task.ID, status task.Status) *StatusSwitched {
	return &StatusSwitched{
		event: event{
			id:         id,
			occurredAt: occurredAt,
		},
		taskID: taskID,
		status: status,
	}
}
