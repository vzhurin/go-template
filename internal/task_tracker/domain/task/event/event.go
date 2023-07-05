package event

import (
	"time"

	"github.com/google/uuid"
)

type event struct {
	id         uuid.UUID
	occurredAt time.Time
}

func (e *event) ID() uuid.UUID {
	return e.id
}

func (e *event) OccurredAt() time.Time {
	return e.occurredAt
}

type Event interface {
	ID() uuid.UUID
	OccurredAt() time.Time
}
