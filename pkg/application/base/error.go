package base

type errorType string

const (
	CONFLICT_ERROR  errorType = "Error.Conflict"
	CUSTOM_ERROR    errorType = "Error.Custom"
	INTERNAL_ERROR  errorType = "Error.Internal"
	NOT_FOUND_ERROR errorType = "Error.NotFound"
	UNKNOWN_ERROR   errorType = "Error.Unknown"
	VALIDATOR_ERROR errorType = "Error.Validator"
)

type Error struct {
	Type errorType
	Err  error
}

func (e *Error) Error() string {
	return e.Err.Error()
}
