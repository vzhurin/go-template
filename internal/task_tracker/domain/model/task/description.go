package task

import "errors"

type Description struct {
	description string
}

func NewDescription(d string) (Description, error) {
	if len(d) > 1000 {
		return Description{}, errors.New("description less than 1000")
	}

	return Description{description: d}, nil
}

func (d Description) String() string {
	return d.description
}
