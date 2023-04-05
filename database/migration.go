package database

import (
	"dumbflix/models"
	"dumbflix/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.Users{})

	if err!= nil {
        fmt.Println(err)
		panic("Migration failed")
    }

	fmt.Println("Migration success")
}