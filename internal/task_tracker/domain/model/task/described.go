package task

import (
	"time"

	"github.com/google/uuid"
)

type Described struct {
	event
	taskID      ID
	title       Title
	description Description
}

func NewDescribed(id uuid.UUID, occurredAt time.Time, taskID ID, title Title, description Description) *Described {
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
