package errores

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ErrorReponseApi struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// ErrResponse no puede recibir un nil
func ErrorResponse(c echo.Context, err error) error {
	var errc *CustomError
	code := 400
	message := "algo paso, hubo un error no esperado"
	if errors.As(err, &errc) {
		code = errc.httpCode
		message = errc.message
	}
	go func(e error, ec *CustomError) {
		if ec == nil {
			logrus.Error(e.Error())
			return
		}
		if ec.GetError() != nil {
			logrus.Error(e.Error())
			return
		}
	}(err, errc)
	return c.JSON(code, &ErrorReponseApi{Code: code, Message: message})
}
func JSONErrorResponse(c echo.Context, err error) error {
	errc := NewBadRequestf(err, ErrInvalidJSON)
	return ErrorResponse(c, errc)
}

///Mejorar la gestion de errores generados desde la base de datos
func NewInternalDBf(err error) error { /// TODO
	return newErrf(err, ErrDatabaseRequest, http.StatusInternalServerError)
}
