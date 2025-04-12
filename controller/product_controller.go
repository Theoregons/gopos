package controller

import (
	"net/http"
	"pos/config"
	"pos/entity"
	"pos/utils"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	var products []entity.Product

	if err := config.DB.Find(&products).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Gagal mengambil data")
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, products)
}
