package task

import (
	"time"

	"github.com/google/uuid"
)

type Assigned struct {
	event
	taskID   ID
	assignee Assignee
}

func NewAssigned(id uuid.UUID, occurredAt time.Time, taskID ID, assignee Assignee) *Assigned {
	return &Assigned{
		event: event{
			id:         id,
			occurredAt: occurredAt,
		},
		taskID:   taskID,
		assignee: assignee,
	}
}
