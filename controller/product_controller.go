package controller

import (
	"errors"
	"fmt"
	"net/http"
	"pos/service"
	"pos/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetProduct(c *gin.Context) {
	search := c.Query("search")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	products, meta, err := service.GetProducts(search, limit, page)

	if err != nil {
		utils.ResponseError(c, http.StatusOK, err.Error())
	}

	utils.ResponseSuccess(c, http.StatusOK, products, "", meta)
}

func CreateProduct(c *gin.Context) {

	var input service.ProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			e := ve[0]
			field := e.Field()
			tag := e.Tag()

			var msg string

			switch tag {
			case "required":
				msg = fmt.Sprintf("field %s cannot be empty", strings.ToLower(field))
			case "min":
				msg = "price and stock must be above 0"
			}

			utils.ResponseError(c, http.StatusBadRequest, msg)
			return
		}
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := service.CreateProduct(input); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, nil, "product added")
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var input service.ProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := service.UpdateProduct(id, input); err != nil {
		utils.ResponseError(c, http.StatusNotFound, err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, nil, "product updated")
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var input service.ProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := service.DeleteProduct(id, input); err != nil {

	}

	utils.ResponseSuccess(c, http.StatusAccepted, nil, "product deleted")
}
