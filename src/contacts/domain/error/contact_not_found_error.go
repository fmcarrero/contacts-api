package errors

type ContactNotFoundError struct {
	Message string
	Type    string
}

func NewContactNotFoundError(message, errorType string) ContactNotFoundError {
	return ContactNotFoundError{Message: message, Type: errorType}
}

func (e ContactNotFoundError) Error() string {
	return e.Message
}
