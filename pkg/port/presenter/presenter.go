package presenter

import (
	"encoding/json"

	base "github.com/lucasfrancaid/go-url-shortener/pkg/application/base"
)

type status int

const (
	SUCCESS_CODE          status = 1
	REDIRECT_CODE         status = 2
	INTERNAL_ERROR_CODE   status = 3
	UNKNOWN_ERROR_CODE    status = 4
	NOT_FOUND_ERROR_CODE  status = 5
	VALIDATION_ERROR_CODE status = 6
	CONFLICT_ERROR_CODE   status = 7
	CUSTOM_ERROR_CODE     status = 8
)

type Presenter struct {
	Headers    string
	Data       any
	StatusCode status
	Error      error
}

func (p *Presenter) ToJSON() ([]byte, error) {
	j, err := json.Marshal(p.Data)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func PresenterSuccess(data any) Presenter {
	return Presenter{
		Data:       data,
		StatusCode: SUCCESS_CODE,
	}
}

func PresenterRedirect(data any) Presenter {
	return Presenter{
		Data:       data,
		StatusCode: REDIRECT_CODE,
	}
}

func PresenterError(err error) Presenter {
	statusCode := UNKNOWN_ERROR_CODE
	if baseErr, ok := err.(*base.Error); ok {
		switch baseErr.Type {
		case base.NOT_FOUND_ERROR:
			statusCode = NOT_FOUND_ERROR_CODE
		case base.VALIDATOR_ERROR:
			statusCode = VALIDATION_ERROR_CODE
		case base.CONFLICT_ERROR:
			statusCode = CONFLICT_ERROR_CODE
		case base.CUSTOM_ERROR:
			statusCode = CUSTOM_ERROR_CODE
		default:
			statusCode = UNKNOWN_ERROR_CODE
		}
	}
	return Presenter{
		StatusCode: statusCode,
		Error:      err,
	}
}

// TODO: Implement another presenters
// func (p *Presenter) gRPC()    {}
// func (p *Presenter) Event()   {}
// func (p *Presenter) Message() {}
