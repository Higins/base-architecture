package constants

import "fmt"

var ErrUnauthorized = fmt.Errorf("unauthorized")
var ErrMissingUserFromRequest = fmt.Errorf("missing user from request")
var ErrMissingUserIDFromRequest = fmt.Errorf("missing user id from request")
var ErrInternalError = fmt.Errorf("internal error")
var ErrUserIDNotMatch = fmt.Errorf("user id not match")
var ErrDatabaseError = fmt.Errorf("database error")
var ErrJWTError = fmt.Errorf("jwt error")
var ErrTokenExpired = fmt.Errorf("token is expired")
var ErrMissingUserFromCtx = fmt.Errorf("missing user id from ctx")
var ErrInvalidEmailOrPassword = fmt.Errorf("invalid email or password")
