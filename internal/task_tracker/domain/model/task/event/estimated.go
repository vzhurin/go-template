package event

import (
	"github.com/google/uuid"
	"github.com/vzhurin/template/internal/task_tracker/domain/model/task"
	"time"
)

type Estimated struct {
	event
	taskID     task.ID
	estimation task.Estimation
}

func NewEstimated(id uuid.UUID, occurredAt time.Time, taskID task.ID, estimation task.Estimation) *Estimated {
	return &Estimated{
		event: event{
			id:         id,
			occurredAt: occurredAt,
		},
		taskID:     taskID,
		estimation: estimation,
	}
}
