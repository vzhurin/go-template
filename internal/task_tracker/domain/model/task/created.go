package task

import (
	"time"

	"github.com/google/uuid"
)

type Created struct {
	event
	taskID      ID
	title       Title
	description Description
}

func NewCreated(id uuid.UUID, occurredAt time.Time, taskID ID, title Title, description Description) *Created {
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
