package errorwrapper

import "net/http"

// Common HTTP Status Code
var (
	// 2xx Success
	StatusOK StatusCode = 200

	// 4xx client errors
	StatusBadRequest   StatusCode = 400
	StatusUnauthorized StatusCode = 401
	StatusForbidden    StatusCode = 403
	StatusNotFound     StatusCode = 404

	// 5xx server errors
	StatusInternalServerError StatusCode = 500
)

// Invalid Condition / Status
var (
	StatusInvalidUsernameErr StatusCode = 1001
)

// Message
var (
	StatusOKMessage = "Success"

	StatusBadRequestMessage          = "Bad request"
	StatusUnauthorizedMessage        = "Unauthorized"
	StatusForbiddenMessage           = "Forbidden"
	StatusInternalServerErrorMessage = "Internal server error"

	// Invalid Condition / Status
	StatusInvalidUsernameErrMessage = "Invalid username"
)

// map Status Code to HTTP Status
var (
	errHTTPStatus = map[StatusCode]int{
		StatusOK:                  http.StatusOK,
		StatusBadRequest:          http.StatusBadRequest,
		StatusUnauthorized:        http.StatusUnauthorized,
		StatusForbidden:           http.StatusForbidden,
		StatusInternalServerError: http.StatusInternalServerError,

		// Invalid Condition / Status
		StatusInvalidUsernameErr: http.StatusUnauthorized,
	}
)

// map Status code to Message
var (
	errStatusMessage = map[StatusCode]string{
		StatusOK:                  StatusOKMessage,
		StatusBadRequest:          StatusBadRequestMessage,
		StatusUnauthorized:        StatusUnauthorizedMessage,
		StatusForbidden:           StatusForbiddenMessage,
		StatusInternalServerError: StatusInternalServerErrorMessage,

		// Invalid Condition / Status
		StatusInvalidUsernameErr: StatusInvalidUsernameErrMessage,
	}
)
