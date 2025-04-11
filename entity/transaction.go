package entity

import (
	"time"
)

type Transaction struct {
	Id                uint
	Invoice           string
	Nama              string
	Total_harga       int
	Total_bayar       int
	Total_kembalian   int
	Metode_pembayaran string
	User_id           uint
	User              User `gorm:"foreignKey:User_id"`
	Created_at        time.Time
}
