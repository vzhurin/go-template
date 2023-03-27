package query

import "github.com/google/uuid"

const (
	Completed           TaskStatus = "completed"
	InDevelopment       TaskStatus = "in_development"
	InReview            TaskStatus = "in_review"
	ReadyForDeploy      TaskStatus = "ready_for_deploy"
	ReadyForDevelopment TaskStatus = "ready_for_development"
	ReadyForReview      TaskStatus = "ready_for_review"
	Unscheduled         TaskStatus = "unscheduled"
)

type TaskStatus string

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Assignee    uuid.UUID
	Status      TaskStatus
	Estimation  uint64
}
