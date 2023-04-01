package api

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for TaskStatus.
const (
	Completed           TaskStatus = "completed"
	InDevelopment       TaskStatus = "in_development"
	InReview            TaskStatus = "in_review"
	ReadyForDeploy      TaskStatus = "ready_for_deploy"
	ReadyForDevelopment TaskStatus = "ready_for_development"
	ReadyForReview      TaskStatus = "ready_for_review"
	Unscheduled         TaskStatus = "unscheduled"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// ParticipantID defines model for ParticipantID.
type ParticipantID = openapi_types.UUID

// PostTask defines model for PostTask.
type PostTask struct {
	Description string `json:"description"`
	Title       string `json:"title"`
}

// Task defines model for Task.
type Task struct {
	Assignee    ParticipantID  `json:"assignee"`
	Description string         `json:"description"`
	Estimation  TaskEstimation `json:"estimation"`
	Id          UUID           `json:"id"`
	Status      TaskStatus     `json:"status"`
	Title       string         `json:"title"`
}

// TaskEstimation defines model for TaskEstimation.
type TaskEstimation = uint64

// TaskStatus defines model for TaskStatus.
type TaskStatus string

// UUID defines model for UUID.
type UUID = openapi_types.UUID

// EstimateTaskJSONBody defines parameters for EstimateTask.
type EstimateTaskJSONBody struct {
	Estimation TaskEstimation `json:"estimation"`
}

// AssignParticipantToTaskJSONBody defines parameters for AssignParticipantToTask.
type AssignParticipantToTaskJSONBody struct {
	ParticipantID ParticipantID `json:"participantID"`
}

// TransitTaskStatusJSONBody defines parameters for TransitTaskStatus.
type TransitTaskStatusJSONBody struct {
	Status TaskStatus `json:"status"`
}

// EstimateTaskJSONRequestBody defines body for EstimateTask for application/json ContentType.
type EstimateTaskJSONRequestBody EstimateTaskJSONBody

// CreateTaskJSONRequestBody defines body for CreateTask for application/json ContentType.
type CreateTaskJSONRequestBody = PostTask

// UpdateTaskJSONRequestBody defines body for UpdateTask for application/json ContentType.
type UpdateTaskJSONRequestBody = PostTask

// AssignParticipantToTaskJSONRequestBody defines body for AssignParticipantToTask for application/json ContentType.
type AssignParticipantToTaskJSONRequestBody AssignParticipantToTaskJSONBody

// TransitTaskStatusJSONRequestBody defines body for TransitTaskStatus for application/json ContentType.
type TransitTaskStatusJSONRequestBody TransitTaskStatusJSONBody
