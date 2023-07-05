package event

import (
	"time"

	"github.com/google/uuid"
)

type Described struct {
	event
	taskID      uuid.UUID
	title       string
	description string
}

func NewDescribed(id uuid.UUID, occurredAt time.Time, taskID uuid.UUID, title string, description string) *Described {
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
