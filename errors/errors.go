package errors

// InvalidNumericParameterInputError
type InvalidNumericParameterInputError struct {
	message string
}

func (err InvalidNumericParameterInputError) Error() string {
  return err.message
}

func NewInvalidNumericParameterInputError() error {
	return InvalidNumericParameterInputError{message: "invalid input parameter type. provide an integer instead"}
}


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

// UnauthorizedError
type UnauthorizedError struct {
	message string
}

func (err UnauthorizedError) Error() string {
  return err.message
}

func NewUnauthorizedError() error {
	return IncorrectQueryParameterError{message: "You are unauthorized to make this request."}
}