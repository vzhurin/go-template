package task

import "github.com/google/uuid"

type Assignee struct {
	uuid uuid.UUID
}

func NewAssignee(i string) (Assignee, error) {
	parsed, err := uuid.Parse(i)
	if err != nil {
		return Assignee{}, err
	}

	return Assignee{uuid: parsed}, nil
}

func (a Assignee) String() string {
	return a.uuid.String()
}
