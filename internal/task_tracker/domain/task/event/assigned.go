package event

import (
	"time"

	"github.com/google/uuid"
)

type Assigned struct {
	event
	taskID   uuid.UUID
	assignee uuid.UUID
}

func NewAssigned(id uuid.UUID, occurredAt time.Time, taskID uuid.UUID, assignee uuid.UUID) *Assigned {
	return &Assigned{
		event: event{
			id:         id,
			occurredAt: occurredAt,
		},
		taskID:   taskID,
		assignee: assignee,
	}
}
