package validation

import "errors"

var ErrInvalidDataType = errors.New("invalid data type")
var ErrIncorrectCredentials = errors.New("username or password incorrect")
var ErrObteinData = errors.New("data wasn't obtained")
var ErrRowsAffected = errors.New("1 row was expected to be affect")
var ErrInvalidToken = errors.New("invalid json web token")
