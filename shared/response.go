package shared

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func BuildResponse(status string, data interface{}) Response {
	res := Response{
		Status: status,
		Data:   data,
	}
	return res
}

func BuildErrorResponse(status string, errorMessage string) ErrorResponse {
	res := ErrorResponse{
		Status:       status,
		ErrorMessage: errorMessage,
	}
	return res
}
