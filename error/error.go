package error

import (
	"fmt"
	"net/http"
)

// Error represents an error in provider layer.
//Error Response model
//swagger:response Error
type Error struct {
	// Error Message
	DevMessage string    `json:"devMessage"`
	// Error Arguments
	Arg        string    `json:"arg"`
	// Error Code
	Code       ErrorCode `json:"errorCode"`
}

// NewError returns a new Error.
func NewError(code ErrorCode, msg string) error {
	err := Error{}
	err.DevMessage = fmt.Sprintf("%v", msg)
	err.Code = code
	return err
}

// Error returns the error message associated with the error
func (err Error) Error() string {
	return fmt.Sprintf("%d:%s", err.Code, err.DevMessage)
}

// GetHTTPStatusCode returns HTTP code based on code
func (err Error) GetHTTPStatusCode(errorCodeMappings ...map[ErrorCode]int) int {
	httpStatusCode, ok := ErrorCodeToHTTPStatusCodeMapping[err.Code]
	if ok {
		return httpStatusCode
	}

	for i := range errorCodeMappings {
		httpStatusCode, ok = (errorCodeMappings[i])[err.Code]
		if ok {
			return httpStatusCode
		}
	}

	return http.StatusInternalServerError
}