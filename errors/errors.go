package errors


type NotFoundError struct {
	message string
}

func (err NotFoundError) Error() string {
  return err.message
}

func NewNotFoundError(message string) error {
	return NotFoundError{message: message}
}