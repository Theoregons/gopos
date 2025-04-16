package controller

import (
	"net/http"
	"pos/config"
	"pos/entity"
	"pos/utils"

	"github.com/gin-gonic/gin"
)

type ItemInput struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
	UnitPrice int  `json:"harga_satuan"`
}

type TransactionInput struct {
	CustomerName  string      `json:"customer_name"`
	PaymentMethod string      `json:"payment_method"`
	TotalPayment  int         `json:"total_payment"`
	Items         []ItemInput `json:"items"`
}

func CreateTransaction(c *gin.Context) {
	var input TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	var totalPrice int
	var items []entity.TransactionItem

	tx := config.DB.Begin()

	for _, item := range input.Items {
		var product entity.Product
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			utils.ResponseError(c, http.StatusNotFound, "product tidak ditemukan")
			return
		}
		if product.Stock < item.Quantity {
			tx.Rollback()
			utils.ResponseError(c, http.StatusBadRequest, `"stock product tidak cukup:" product.Name`)
			return
		}

		product.Stock -= item.Quantity
		tx.Save(&product)

		totalPrice += item.Quantity * item.UnitPrice
		items = append(items, entity.TransactionItem{
			Product_id: item.ProductID,
			Quantity:   item.Quantity,
			Unit_price: item.UnitPrice,
		})
	}

	totalChange := input.TotalPayment - totalPrice

	if totalChange < 0 {
		tx.Rollback()
		utils.ResponseError(c, http.StatusBadRequest, "total bayar kurang dari total harga")
		return
	}

	transaction := entity.Transaction{
		Invoice:          "INV - " + utils.RandString(8),
		Customer_name:    input.CustomerName,
		Payment_method:   input.PaymentMethod,
		Total_price:      totalPrice,
		Total_payment:    input.TotalPayment,
		Total_change:     totalChange,
		User_id:          c.GetUint("user_id"),
		TransactionItems: items,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		utils.ResponseError(c, http.StatusInternalServerError, "gagal membuat transaksi")
		return
	}

	tx.Commit()
	utils.ResponseSuccess(c, http.StatusCreated, transaction, "")

}

func GetTransactions(c *gin.Context) {
	var transactions []entity.Transaction
	if err := config.DB.Preload("User").
		Preload("TransactionItems").
		Preload("TransactionItems.Product").
		Order("created_at DESC").
		Find(&transactions).Error; err != nil {

		utils.ResponseError(c, http.StatusNotFound, "data tidak ditemukan")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, transactions, "ok")
}
