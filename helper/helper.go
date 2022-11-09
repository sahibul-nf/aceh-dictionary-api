package helper

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type Response struct {
	Meta  Meta        `json:"meta"`
	Error interface{} `json:"errors"`
	Data  interface{} `json:"data"`
}

func APIResponse(message string, code int, data interface{}, errors interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  http.StatusText(code),
	}

	response := Response{
		Meta:  meta,
		Error: errors,
		Data:  data,
	}

	return response
}

type ErrorValidation struct {
	ExpectedKeyValue  string `json:"expected_key_value"`
	ExpectedTypeValue string `json:"expected_value"`
	Message           string `json:"message"`
}

func ErrorValidationFormat(err error) []ErrorValidation {
	var errors []ErrorValidation

	for _, e := range err.(validator.ValidationErrors) {
		valid := ErrorValidation{
			ExpectedKeyValue:  strings.ToLower(e.Field()),
			ExpectedTypeValue: e.Type().Name(),
			Message:           e.Error(),
		}
		errors = append(errors, valid)
	}

	return errors
}
