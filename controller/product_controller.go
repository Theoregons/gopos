package controller

import (
	"net/http"
	"pos/config"
	"pos/entity"
	"pos/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	var products []entity.Product
	var totalData int64

	search := c.Query("search")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	query := config.DB.Model(&entity.Product{})
	query.Count(&totalData)

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	totalPage := int((totalData + int64(limit) - 1) / int64(limit))

	if page > totalPage {
		utils.ResponseError(c, http.StatusNotFound, "halaman tidak ditemukan")
		return
	}

	query.Offset((page - 1) * limit).Limit(limit).Find(&products)

	meta := utils.PaginationMeta{
		TotalData:   totalData,
		Limit:       limit,
		CurrentPage: page,
		TotalPage:   totalPage,
	}

	utils.ResponseSuccess(c, http.StatusOK, products, "", meta)
}

func CreateProduct(c *gin.Context) {
	var input entity.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Price < 0 || input.Stock < 0 {
		utils.ResponseError(c, http.StatusBadRequest, "price and stock must be above 0 ")
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, nil, "product added")
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product entity.Product
	var input entity.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.First(&product, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "product not found")
		return
	}

	config.DB.Model(&product).Updates(input)

	utils.ResponseSuccess(c, http.StatusCreated, nil, "product updated")
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var product entity.Product
	var input entity.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.First(&product, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, "product not found")
		return
	}

	config.DB.Model(&product).Delete(&product)
	utils.ResponseSuccess(c, http.StatusAccepted, nil, "product deleted")

}
