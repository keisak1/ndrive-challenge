package services

import (
	"api/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService interface {
	AddProduct(*models.ProductInsert) (*models.DBProductResponse, error)
	DeleteProduct(string) (*mongo.DeleteResult, error)
	FindAll() ([]*models.Product, error)
	FindOne(string) (*models.Product, error)
	EditProduct(*models.ProductEdit) (*models.ProductEdit, error)
	SearchProduct(string) ([]*models.Product, error)
}
