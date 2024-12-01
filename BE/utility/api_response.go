package utility

type APIResponse struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func ResponseSuccess(message string, data any, code int) APIResponse {
	return APIResponse{
		Status:  true,
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func ResponseError(message string, err string, code int) APIResponse {
	return APIResponse{
		Status:  false,
		Code:    code,
		Message: message,
		Error:   err,
	}
}
