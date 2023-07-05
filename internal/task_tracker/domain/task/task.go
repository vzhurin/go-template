package task

import (
	"fmt"
	"github.com/vzhurin/template/internal/task_tracker/domain/task/event"
	"time"

	"github.com/google/uuid"
)

var statusTransitions = map[Status][]Status{
	Unscheduled:         {ReadyForDevelopment},
	ReadyForDevelopment: {InDevelopment},
	InDevelopment:       {ReadyForReview},
	ReadyForReview:      {InReview},
	InReview:            {ReadyForDeploy},
	ReadyForDeploy:      {Completed},
	Completed:           {},
}

type Task struct {
	id          ID
	title       Title
	description Description
	assignee    Assignee
	status      Status
	estimation  Estimation
	events      []event.Event
}

func NewTask(id ID, title Title, description Description) *Task {
	return &Task{
		id:          id,
		title:       title,
		description: description,
		status:      Unscheduled,
		events:      []event.Event{event.NewCreated(uuid.New(), time.Now(), id.uuid, title.title, description.description)},
	}
}

func (t *Task) ID() ID {
	return t.id
}

func (t *Task) Describe(title Title, description Description) {
	t.title = title
	t.description = description
	t.events = append(t.events, event.NewDescribed(uuid.New(), time.Now(), t.id.uuid, title.title, description.description))
}

func (t *Task) Assign(assignee Assignee) {
	t.assignee = assignee
	t.events = append(t.events, event.NewAssigned(uuid.New(), time.Now(), t.id.uuid, assignee.id.uuid))
}

func (t *Task) SwitchStatus(status Status) error {
	for _, s := range statusTransitions[t.status] {
		if s == status {
			t.status = status
			t.events = append(t.events, event.NewStatusSwitched(uuid.New(), time.Now(), t.id.uuid, status.status))

			return nil
		}
	}

	return fmt.Errorf("invalid transition: %s -> %s", t.status, status)
}

func (t *Task) Estimate(estimation Estimation) {
	t.estimation = estimation
	t.events = append(t.events, event.NewEstimated(uuid.New(), time.Now(), t.id.uuid, estimation.estimation))
}

func (t *Task) Title() Title {
	return t.title
}

func (t *Task) Description() Description {
	return t.description
}

func (t *Task) Assignee() Assignee {
	return t.assignee
}

func (t *Task) Status() Status {
	return t.status
}

func (t *Task) Estimation() Estimation {
	return t.estimation
}
