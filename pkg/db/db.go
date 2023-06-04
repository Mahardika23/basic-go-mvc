package db

import (
	"fmt"
	"ketemuditengah/hello/pkg/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func GetDbConnection(context string) *gorm.DB {

	godotenv.Load()

	if DB != nil {
		fmt.Println("DB Conn already exists", context)

		return DB
	}

	DB = InitiateDbConnection("FROM GET DB CONN", context)

	return DB
}

func InitiateDbConnection(context string, parentcontext string) *gorm.DB {
	fmt.Println("Initiating new db conn", context, parentcontext)
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	return db
}

func MigrateDb() {
	conn := GetDbConnection("Migrate Db")

	conn.AutoMigrate(&model.User{}, &model.UserAddress{})
}
