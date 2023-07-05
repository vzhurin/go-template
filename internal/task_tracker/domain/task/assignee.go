package task

type Assignee struct {
	id UserID
}

func NewAssignee(id UserID) (Assignee, error) {
	return Assignee{id: id}, nil
}

func (a Assignee) String() string {
	return a.id.String()
}
