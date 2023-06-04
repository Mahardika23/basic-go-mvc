package util

import (
	"fmt"
	ApplicationErrors "ketemuditengah/hello/pkg/errors"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func HttpHandler[T any](echoContext echo.Context, f func() (T, error)) error {
	var res T
	res, err := f()

	if err != nil {

		fmt.Println(reflect.TypeOf(err))

		_, isResourceNotFoundError := err.(*ApplicationErrors.ResourceNotFoundError)
		if isResourceNotFoundError {
			return echoContext.JSON(http.StatusNotFound, err.Error())
		}

		_, isValidationError := err.(validator.ValidationErrors)

		if isValidationError {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}
		// Add more http cases here to be auto handled

		return echoContext.JSON(http.StatusInternalServerError, "Internal Server Error, please contact developer aryagamas@gmail.com")
	}

	return echoContext.JSON(http.StatusOK, res)
}
