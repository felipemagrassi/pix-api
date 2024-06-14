package rest_err

import (
	"net/http"

	"github.com/felipemagrassi/pix-api/internal/internal_error"
)

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func ConvertError(internalErr *internal_error.InternalError) *RestErr {
	switch internalErr.Err {
	case "bad_request":
		causes := make([]Causes, 0)
		for _, cause := range internalErr.Causes {
			causes = append(causes, Causes{Field: cause.Field, Message: cause.Message})
		}
		return NewBadRequestError(internalErr.Message, causes...)
	case "not_found":
		return NewNotFoundError(internalErr.Message)
	default:
		return NewInternalServerError(internalErr.Message, internalErr.OriginalError)
	}
}

func NewBadRequestError(message string, causes ...Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}

	if err != nil {
		result.Causes = append(result.Causes, Causes{Field: "internal", Message: err.Error()})
	}

	return result
}
