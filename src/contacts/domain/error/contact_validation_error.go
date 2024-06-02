package errors

type ContactValidationError struct {
	Message string
	Type    string
}

func NewContactValidationError(message, typeError string) ContactValidationError {
	return ContactValidationError{
		Message: message,
		Type:    typeError,
	}

}
func (e ContactValidationError) Error() string {
	return e.Message
}
