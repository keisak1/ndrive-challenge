package services

import (
	"api/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService interface {
	AddProduct(*models.ProductInsert) (*models.DBProductResponse, error)
	DeleteProduct(string) (*mongo.DeleteResult, error)
	FindAll() ([]*models.Product, error)
}
