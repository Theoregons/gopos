package entity

import (
	"time"
)

type Product struct {
	Id         uint `gorm:"primaryKey"`
	Jenis      string
	Stok       int
	Harga      int
	Created_at time.Time
}
