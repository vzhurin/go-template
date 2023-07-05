package task

var (
	Completed           = Status{"completed"}
	InDevelopment       = Status{"in_development"}
	InReview            = Status{"in_review"}
	ReadyForDeploy      = Status{"ready_for_deploy"}
	ReadyForDevelopment = Status{"ready_for_development"}
	ReadyForReview      = Status{"ready_for_review"}
	Unscheduled         = Status{"unscheduled"}
)

type Status struct {
	status string
}

func (s Status) String() string {
	return s.status
}
