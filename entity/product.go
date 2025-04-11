package entity

import (
	"time"
)

type Product struct {
	Id         uint `gorm:"primaryKey"`
	Nama       string
	Jenis      string
	Stok       int
	Harga      int
	Created_at time.Time
}
