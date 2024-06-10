package value_object

import (
	"fmt"
	"regexp"

	"github.com/felipemagrassi/pix-api/internal/internal_error"
)

type Email string

const (
	EmailKeyPattern = `^[a-z0-9+_.-]+@[a-z0-9.-]+$`
)

func NewEmail(email string) (Email, *internal_error.InternalError) {
	newEmail := Email(email)

	if err := newEmail.Validate(); err != nil {
		return "", err
	}

	return newEmail, nil
}

func (e Email) String() string {
	return string(e)
}

func (e Email) Validate() *internal_error.InternalError {
	re, err := regexp.Compile(EmailKeyPattern)
	if err != nil {
		message := fmt.Sprintf("Error compiling regex: %s", err.Error())
		return internal_error.NewInternalServerError(message)
	}

	if !re.MatchString(e.String()) {
		return internal_error.NewBadRequestError("Invalid Email")
	}

	return nil
}
