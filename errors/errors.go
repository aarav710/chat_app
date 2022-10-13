package errors

import "errors"


var (
	UnauthorizedError = errors.New("you are unauthorized to make this request")
	InvalidNumericParameterInputError = errors.New("invalid input parameter type. provide an integer instead")
	InternalServerError = errors.New("an internal error has occured in the server")
)


// IncorrectQueryParameterError
type IncorrectQueryParameterError struct {
	message string
}

func (err IncorrectQueryParameterError) Error() string {
  return err.message
}

func NewIncorrectQueryParameterError(message string) error {
	return IncorrectQueryParameterError{message: message}
}
