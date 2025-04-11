package entity

type TransactionItem struct {
	Id             uint `gorm:"primaryKey"`
	Product_id     uint
	Product        Product `gorm:"foreignKey:Product_id"`
	Harga_satuan   int
	Quantity       int
	Transaction_id uint
	Transaction    Transaction `gorm:"foreignKey:Transaction_id"`
}
