package nbtag

import "fmt"

type NBError struct {
	err     string
	Context []string
}

func newErrorf(reader NBReader, format string, args ...interface{}) NBError {
	return NBError{err: fmt.Sprintf(format, args...), Context: reader.Context()}
}

func (e NBError) Error() string {
	return e.err
}
