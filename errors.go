package datafile

import (
	"github.com/pkg/errors"
)

var ErrNotImplemented = errors.New("Not Implemented")
var ErrPrevExists = errors.New("Previous DataFile stuct exists")

func NotImplemented(where string) error {
	return errors.Wrapf(ErrNotImplemented, "func %q", where)
}

func PrevExists(where string) error {
	return errors.Wrapf(ErrPrevExists, "func %q", where)
}
