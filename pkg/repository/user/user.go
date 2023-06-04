package user

import (
	"fmt"
	dbprovider "ketemuditengah/hello/pkg/db"
	"ketemuditengah/hello/pkg/model"

	"gorm.io/gorm"
)

var conn *gorm.DB = dbprovider.GetDbConnection("user repository")

func GetUserById(id int32) (*gorm.DB, model.User, error) {
	var user model.User

	result := conn.First(&user, id)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, user, result.Error
	}

	return result, user, nil
}

func Create(user model.User) (*gorm.DB, model.User, error) {

	result := conn.Create(&user)

	if result.Error != nil {
		panic(result.Error)
	}

	return result, user, nil
}
