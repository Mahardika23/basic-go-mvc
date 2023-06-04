package userservice

import (
	"errors"
	ApplicationErrors "ketemuditengah/hello/pkg/errors"
	"ketemuditengah/hello/pkg/model"
	userrepository "ketemuditengah/hello/pkg/repository/user"
	"ketemuditengah/hello/pkg/util"

	"gorm.io/gorm"
)

func GetUserById(id int32) (model.User, error) {
	_, user, err := userrepository.GetUserById(int32(id))

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, &ApplicationErrors.ResourceNotFoundError{
			Err:             err.Error(),
			Message:         "Resource not found ",
			Code:            404,
			ResourceType:    "User",
			IdentifierType:  "Id",
			IdentifierValue: string(id),
		}
	}

	return user, nil
}

func Create(user model.User) (model.User, error) {
	return util.DbErrorHandler(func() (*gorm.DB, model.User, error) {
		return userrepository.Create(user)
	})
}
