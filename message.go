package errores

const (
	ErrInvalidJSON           = "estructura json invalida"
	ErrInvalidToken          = "el token es invalido"
	ErrTokenNull             = "no se encontro el token"
	ErrSigningTokenString    = "no se pudo authentificar"
	ErrNoDefined             = "hubo un error, no esperado"
	ErrDatabaseRequest       = "no se pudo realizar la operacion"
	ErrRecordNotFaund        = "no se encontrar el registro"
	ErrRecord                = "no se pudo guardar el registro"
	ErrUsernameExists        = "el usuario ya existe"
	ErrAuthorizationHeader   = "Authorization header no encontrado"
	ErrUserOrPasswordInvalid = "usuario o password incorrectos"
)
