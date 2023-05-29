package task

import (
	"time"

	"github.com/google/uuid"
)

type StatusSwitched struct {
	event
	taskID ID
	status Status
}

func NewStatusSwitched(id uuid.UUID, occurredAt time.Time, taskID ID, status Status) *StatusSwitched {
	return &StatusSwitched{
		event: event{
			id:         id,
			occurredAt: occurredAt,
		},
		taskID: taskID,
		status: status,
	}
}
