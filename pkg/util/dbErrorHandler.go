package util

import (
	"errors"
	standardizedErrors "ketemuditengah/hello/pkg/errors"
	"reflect"

	"gorm.io/gorm"
)

func DbErrorHandler[T any](callable func() (*gorm.DB, T, error)) (T, error) {

	_, result, err := callable()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *new(T), &standardizedErrors.ResourceNotFoundError{
				Err:             err.Error(),
				Message:         "Resource not found ",
				Code:            404,
				ResourceType:    reflect.TypeOf(*new(T)).String(),
				IdentifierType:  "Id",
				IdentifierValue: "",
			}
		}

		return *new(T), err
	}

	return result, nil
}
