package task

import "github.com/google/uuid"

type ID struct {
	uuid uuid.UUID
}

func NewID(i string) (ID, error) {
	parsed, err := uuid.Parse(i)
	if err != nil {
		return ID{}, err
	}

	return ID{uuid: parsed}, nil
}

func (id ID) String() string {
	return id.uuid.String()
}
