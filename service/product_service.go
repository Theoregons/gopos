package service

import (
	"errors"
	"pos/config"
	"pos/entity"
	"pos/utils"
)

type ProductInput struct {
	Name     string `json:"name" binding:"required,min=3"`
	Category string `json:"category" binding:"required"`
	Price    int    `json:"price" binding:"required,min=0"`
	Stock    int    `json:"stock" binding:"required,min=0"`
}

func GetProducts(search string, limit, page int) ([]entity.Product, utils.PaginationMeta, error) {
	var products []entity.Product
	var totalData int64

	query := config.DB.Model(&entity.Product{})
	query.Count(&totalData)

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	totalPage := int((totalData + int64(limit) - 1) / int64(limit))

	if page > totalPage && totalPage != 0 {
		return nil, utils.PaginationMeta{}, errors.New("halaman tidak ditemukan")
	}

	query.Offset((page - 1) * limit).Limit(limit).Find(&products)

	meta := utils.PaginationMeta{
		TotalData:   totalData,
		Limit:       limit,
		CurrentPage: page,
		TotalPage:   totalPage,
	}

	return products, meta, nil
}

func CreateProduct(input ProductInput) error {
	product := entity.Product{
		Name:     input.Name,
		Category: input.Category,
		Price:    input.Price,
		Stock:    input.Stock,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func UpdateProduct(id string, input ProductInput) error {
	var product entity.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		return errors.New("product not found")
	}

	config.DB.Model(&product).Updates(input)
	return nil
}

func DeleteProduct(id string, input ProductInput) error {
	var product entity.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		return errors.New("product not found")
	}

	config.DB.Model(&product).Delete(&product)
	return nil
}
