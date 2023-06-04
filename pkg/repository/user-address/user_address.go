package user_address

import (
	"fmt"
	dbprovider "ketemuditengah/hello/pkg/db"
	"ketemuditengah/hello/pkg/model"

	"gorm.io/gorm"
)

var conn *gorm.DB = dbprovider.GetDbConnection("user address repository")

func GetUserAddressByUserId(userId int32) (*gorm.DB, []model.UserAddress, error) {
	var userAddress []model.UserAddress

	results := conn.Where("user_id = ?", userId).Find(&userAddress)

	if results.Error != nil {
		fmt.Println(results.Error)
		return nil, userAddress, results.Error
	}

	return results, userAddress, nil
}
