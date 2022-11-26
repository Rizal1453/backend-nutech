package database

import (
	"fmt"
	"nutech/models"
	"nutech/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.Products{},&models.User{})
	if err != nil{
		fmt.Println(err)
		panic("Migration Failed")
		
	}

	fmt.Println("Migration Success")
}