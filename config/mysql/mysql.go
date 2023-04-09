package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseInit() {

	// var DB_HOST = os.Getenv("DB_HOST")
	// var DB_USER = os.Getenv("DB_USER")
	// var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	// var DB_NAME = os.Getenv("DB_NAME")
	// var DB_PORT = os.Getenv("DB_PORT")

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//todo if using mysql
	dsn := "root:@tcp(localhost:3306)/test-golang?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	// dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_HOST, DB_PORT, DB_NAME)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}
