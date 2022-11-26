package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func DatabaseInit(){
	var err error

	dsn := "root:@tcp(localhost:3306)/nutech?charset=utf8mb4&parseTime=True&loc=Local"
	DB,err = gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil {
	fmt.Println("database error")
	}
	fmt.Println("database connected")
}