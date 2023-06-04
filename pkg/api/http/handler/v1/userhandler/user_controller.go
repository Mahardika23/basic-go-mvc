package userhandler

import (
	"fmt"
	userservice "ketemuditengah/hello/pkg/services/user"
	user_address_service "ketemuditengah/hello/pkg/services/user-address"

	"ketemuditengah/hello/pkg/util"
	"net/http"

	"ketemuditengah/hello/pkg/model"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var input model.User

	err := c.Bind(&input)

	fmt.Println(input)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return util.HttpHandler(c, func() (model.User, error) {
		if err = c.Validate(input); err != nil {
			return *new(model.User), err
		}

		return userservice.Create(input)
	})
}

func GetUserByIdHandler(c echo.Context) error {
	var user model.User

	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return util.HttpHandler(c, func() (model.User, error) {
		return userservice.GetUserById(int32(user.ID))
	})
}

func GetUserAddressByUserId(c echo.Context) error {
	var user model.User

	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return util.HttpHandler(c, func() ([]model.UserAddress, error) {
		return user_address_service.GetUserAddressByUserId(int32(user.ID))
	})

}
