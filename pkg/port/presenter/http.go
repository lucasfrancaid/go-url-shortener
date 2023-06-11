package presenter

import (
	"encoding/json"
	"net/http"
)

type PresenterHTTP struct {
	Headers    string `json:"headers"`
	Data       []byte `json:"data"`
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
	var data []byte
	if p.Error != nil {
		data, _ = json.Marshal(ErrorResponseHTTP{Error: p.Error.Error()})
	} else {
		data = p.Data
	}

	return PresenterHTTP{
		Headers:    p.Headers,
		Data:       data,
		StatusCode: statusToHTTP[p.StatusCode],
	}
}
