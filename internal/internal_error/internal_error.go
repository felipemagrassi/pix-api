package internal_error

type InternalError struct {
	Message string
	Err     string
}

func (ie *InternalError) Error() string {
	return ie.Message
}

func NewBadRequestError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "bad_request",
	}
}
