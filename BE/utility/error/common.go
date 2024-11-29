package error

import "errors"

const (
	MESSAGE_FAILED_TO_GET_DATA_FROM_BODY = "Failed to get data from body"
)

var (
	ErrInternalServer = errors.New("Internal server error")
	ErrBadRequest     = errors.New("Bad request")
)
