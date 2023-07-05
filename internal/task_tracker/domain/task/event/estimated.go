package event

import (
	"time"

	"github.com/google/uuid"
)

type Estimated struct {
	event
	taskID     uuid.UUID
	estimation uint64
}

func NewEstimated(id uuid.UUID, occurredAt time.Time, taskID uuid.UUID, estimation uint64) *Estimated {
	return &Estimated{
		event: event{
			id:         id,
			occurredAt: occurredAt,
		},
		taskID:     taskID,
		estimation: estimation,
	}
}
