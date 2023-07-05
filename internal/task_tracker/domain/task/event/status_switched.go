package event

import (
	"time"

	"github.com/google/uuid"
)

type StatusSwitched struct {
	event
	taskID uuid.UUID
	status string
}

func NewStatusSwitched(id uuid.UUID, occurredAt time.Time, taskID uuid.UUID, status string) *StatusSwitched {
	return &StatusSwitched{
		event: event{
			id:         id,
			occurredAt: occurredAt,
		},
		taskID: taskID,
		status: status,
	}
}
