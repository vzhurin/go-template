package task

import "github.com/google/uuid"

type UserID struct {
	uuid uuid.UUID
}

func NewUserID(i string) (UserID, error) {
	parsed, err := uuid.Parse(i)
	if err != nil {
		return UserID{}, err
	}

	return UserID{uuid: parsed}, nil
}

func (id UserID) String() string {
	return id.uuid.String()
}
