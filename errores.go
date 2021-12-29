package errores

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	httpCode int
	err      error
	message  string
}

func newErrf(err error, message string, httpcode int) error {
	return &CustomError{
		err:      err,
		message:  message,
		httpCode: httpcode,
	}
}

func NewBadRequestf(err error, format string, a ...interface{}) error {
	return newErrf(err, fmt.Sprintf(format, a...), http.StatusBadRequest)
}
func NewInternalf(err error, format string, a ...interface{}) error {
	return newErrf(err, fmt.Sprintf(format, a...), http.StatusInternalServerError)
}
func NewUnsupported(err error, format string, a ...interface{}) error {
	return newErrf(err, fmt.Sprintf(format, a...), http.StatusUnsupportedMediaType)
}
func NewUnauthorizedf(err error, format string, a ...interface{}) error {
	return newErrf(err, fmt.Sprintf(format, a...), http.StatusUnauthorized)
}
func NewForbiddenf(err error, format string, a ...interface{}) error {
	return newErrf(err, fmt.Sprintf(format, a...), http.StatusForbidden)
}
func NewNotFoundf(err error, format string, a ...interface{}) error {
	return newErrf(err, fmt.Sprintf(format, a...), http.StatusNotFound)
}
func (e *CustomError) Error() string {
	if e.err == nil {
		return "error trivial "
	}
	return e.err.Error()
}
func (e *CustomError) GetError() error {
	return e.err
}
func (e *CustomError) Message() string {
	return e.message
}
