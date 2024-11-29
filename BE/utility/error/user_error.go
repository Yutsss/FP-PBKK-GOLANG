package error

import "errors"

const (
	MESSAGE_FAILED_REGISTER_USER = "Failed to register user"
	MESSAGE_FAILED_LOGIN_USER    = "Failed to login user"
	MESSAGE_FAILED_GET_USER      = "Failed to get user"
	MESSAGE_FAILED_LOGOUT_USER   = "Failed to logout user"
)

var (
	ErrEmailAlreadyExist = errors.New("Email already exist")
)
