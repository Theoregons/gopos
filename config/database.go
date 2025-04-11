package config

import (
	"fmt"
	"log"
	"pos/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/gopos?parseTime=true"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = DB.AutoMigrate(
		&entity.User{},
		&entity.Product{},
		&entity.Transaction{},
		&entity.TransactionItem{},
	)
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	fmt.Println("Database migrated successfully!")
}
