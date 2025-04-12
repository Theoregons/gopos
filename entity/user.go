package entity

import "time"

type User struct {
	Id         uint      `gorm:"primaryKey"`
	Name       string    `json:"name"`
	Email      string    `gorm:"unique"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	Created_at time.Time `gorm:"autoCreateTime"`
}
