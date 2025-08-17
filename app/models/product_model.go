// Package models for entities and other struct
package models

import (
	"strings"

	"github.com/aysmdb/ojire-casetiga/pkg/database"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
	Description string `json:"description" gorm:"type:text"`
	ImageURL    string `gorm:"default:'https://picsum.photos/200/300'"`
}

type CheckoutRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func GetProductList(page int, take int, search string) []Product {
	db := database.DBConn

	var products []Product
	if search == "" {
		db.Offset((page - 1) * take).Limit(take).Find(&products)
	} else {
		db.Where("lower(name) LIKE ?", "%"+strings.ToLower(search)+"%").Offset((page - 1) * take).Limit(take).Find(&products)
	}

	return products
}

func GetProductByID(id int) Product {
	db := database.DBConn
	var product Product
	db.Find(&product, id)

	return product
}

func UpdateProductQuantity(data CheckoutRequest) error {
	db := database.DBConn
	var product Product
	if err := db.First(&product, data.ProductID).Error; err != nil {
		return err
	}

	db.Model(&product).Update("quantity", product.Quantity-data.Quantity)
	return nil
}

func SeedProduct() error {
	products := []Product{
		{
			Name:     "Joran Premium",
			Quantity: 12,
			Price:    1500000,
		},
		{
			Name:     "Joran Standar",
			Quantity: 12,
			Price:    1000000,
		},
		{
			Name:     "Joran Ekonomi",
			Quantity: 12,
			Price:    500000,
		},
		{
			Name:     "Reel Premium",
			Quantity: 12,
			Price:    500000,
		},
		{
			Name:     "Reel Standar",
			Quantity: 12,
			Price:    300000,
		},
		{
			Name:     "Reel Ekonomi",
			Quantity: 12,
			Price:    100000,
		},
		{
			Name:     "Umpan Premium",
			Quantity: 12,
			Price:    50000,
		},
		{
			Name:     "Umpan Standar",
			Quantity: 12,
			Price:    20000,
		},
		{
			Name:     "Umpan Ekonomi",
			Quantity: 12,
			Price:    10000,
		},
	}

	err := database.DBConn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("name is not null").Delete(&Product{}).Error; err != nil {
			return err
		}

		if err := tx.Create(&products).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
