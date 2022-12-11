package errs

import "fmt"

type HttpError struct {
	Code    int
	Message string
}

func (e HttpError) Error() string {
	return fmt.Sprint(e.Code) + e.Message
}

func NewHttpError(code int, message string) HttpError {
	return HttpError{
		Code:    code,
		Message: message,
	}
}

func (e HttpError) Is(err error) bool {
	return e.Error() == err.Error()
}
