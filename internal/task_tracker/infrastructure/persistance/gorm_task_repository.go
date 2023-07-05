package persistance

import (
	"context"
	"github.com/google/uuid"
	"github.com/vzhurin/template/internal/task_tracker/domain/task"
	"gorm.io/gorm"
)

type GormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) *GormTaskRepository {
	return &GormTaskRepository{db: db}
}

func (r *GormTaskRepository) Get(ctx context.Context, id task.ID) (*task.Task, error) {
	var dto GormTask
	err := r.db.WithContext(ctx).First(&dto, "uuid = ?", id.String()).Error
	if err != nil {
		return nil, err
	}

	return dto.ToEntity()
}

func (r *GormTaskRepository) Save(ctx context.Context, task *task.Task) (*task.Task, error) {
	dto := NewTaskGorm(task)
	err := r.db.WithContext(ctx).Save(dto).Error
	if err != nil {
		return nil, err
	}

	task, err = dto.ToEntity()
	if err != nil {
		return nil, err
	}

	return task, nil
}

// TODO go types
type GormTask struct {
	ID          uuid.UUID `gorm:"primary_key"`
	Title       string
	Description string
	Assignee    uuid.UUID
	Status      string
	Estimation  uint64
}

func NewTaskGorm(task *task.Task) *GormTask {
	id, _ := uuid.Parse(task.ID().String())
	assignee, _ := uuid.Parse(task.Assignee().String())

	return &GormTask{
		ID:          id,
		Title:       task.Title().String(),
		Description: task.Description().String(),
		Assignee:    assignee,
		Status:      task.Status().String(),
		Estimation:  task.Estimation().Uint64(),
	}
}

func (t *GormTask) ToEntity() (*task.Task, error) {
	id, err := task.NewID(t.ID.String())
	if err != nil {
		return nil, err
	}

	title, err := task.NewTitle(t.Title)
	if err != nil {
		return nil, err
	}

	description, err := task.NewDescription(t.Description)
	if err != nil {
		return nil, err
	}

	return task.NewTask(
		id,
		title,
		description,
	), nil
}
