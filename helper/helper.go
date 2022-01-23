package helper

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func APIResponse(status string, message string, code int, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return response
}
