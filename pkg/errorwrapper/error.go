package errorwrapper

import (
	"errors"
	"fmt"
	"net/http"
)

type StatusCode int16

type ErrorWrapper struct {
	HttpStatus int
	Code       StatusCode
	Message    string
	devMessage error
}

func New(code StatusCode, err error, message string) error {
	// get http status
	httpStatus, ok := errHTTPStatus[code]
	if !ok {
		httpStatus = http.StatusInternalServerError
	}

	// get msg
	var msg string
	if message != "" {
		msg = message
	} else {
		// get msg
		msg, ok = errStatusMessage[code]
		if !ok {
			msg = StatusInternalServerErrorMessage
		}
	}

	if err == nil {
		err = errors.New("")
	}

	mod := &ErrorWrapper{
		HttpStatus: httpStatus,
		Code:       code,
		Message:    msg,
		devMessage: err,
	}

	return mod
}

func (e *ErrorWrapper) Error() string {
	output := fmt.Sprintf("%v %s", e.Code, e.Message)
	if e.devMessage != nil {
		output = fmt.Sprintf("%s %s", output, e.devMessage.Error())
	}

	return output
}

func CastToErrorWrapper(err error) *ErrorWrapper {
	if err == nil {
		return nil
	}

	errWrapper, ok := err.(*ErrorWrapper)
	if !ok || errWrapper == nil {
		return New(StatusInternalServerError, err, "").(*ErrorWrapper)
	}

	return errWrapper
}
