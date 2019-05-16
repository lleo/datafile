package datafile

import (
	"fmt"
)

var _ = fmt.Print

type txn struct {
	tasks []task
}

type task interface {
	Execute() error
	Cancel() error
}

type createDatafileTask struct {
	filename string
	pageSize uint32
}

func (t *createDatafileTask) Execute() error {
	return NotImplemented("Execute")
}

func (t *createDatafileTask) Cancel() error {
	return NotImplemented("Cancel")
}

type openDatafileTask struct {
	filename string
}

func (t *openDatafileTask) Execute() error {
	return NotImplemented("Execute")
}

func (t *openDatafileTask) Cancel() error {
	return NotImplemented("Cancel")
}

type writeHeadersTask struct {
}

func (t *writeHeadersTask) Execute() error {
	return NotImplemented("Execute")
}

func (t *writeHeadersTask) Cancel() error {
	return NotImplemented("Cancel")
}

//type readHeadersTask struct {
//
//}
//
//func (t *readHeadersTask) Execute() error {
//	return NotImplemented("Execute")
//}
//
//func (t *readHeadersTask) Cancel() error {
//	return NotImplemented("Cancel")
//}
