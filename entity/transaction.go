package entity

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Invoice        string
	Customer_name  string
	Total_price    int
	Total_payment  int
	Total_change   int
	Payment_method string
	User_id        uint
	User           User `gorm:"foreignKey:User_id"`
}
