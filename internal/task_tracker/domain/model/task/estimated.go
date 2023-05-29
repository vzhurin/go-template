package task

import (
	"time"

	"github.com/google/uuid"
)

type Estimated struct {
	event
	taskID     ID
	estimation Estimation
}

func NewEstimated(id uuid.UUID, occurredAt time.Time, taskID ID, estimation Estimation) *Estimated {
	return &Estimated{
		event: event{
			id:         id,
			occurredAt: occurredAt,
		},
		taskID:     taskID,
		estimation: estimation,
	}
}
