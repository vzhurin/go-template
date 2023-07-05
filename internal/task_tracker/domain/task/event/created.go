package event

import (
	"time"

	"github.com/google/uuid"
)

type Created struct {
	event
	taskID      uuid.UUID
	title       string
	description string
}

func NewCreated(id uuid.UUID, occurredAt time.Time, taskID uuid.UUID, title string, description string) *Created {
	return &Created{
		event: event{
			id:         id,
			occurredAt: occurredAt,
		},
		taskID:      taskID,
		title:       title,
		description: description,
	}
}
