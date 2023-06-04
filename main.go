package main

import (
	"fmt"
	"ketemuditengah/hello/pkg/api/http/handler/v1/userhandler"
	"ketemuditengah/hello/pkg/db"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}

	return nil
}

func init() {
	db.InitiateDbConnection("MAIN", "MAIN")
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	userRoutes := e.Group("/users")

	userRoutes.POST("/register", userhandler.Register)
	userRoutes.GET("/:id", userhandler.GetUserByIdHandler)
	userRoutes.GET("/:id/addresses", userhandler.GetUserAddressByUserId)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
