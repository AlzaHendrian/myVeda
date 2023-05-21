package database

import (
	"fmt"
	"myveda/models"
	"myveda/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Transaction{},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success to Migration!")
}
