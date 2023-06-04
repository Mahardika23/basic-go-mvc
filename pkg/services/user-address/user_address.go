package useraddress

import (
	"ketemuditengah/hello/pkg/model"
	user_address_repository "ketemuditengah/hello/pkg/repository/user-address"
	"ketemuditengah/hello/pkg/util"

	"gorm.io/gorm"
)

func GetUserAddressByUserId(userId int32) ([]model.UserAddress, error) {
	return util.DbErrorHandler(func() (*gorm.DB, []model.UserAddress, error) {
		return user_address_repository.GetUserAddressByUserId(userId)
	})
}
