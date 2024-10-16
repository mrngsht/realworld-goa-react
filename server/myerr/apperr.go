package myerr

import (
	"errors"
)

type appErr struct {
	text string
}

func NewAppErr(text string) error {
	return &appErr{text: text}
}

func (ae *appErr) Error() string {
	return ae.text
}

func AsAppErr(err error) (*appErr, bool) {
	if ae := (*appErr)(nil); errors.As(err, &ae) {
		return ae, true
	}
	return nil, false
}
