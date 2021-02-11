// Package wraperr implements methods for wrapping error in Go.

package wraperr

import (
	"errors"
)

// wrapError is an implementation of error that has both the
// outer and inner errors.
type wrapError struct {
	Outer error
	Inner error
}

// Wrap defines that outer wraps inner
func Wrap(outer, inner error) error {
	return &wrapError{
		Outer: outer,
		Inner: inner,
	}
}

// Wrapf wraps an error
func Wrapf(err error, format string) error {

	outMsg := "<nil>"
	if err != nil {
		outMsg = err.Error()
	}
	format = format + " : " + outMsg
	outer := errors.New(format)

	return Wrap(outer, err)
}

// Contains checks if the given error contains an error with the message
func Contains(err error, msg string) bool {
	return len(GetAll(err, msg)) > 0
}

// GetAll gets all the errors that might be wrapped in err with the given message.
func GetAll(err error, msg string) []error {
	var result []error
	newErr := err
	for {
		if newErr.Error() == msg {
			result = append(result, newErr)
		}
		newErr = Walk(newErr)
		if newErr == nil {
			break
		}
	}

	return result
}

// Walk walks all the wrapped errors in err
func Walk(err error) error {
	if err == nil {
		return nil
	}

	switch e := err.(type) {
	case *wrapError:
		return e.Inner
	default:
		return nil
	}
}

func (w *wrapError) Error() string {
	return w.Outer.Error()
}
