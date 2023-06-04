package model

import "time"

type User struct {
	ID          uint   `param:"id" query:"id" form:"id" json:"id" xml:"id" gorm:"autoIncrement"`
	Name        string `json:"name" validate:"required"`
	Age         int16  `json:"age" validate:"required,lte=12"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserAddress struct {
	ID           uint32 `json:"id"`
	AddressLine1 string `json:"address_line_1"`
	Province     string `json:"province"`
	District     string `json:"district"`
	Subdistrict  string `json:"sub_district"`
	Village      string `json:"village"`
	UserID       uint   `json:"user_id"`
	User         User
}
