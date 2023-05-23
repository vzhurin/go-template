package task

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/vzhurin/template/internal/task_tracker/domain/model/task/event"
	"time"
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
		events:      []event.Event{event.NewCreated(uuid.New(), time.Now(), id, title, description)},
	}
}

func (t *Task) ID() ID {
	return t.id
}

func (t *Task) Describe(title Title, description Description) {
	t.title = title
	t.description = description
	t.events = append(t.events, event.NewDescribed(uuid.New(), time.Now(), t.id, title, description))
}

func (t *Task) Assign(assignee Assignee) {
	t.assignee = assignee
	t.events = append(t.events, event.NewAssigned(uuid.New(), time.Now(), t.id, assignee))
}

func (t *Task) SwitchStatus(status Status) error {
	for _, s := range statusTransitions[t.status] {
		if s == status {
			t.status = status
			t.events = append(t.events, event.NewStatusSwitched(uuid.New(), time.Now(), t.id, status))

			return nil
		}
	}

	return fmt.Errorf("invalid transition: %s -> %s", t.status, status)
}

func (t *Task) Estimate(estimation Estimation) {
	t.estimation = estimation
	t.events = append(t.events, event.NewEstimated(uuid.New(), time.Now(), t.id, estimation))
}

func UnmarshalFromDatabase(
	id string,
	title string,
	description string,
	assignee string,
	status string,
	estimation uint64,
) *Task {
	i, _ := NewID(id)
	t, _ := NewTitle(title)
	d, _ := NewDescription(description)
	a, _ := NewAssignee(assignee)
	s := Status{status: status}
	e := NewEstimation(estimation)

	return &Task{
		id:          i,
		title:       t,
		description: d,
		assignee:    a,
		status:      s,
		estimation:  e,
	}
}

func (t *Task) MarshalToDatabase() map[string]interface{} {
	m := make(map[string]interface{})

	m["id"] = t.ID().String()
	m["title"] = t.title.String()
	m["description"] = t.description.String()
	m["assignee"] = t.assignee.String()
	m["status"] = t.status.String()
	m["estimation"] = t.estimation.estimation

	return m
}
