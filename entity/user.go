package entity

import "time"

type User struct {
	Id         uint `gorm:"primaryKey"`
	Name       string
	Email      string `gorm:"unique"`
	Password   string
	Role       string
	Created_at time.Time
}
