package errors


type InvalidNumericParameterInputError struct {
	message string
}

func (err InvalidNumericParameterInputError) Error() string {
  return err.message
}

func NewInvalidNumericParameterInputError() error {
	return InvalidNumericParameterInputError{message: "invalid input parameter type. provide an integer instead"}
}