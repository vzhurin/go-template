package task

import "errors"

type Title struct {
	title string
}

func NewTitle(t string) (Title, error) {
	if len(t) > 255 || t == "" {
		return Title{}, errors.New("title length must greater then 0 and less than 255")
	}

	return Title{title: t}, nil
}

func (t Title) String() string {
	return t.title
}
