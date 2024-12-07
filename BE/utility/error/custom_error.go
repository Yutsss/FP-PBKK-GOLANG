package error

import (
	"errors"
	"net/http"
)

type CustomError interface {
	Error() string
	Code() int
	UnWrap() error
}

type customError struct {
	err  error
	code int
}

func NewCustomError(err error, code int) CustomError {
	return &customError{
		err:  err,
		code: code,
	}
}

func (e *customError) Error() string {
	return e.err.Error()
}

func (e *customError) Code() int {
	return e.code
}

func (e *customError) UnWrap() error {
	return e.err
}

const (
	MESSAGE_FAILED_TO_GET_DATA_FROM_BODY = "Failed to get data from body"
	MESSAGE_FAILED_TO_AUTHORIZE_USER     = "Failed to authorize user"
	MESSAGE_FAILED_REGISTER_USER         = "Failed to register user"
	MESSAGE_FAILED_LOGIN_USER            = "Failed to login user"
	MESSAGE_FAILED_GET_USER              = "Failed to get user"
	MESSAGE_FAILED_GET_ALL_USER          = "Failed to get all user"
	MESSAGE_FAILED_LOGOUT_USER           = "Failed to logout user"
	MESSAGE_FAILED_CREATE_TICKET         = "Failed to create ticket"
	MESSAGE_FAILED_GET_ALL_TICKET        = "Failed to get all ticket"
)

var (
	ErrInternalServer       = NewCustomError(errors.New("Internal server error"), http.StatusInternalServerError)
	ErrBadRequest           = NewCustomError(errors.New("Bad request"), http.StatusBadRequest)
	ErrNotAllowed           = NewCustomError(errors.New("Not allowed"), http.StatusMethodNotAllowed)
	ErrEmailAlreadyExist    = NewCustomError(errors.New("Email already exist"), http.StatusConflict)
	ErrWrongEmailOrPassword = NewCustomError(errors.New("Wrong email or password"), http.StatusBadRequest)
	ErrUnauthorized         = NewCustomError(errors.New("Unauthorized"), http.StatusUnauthorized)
	ErrUserNotFound         = NewCustomError(errors.New("User not found"), http.StatusNotFound)
)
