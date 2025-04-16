package controller

import (
	"net/http"
	"pos/config"
	"pos/entity"
	"pos/utils"

	"github.com/gin-gonic/gin"
)

type ItemInput struct {
	Id         uint `json:"product_id"`
	Quantity   int  `json:"quantity"`
	Unit_price int  `json:"harga_satuan"`
}

type TransactionInput struct {
	Customer_name  string      `json:"customer_name"`
	Payment_method string      `json:"payment_method"`
	Total_payment  int         `json:"total_payment"`
	Items          []ItemInput `json:"items"`
}

func CreateTransaction(c *gin.Context) {
	var input TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var totalHarga int
	var items []entity.TransactionItem

	tx := config.DB.Begin()

	for _, item := range input.Items {
		var product entity.Product
		if err := tx.First(&product, item.Id).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Produk tidak ditemukan"})
			return
		}

		if product.Stock < item.Quantity {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Stok produk tidak cukup: " + product.Category})
			return
		}

		// Kurangi stok
		product.Stock -= item.Quantity
		tx.Save(&product)

		totalHarga += item.Quantity * item.Unit_price
		items = append(items, entity.TransactionItem{
			Product_id: item.Id,
			Quantity:   item.Quantity,
			Unit_price: item.Unit_price,
		})
	}
	c.JSON(http.StatusAccepted, input.Total_payment)

	totalKembalian := input.Total_payment - totalHarga
	if totalKembalian < 0 {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Total bayar kurang dari total harga"})
		return
	}

	transaction := entity.Transaction{
		Invoice:          "INV-" + utils.RandString(8),
		Customer_name:    input.Customer_name,
		Payment_method:   input.Payment_method,
		Total_price:      totalHarga,
		Total_payment:    input.Total_payment,
		Total_change:     totalKembalian,
		User_id:          1, //c.GetUint("user_id"),
		TransactionItems: items,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat transaksi"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusCreated, transaction)
}
