package entity

type TransactionItem struct {
	Id             uint `gorm:"primaryKey"`
	Product_id     uint
	Product        Product `gorm:"foreignKey:Product_id"`
	Unit_price     int
	Quantity       int
	Transaction_id uint
	Transaction    Transaction `json:"-"`
}
