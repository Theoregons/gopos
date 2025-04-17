package controller

import (
	"net/http"
	"pos/config"
	"pos/entity"
	"pos/utils"

	"github.com/gin-gonic/gin"
)

type TopProductReport struct {
	ProductID uint
	Name      string
	TotalQty  int
}

type LowStockReport struct {
	ID    uint
	Name  string
	Stock int
}

func GetReports(c *gin.Context) {
	var topProducts []TopProductReport
	var lowStocks []LowStockReport

	config.DB.
		Table("transaction_items as ti").
		Select("p.id as product_id, p.name as nama, SUM(ti.quantity) as total_qty").
		Joins("JOIN products p ON p.id = ti.product_id").
		Group("p.id, p.name").
		Order("total_qty").
		Limit(5).
		Scan(&topProducts)

	config.DB.
		Model(&entity.Product{}).
		Select("id, name, stock").
		Where("stock < ?", 10).
		Order("stock ASC").
		Scan(&lowStocks)

	utils.ResponseSuccess(c, http.StatusOK, gin.H{"top_selling_products": topProducts, "low_stock_products": lowStocks}, "")
}
