package services

import (
	"api/models"
)

type CategoryService interface {
	CreateCategory(*models.Category) (*models.Category, error)
	FindAll() ([]*models.Category, error)
}
