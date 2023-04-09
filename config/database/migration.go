package database

import (
	"fmt"
	"test-golang/config/mysql"
	"test-golang/models"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Product{},
		&models.User{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println(("Migration Success"))
}
