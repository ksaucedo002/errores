## Errores
Gestiona los errores, logs y los mensajes que se devuelven al cliente HTTP
### Instalacion
`go get github.com/ksaucedo002/errores`

### Error DefaultMessages
```go
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
```

### Example

```go
    if err:=db.DoSomething();err!=nil{                
        // Este error imprimir aun log, mostrando el valor de la variable err
        // en caso esta se nil, no se imprimirá nada,
        // el mensaje sera el que se envíe al cliente que realizo la petición del 
        // recurso
        return errores.NewInternalf(err,"<mensaje para el cliente>")        
    }
    return nil
```
Echo Framework Handler
```go
    func BooksHandler(c echo.Context) error{
        data,err:=h.service.FindBooks()
        if err!=nil{
            // Errors, determina el mensaje correspondiente
            // Para el error
            return errors.ErrorResponse(c,err)
        }
        return c.JSON(http.StatusOK,<data>)
    }
```
Echo Framework Handler, JSON Bind
```go
    func BooksHandler(c echo.Context) error{
       books := &models.Books{}
       if err:=c.Bind(books);err!=nil{
           // Genera mensaje y log para json parsing error
           return errores.JSONErrorResponse(c,err)
       }
    }
```