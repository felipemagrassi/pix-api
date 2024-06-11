package internal_error

type InternalError struct {
	Message       string
	Err           string
	OriginalError error
	Causes        []Causes
}

type Causes struct {
	Field   string
	Message string
}

func (ie *InternalError) Error() string {
	return ie.Message
}

func NewBadRequestError(message string, causes ...Causes) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "bad_request",
		Causes:  causes,
	}
}

func NewNotFoundError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "not_found",
	}
}

func NewInternalServerError(message string, err error) *InternalError {
	return &InternalError{
		Message:       message,
		Err:           "internal_server_error",
		OriginalError: err,
	}
}
