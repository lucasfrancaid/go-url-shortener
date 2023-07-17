package presenter

import (
	"encoding/json"
	"net/http"
)

type PresenterHTTP struct {
	Data       any
	Headers    string `json:"headers"`
	JsonData   []byte `json:"data"`
	StatusCode int    `json:"statusCode"`
}

var statusToHTTP = map[status]int{
	SUCCESS_CODE:          http.StatusOK,
	REDIRECT_CODE:         http.StatusSeeOther,
	CUSTOM_ERROR_CODE:     http.StatusBadRequest,
	NOT_FOUND_ERROR_CODE:  http.StatusNotFound,
	CONFLICT_ERROR_CODE:   http.StatusConflict,
	VALIDATION_ERROR_CODE: http.StatusUnprocessableEntity,
	INTERNAL_ERROR_CODE:   http.StatusInternalServerError,
	UNKNOWN_ERROR_CODE:    http.StatusInternalServerError,
}

type ErrorResponseHTTP struct {
	Error string `json:"error"`
}

func (p *Presenter) HTTP() PresenterHTTP {
	var data any
	var jsonData []byte
	if p.Error != nil {
		data = ErrorResponseHTTP{Error: p.Error.Error()}
		jsonData, _ = json.Marshal(data)
	} else {
		data = p.Data
		jsonData, _ = p.ToJSON()
	}

	return PresenterHTTP{
		Data:       data,
		Headers:    p.Headers,
		JsonData:   jsonData,
		StatusCode: statusToHTTP[p.StatusCode],
	}
}
