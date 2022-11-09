package helper

import "net/http"

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func APIResponse(message string, code int, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  http.StatusText(code),
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return response
}
