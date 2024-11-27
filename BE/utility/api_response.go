package utility

type APIResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func ResponseSuccess(message string, data any) APIResponse {
	return APIResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func ResponseError(message string, err any) APIResponse {
	return APIResponse{
		Status:  false,
		Message: message,
		Error:   err,
	}
}
