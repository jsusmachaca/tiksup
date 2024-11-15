package auth

import (
	"encoding/json"
	"errors"
	"io"
)

func UserValidation(body io.ReadCloser, model any) error {
	err := json.NewDecoder(body).Decode(model)
	if err != nil {
		var unmarshallError *json.UnmarshalTypeError
		if errors.As(err, &unmarshallError) {
			return ErrInvalidDataType
		}
		return err
	}

	return nil
}
