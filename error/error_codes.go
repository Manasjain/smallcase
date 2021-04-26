package error

import "net/http"

type ErrorCode int

// Basic Service Side Errors - Use error codes from 4000 to 4999
const (
	// InternalServerError ...
	InternalServerError ErrorCode = 4000

	// BadRequest ...
	BadRequest ErrorCode = 4002
)

// DatabaseFailures - Use error codes from 2000 to 2999
const (
	// DatabaseServiceFailure
	DatabaseServiceFailure ErrorCode = 2000

	// DatabaseRecordNotFound ...
	DatabaseRecordNotFound ErrorCode = 2001
)

// ErrorCodeToHTTPStatusCodeMapping ...
var ErrorCodeToHTTPStatusCodeMapping = map[ErrorCode]int{
	InternalServerError:    http.StatusInternalServerError,
	BadRequest:             http.StatusBadRequest,
	DatabaseServiceFailure: http.StatusInternalServerError,
	DatabaseRecordNotFound: http.StatusBadRequest,
}

func (code ErrorCode) String() string {
	strings := map[ErrorCode]string{
		InternalServerError:    "InternalServerError",
		BadRequest:             "BadRequest",
		DatabaseServiceFailure: "DatabaseServiceFailure",
		DatabaseRecordNotFound: "DatabaseRecordNotFound",
	}
	return strings[code]
}
