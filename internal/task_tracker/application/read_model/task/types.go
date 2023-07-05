package task

import "github.com/google/uuid"

const (
	Completed           Status = "completed"
	InDevelopment       Status = "in_development"
	InReview            Status = "in_review"
	ReadyForDeploy      Status = "ready_for_deploy"
	ReadyForDevelopment Status = "ready_for_development"
	ReadyForReview      Status = "ready_for_review"
	Unscheduled         Status = "unscheduled"
)

type Status string

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Assignee    uuid.UUID
	Status      Status
	Estimation  uint64
}
