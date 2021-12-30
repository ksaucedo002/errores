package errores

import (
	"errors"
	"net/http"

	"github.com/jackc/pgconn"
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
func JSONErrorResponse(c echo.Context) error {
	return ErrorResponse(c, NewBadRequestf(nil, ErrInvalidJSON))
}
func QueryErrorResponse(c echo.Context) error {
	return ErrorResponse(c, NewBadRequestf(nil, ErrInvalidQueryParam))
}

///Getion de errores DB Postgres
type action struct {
	Message          string
	HttpResponseCode int
	Loggable         bool
}

var pgErrorMessage = map[string]action{
	"23505": {"el registro ya existe", http.StatusBadRequest, false},
	"23514": {"formato incorrecto de datos, consulte la documentación", http.StatusBadRequest, false},
	"23503": {"referencia incompatible con recurso, consulte la documentación", http.StatusBadRequest, false},
	"23000": {"operación restringida, problema de integridad con los datos, consulte documentación", http.StatusBadRequest, false},
	"25000": {"no se pudo completar con las operaciones", http.StatusInternalServerError, true},
	"26000": {"hubo un problema interno, por favor reporte la incidencia al equipo técnico respectivo", http.StatusInternalServerError, true},
	"28000": {"acceso restringido", http.StatusUnauthorized, true},
	"2D000": {"transacción inválida", http.StatusInternalServerError, true},
}

func NewInternalDBf(err error) error {
	var pgerr *pgconn.PgError
	if errors.As(err, &pgerr) {
		act, ok := pgErrorMessage[pgerr.Code]
		if ok {
			if act.Loggable {
				return newErrf(err, act.Message, act.HttpResponseCode)
			}
			return newErrf(nil, act.Message, act.HttpResponseCode)
		}
	}
	return newErrf(err, ErrDatabaseInternal, http.StatusInternalServerError)
}
