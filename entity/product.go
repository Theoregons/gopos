package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string
	Category string
	Stock    int
	Price    int
}
