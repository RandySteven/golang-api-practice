package apperror

type CustomError struct {
	Resource string
	Field    string
	Value    string
	Message  string
	ErrType  string
	Err      error
}

type (
	ErrFieldValidation   error
	ErrDataNotFound      error
	ErrDuplicatedRequest error
)

func NewCustomError(resource, field, value, message, errType string, err error) *CustomError {
	return &CustomError{
		Resource: resource,
		Field:    field,
		Value:    value,
		Message:  message,
		ErrType:  errType,
		Err:      err,
	}
}

// func (e *CustomError) Error() string {
// 	return e.Message
// }

// func (e *ErrDataNotFound) Error() string {
// 	return e.Err.Error()
// }

// func (e *ErrFieldValidation) Error() string {
// 	return e.Err.Error()
// }

// func (e *ErrDuplicatedRequest) Error() string {
// 	return e.Err.Error()
// }
