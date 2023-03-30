package task

import (
	"fmt"
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
}

func NewTask(id ID, title Title, description Description) *Task {
	// TODO domain event

	return &Task{
		id:          id,
		title:       title,
		description: description,
		status:      Unscheduled,
	}
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

func (t *Task) ID() ID {
	return t.id
}

func (t *Task) Describe(title Title, description Description) {
	t.title = title
	t.description = description

	// TODO domain event
}

func (t *Task) Assign(assignee Assignee) {
	t.assignee = assignee

	// TODO domain event
}

func (t *Task) SwitchStatus(status Status) error {
	for _, s := range statusTransitions[t.status] {
		if s == status {
			t.status = status
			return nil
		}
	}

	// TODO domain event

	return fmt.Errorf("invalid transition: %s -> %s", t.status, status)
}

func (t *Task) Estimate(estimation Estimation) {
	t.estimation = estimation

	// TODO domain event
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
