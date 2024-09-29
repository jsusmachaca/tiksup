package validation

import "errors"

var ErrInvalidDataType = errors.New("invalid data type")
var ErrIncorrectCredentials = errors.New("username or password incorrect")
