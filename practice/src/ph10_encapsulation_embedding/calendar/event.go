package calendar

import "errors"

type Event struct {
	title string
	Date
}

func (e *Event) SetTitle(title string) error {
	if len(title) > 5 {
		return errors.New("invalid title!")
	}
	e.title = title
	return nil
}

func (e *Event) Title() string {
	return e.title
}
