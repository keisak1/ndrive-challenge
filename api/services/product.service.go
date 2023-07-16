package services

import "api/models"

type ProductService interface {
	AddProduct(*models.ProductInsert) (*models.DBProductResponse, error)
	DeleteProduct(*models.ProductDelete) (*models.DBProductResponse, error)
	FindAll() ([]models.Product, error)
}
